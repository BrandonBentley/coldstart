run:
	@go run main.go

test:
	@go test -coverprofile /dev/null ./... 

coverage:
	@go test -coverprofile cover.out ./... 
	@go tool cover -html=cover.out

gen:
	@go generate ./...

docker:
	@docker build -t coldstart:latest -t coldstart:$$(git rev-parse --short HEAD) .