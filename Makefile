run:
	@go run cmd/main.go

test:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

docker-up:
	@sudo docker-compose up

docker-down:
	@sudo docker-compose down