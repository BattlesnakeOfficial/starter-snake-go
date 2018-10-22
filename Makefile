get:
	@echo ">> Getting any missing dependencies.."
	go get -t ./...
.PHONY: get

install:
	go install github.com/battlesnakeio/starter-snake-go
.PHONY: install

run: install
	starter-snake-go server
.PHONY: run

test:
	go test ./...
.PHONY: test

fmt:
	@echo ">> Running Gofmt.."
	gofmt -l -s -w .
