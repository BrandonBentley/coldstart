package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

//go:embed api client config entity service main/main.go go.mod go.sum Makefile .gitignore Dockerfile
var appDir embed.FS

var folderNames = []string{
	"api",
	"client",
	"config",
	"entity",
	"service",
}

var files = []string{
	"main.go",
	"go.mod",
	"go.sum",
	"Makefile",
	".gitignore",
	"Dockerfile",
}

var coldstartRegex = regexp.MustCompile("coldstart")

var coldstartModuleName = regexp.MustCompile("github.com/BrandonBentley/coldstart")

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Module name is required")
		os.Exit(1)
	}

	if os.Args[1] == "version" {
		printVersion()
		os.Exit(0)
	}

	if os.Args[1] == "help" {
		printHelp()
		os.Exit(0)
	}

	newModulePath := os.Args[1]

	newModuleName := filepath.Base(newModulePath)

	_, err := os.Stat(newModuleName)
	if err == nil {
		fmt.Printf("directory already exists -> %v\n", newModuleName)
		os.Exit(1)
	}

	for _, folderName := range folderNames {
		readFolder(folderName)
	}
	for _, file := range files {
		inputFileName := file
		if inputFileName == "main.go" {
			inputFileName = "main/" + inputFileName
		}
		inFile, err := appDir.Open(inputFileName)
		if err != nil {
			slog.Default().Error(
				"failed to open file",
				"file", file,
				"err", err,
			)
			os.Exit(1)
		}
		defer inFile.Close()

		newFilePath := filepath.Join(newModuleName, file)
		os.MkdirAll(filepath.Dir(newFilePath), 0755)

		outFile, err := os.Create(newFilePath)
		if err != nil {
			slog.Default().Error(
				"failed to create file",
				"file", file,
				"err", err,
			)
			os.Exit(1)
		}
		defer outFile.Close()

		scanner := bufio.NewScanner(inFile)
		for scanner.Scan() {
			fmt.Fprintln(outFile,
				coldstartRegex.ReplaceAllString(
					coldstartModuleName.ReplaceAllString(scanner.Text(), newModulePath),
					newModuleName,
				),
			)
		}
	}

	readMeFile, err := os.Create(filepath.Join(newModuleName, "README.md"))
	if err != nil {
		slog.Default().Error(
			"failed to create file",
			"file", readMeFile,
			"err", err,
		)
		os.Exit(1)
	}
	fmt.Fprintf(readMeFile, readmeFmtString, newModuleName)

	cmd := exec.Command("sh", "-c", fmt.Sprintf(initRepoScriptFmtString, newModuleName))

	err = cmd.Run()
	if err != nil {
		slog.Default().Error(
			"failed to run init repo script",
			"err", err,
		)
		os.Exit(1)
	}
}

const initRepoScriptFmtString = `
#/bin/sh
cd %v
go mod tidy
git init
git add .gitignore
git add *
git commit -m "initial commit from coldstart"
`

const readmeFmtString = `
# %v
Initialized via [coldstart](https://github.com/BrandonBentley/coldstart)
`

func readFolder(folderName string) {
	dirEntries, err := appDir.ReadDir(folderName)
	if err != nil {
		log.Fatalln("could not open folder", folderName)
	}
	for _, de := range dirEntries {
		fullPath := filepath.Join(folderName, de.Name())
		if de.IsDir() {
			readFolder(fullPath)
		} else {
			files = append(files, fullPath)
		}
	}
}

func printVersion() {
	fmt.Println("coldstart", "v1.1.0")
}

func printHelp() {
	fmt.Println("usage:")
	fmt.Println("   ", "coldstart", "[module_name]")
}
