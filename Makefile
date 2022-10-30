build: setup
	@echo "--- Building binary file ---"
	@go build -o ./main server/grpc/main.go

grpc:
	@echo "--- running gRPC server in dev mode ---"
	@go run server/grpc/main.go

tidy:
	@go mod tidy

setup:
	@echo " --- Setup and generate configuration --- "
	@cp config/example/server.yml.example config/server/server.yml

build-docker: build
	@docker build --tag bussiness-services .

protoc-docker:
	@docker container create --name bussiness-services -p 9905:9905 bussiness-services
