
ENV ?= dev

clean:
	rm -rf bin/*

fmt:
	go fmt ./cmd/* && \
	go fmt ./pkg/*

wire:
	go run github.com/google/wire/cmd/wire cmd/server/wire.go

run:
	go run -tags dev cmd/server/main.go cmd/server/wire_gen.go

build: 
	go build \
	-tags $(ENV) \
	-o bin/server \
	cmd/server/main.go cmd/server/wire_gen.go

up:
	docker-compose build base-service-dev && \
	docker-compose up -d base-service-dev

down:
	docker-compose down
