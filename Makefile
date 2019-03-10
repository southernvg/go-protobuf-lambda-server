.PHONY: build clean deploy
PROTOC=$(HOME)/bin/protoc

build:
	go mod tidy
	$(PROTOC) -I ping ping.proto --go_out=plugins=grpc:ping
	env GOOS=linux go build -ldflags="-s -w" -o server server.go

buildlocal:
	go mod tidy
	$(PROTOC) -I proto ping.proto --go_out=plugins=grpc:.
	go build -ldflags="-s -w" -o server server.go ping.pb.go

clean:
	rm -rf server ./vendor

deploy: clean build
	sls deploy --verbose
