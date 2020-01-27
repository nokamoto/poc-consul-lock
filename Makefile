
all:
	go fmt ./...
	go test ./...
	docker-compose build
	docker-compose down
	docker-compose up
