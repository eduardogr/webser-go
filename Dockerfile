FROM golang:1.24.1-alpine3.21

ARG BUILD_ENV=dev

# Creating working directory
RUN mkdir /app
ADD . /app
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

## Add this go mod download command to pull in any dependencies
RUN go mod download

## Our project will now successfully build with the necessary go libraries included.
RUN go build -tags ${BUILD_ENV} -o bin/server cmd/server/main.go cmd/server/wire_gen.go

## Our start command which kicks off
## our newly created binary executable
CMD ["/app/bin/server"]
