APP_NAME=template

format:
	gofmt -w .

init:
	air init

dev:
	air

test:
	go test  -tags=test -v ./... 

coverage:
	go test -tags=test  -v ./... -coverprofile cover.out && go tool cover -func cover.out

build:
	go build -o $(APP_NAME) main.go

run:
	go run main.go

.PHONY: format init dev test coverage build run
