CURRENT_DIR=$(shell pwd)

.DEFAULT_GOAL = build

# init env
.PHONY: env
env:
	export $(grep -v '^#' .env | xargs)

# build for current os
.PHONY: build
build:
	go build -o mh-access-log-server  -ldflags="-s -w" cmd/main.go

# build for linux amd64
.PHONY: build-linux
build-linux:
	GOOS="linux" GOARCH="amd64" go build -o mh-access-log-server -ldflags="-s -w" cmd/main.go

# run service
.PHONY: run
run:
	go run cmd/main.go

# migrate	
.PHONY: migrate
migrate:
	migrate -source file://migrations -database postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable up

# proto	
.PHONY: proto
proto:
	./scripts/gen-proto.sh
	