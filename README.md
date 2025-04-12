# Cold Start
The purpose of this project is to provide a structured starting template for an Go HTTP service that provides the following:
- Basic Project Structure
- Config Setup (via [env](https://github.com/caarlos0/env))
- Dependency Injection (via [fx](https://go.uber.org/fx))
- Mock Generation for Testing (via [mock](https://go.uber.org/mock))
- HTTP Server Setup (via [gin](https://github.com/gin-gonic/gin))


### Installation
``` bash
go install github.com/BrandonBentley/coldstart@v1.2.0
```

### Setup Requirements

Install Go
``` bash
brew install go
```

Add GOBIN to $Path in `.zprofile`
``` bash
printf "\nPATH=\$PATH:\$(go env GOPATH)/bin\n" >> ~/.zprofile
```

Install Mockgen (For mock generation)
``` bash
go install go.uber.org/mock/mockgen@latest
```
