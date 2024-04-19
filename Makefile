BINARY_NAME=todo-api

.PHONY: run
run: build
	./$(BINARY_NAME)

.PHONY: build
build:
	go build -o $(BINARY_NAME) ./api/

.PHONY: test
test:
	go test ./... -race -cover