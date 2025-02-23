run:
	@go run main.go

test:
	@go test -coverprofile /dev/null ./... 

coverage:
	@go test -coverprofile cover.out ./... 
	@go tool cover -html=cover.out

gen:
	@go generate ./...