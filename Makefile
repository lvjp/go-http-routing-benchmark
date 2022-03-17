all: lint test

lint:
	@docker run \
		--rm \
		--mount "type=bind,src=${PWD},dst=/app,ro" \
		--workdir /app \
		golangci/golangci-lint:v1.44.2 \
		golangci-lint run

test:
	go test .

bench:
	go test -bench .

.PHONY: all bench lint test
