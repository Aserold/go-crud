run:
	go run ./cmd/server/main.go

build:
	go build ./cmd/server/main.go

test:
	go test ./...

develop:
	echo "Starting docker environment"
	docker-compose -f docker-compose.dev.yml up --build
	# docker compose -f docker-compose.dev.yml up --build
