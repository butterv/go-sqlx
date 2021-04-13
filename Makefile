init:
	go mod download

build:
	go build

lint:
	golangci-lint run ./app/...

test:
	go test -v -race ./...

test-coverage:
	go test -race -cover -coverprofile=coverage.out ./...

view-coverage-report:
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

benchmark:
	go test -bench . -benchmem

generate-pb:
	@for file in `\find proto/v1 -type f -name '*.proto'`; do \
		protoc \
			$$file \
			-I ./proto/v1/ \
			-I $(GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.5.1 \
			--go_out=./app/interface/rpc/v1 \
			--go_opt=paths=source_relative \
			--go-grpc_out=./app/interface/rpc/v1 \
			--go-grpc_opt=paths=source_relative; \
	done
