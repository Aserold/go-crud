# Initial stage: download modules
FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

# Intermediate stage: Build the binary
FROM golang:1.22.5-alpine as runner

WORKDIR /app

COPY --from=builder /app /app

# Install CompileDaemon using go install with a specific version
RUN go install github.com/githubnemo/CompileDaemon@latest

ENV config=docker

EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build cmd/server/main.go" --command=./main
