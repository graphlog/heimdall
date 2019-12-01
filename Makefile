DOCKER_TARGET?=ci

vendor:
	go get github.com/google/wire/cmd/wire
	go mod vendor

start: wire
	go run cmd/main.go

generate: wire

wire:
	cd pkg && go generate

test:
	go test -v github.com/graphlog/heimdall/pkg/...

test-ci:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v github.com/graphlog/heimdall/pkg/...

lint:
	go get -u golang.org/x/lint/golint
	test -z "$$(golint pkg/... cmd/...)"

cleanup:
	rm -rf ./build
	rm -rf ./vendor

setup: vendor generate

migrate:
	cd db && migrate @ 

default:
	setup
