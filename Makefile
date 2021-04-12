init:
	GO111MODULE=on go mod download

build:
	GO111MODULE=on go build

lint:
	GO111MODULE=on golangci-lint run ./app/...

test:
	go test -v ./...

test-coverage:
	go test -cover ./...

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

app-build:
ifeq ($(tag),)
	@echo "Please execute this command with the docker image tag."
	@echo "Usage:"
	@echo "	$$ make app-build tag=<version>"
else
	docker build -f ./Dockerfile -t istsh/gitops-sample-app:${tag} ./
endif

app-push:
ifeq ($(tag),)
	@echo "Please execute this command with the docker image tag."
	@echo "Usage:"
	@echo "	$$ make app-push tag=<version>"
else
	docker push istsh/gitops-sample-app:${tag}
endif

migration-build:
ifeq ($(tag),)
	@echo "Please execute this command with the docker image tag."
	@echo "Usage:"
	@echo "	$$ make migration-build tag=<version>"
else
	docker build -f ./Dockerfile.migration -t istsh/gitops-sample-migration:${tag} ./
endif

migration-push:
ifeq ($(tag),)
	@echo "Please execute this command with the docker image tag."
	@echo "Usage:"
	@echo "	$$ make migration-push tag=<version>"
else
	docker push istsh/gitops-sample-migration:${tag}
endif

create-migration-file:
ifeq ($(name),)
	@echo "Please execute this command with the migration file name."
	@echo "Usage:"
	@echo "	$$ make create-migration-file name=<name>"
else
	migrate create -dir db/migrations/ -ext sql ${name}
endif
