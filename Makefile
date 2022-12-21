NAME=gist-clone

.PHONY: run
run: build wait-for-db
	./build/$(NAME)

.PHONY: run-prod
run-prod: build wait-for-db
	GIST_CLONE_ENV=prod ./build/$(NAME)

.PHONY: build
build:
	go build -o ./build/$(NAME)

.PHONY: wait-for-db
wait-for-db:
	./wait-for-it.sh -s -t 30 127.0.0.1:27017

.PHONY: clean
clean:
	rm -f ./build/$(NAME)

.PHONY: deps
deps:
	go mod download

.PHONY: test
test:
	go test -v ./tests/*
