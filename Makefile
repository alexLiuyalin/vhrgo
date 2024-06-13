ifeq ($(OS),Windows_NT)     # is Windows_NT on XP, 2000, 7, Vista, 10...
	GIT_BASH="$(subst cmd\git.exe,bin\bash.exe,$(shell where git))"
else
	GIT_BASH=$(shell which bash)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go get github.com/google/wire/cmd/wire@latest

.PHONY: api
# generate api
api:
	$(GIT_BASH) -c "find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'"

.PHONY: wire
# generate wire
wire:
	$(GIT_BASH) -c "find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'"

.PHONY: config
# generate api
config:
	$(GIT_BASH) -c "find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) config'"

.PHONY: proto
# generate proto
proto:
	$(GIT_BASH) -c "find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) proto'"

.PHONY: ent
# generate ent
ent:
	go run -mod=mod ./pkg/ent/entc/entc.go ./data/schema

.PHONY: common
# generate common
common:
	cd api/common && protoc --proto_path=. \
		--proto_path=../../third_party \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--validate_out=paths=source_relative,lang=go:. \
		*.proto

.PHONY: clear
# generate clear
clear:
	$(shell $(GIT_BASH) -c "find . -type f -name *.pb.go -print0 -o -name *.pb.validate.go -print0 -o -name *.swagger.json -print0 | xargs -0 rm -rf")

.PHONY: help
help:
	@echo $(GIT_BASH)