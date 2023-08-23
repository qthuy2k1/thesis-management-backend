gen-proto:
	protoc -I./internal/proto \
  	--go_out ./internal/proto --go_opt paths=source_relative \
  	--go-grpc_out ./internal/proto --go-grpc_opt paths=source_relative \
  	--grpc-gateway_out ./internal/proto --grpc-gateway_opt paths=source_relative \
  	--openapiv2_out ./internal/proto \
    --openapiv2_opt logtostderr=true \
	--validate_out="lang=go,paths=source_relative:./internal/proto" \
  	./internal/proto/$(name)/*.proto

run-server:
	go run cmd/classroom-svc/server/server.go

run-client:
	go run cmd/classroom-svc/client/client.go

postgres:
	psql -U postgres -d thesis_management_$(db) -h localhost -p 5432

create_migration:
	migrate create -ext sql -dir data/migrations/ -seq $(filename)

migrate_up:
	migrate -source file://$(PWD)/$(db)-svc/data/migrations/ -database "postgres://postgres:root@localhost:5432/thesis_management_$(db)s?sslmode=disable" up 

migrate_down:
	migrate -source file://$(PWD)/data/migrations/ -database postgres://postgres:root@localhost:5432/thesis_management_$(db)?sslmode=disable down
	
migrate_force:
	migrate -source file://$(PWD)/data/migrations/ -database "postgres://postgres:root@localhost:5432/thesis_management_$(db)?sslmode=disable" force $(ver)

docker_volume_down:
	docker-compose down --volumes


proto-api:
	@echo "--> Generating gRPC clients for API"
	@protoc -I ./api-gw/api/v1 \
		--go_out ./api-gw/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./api-gw/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./api-gw/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./api-gw/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./api-gw/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 api_classroom.proto
	@echo "Done"

proto-classroom:
	@echo "--> Generating gRPC clients for classroom API"
	@protoc -I ./classroom-svc/api/v1 \
		--go_out ./classroom-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./classroom-svc/api/goclient --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./classroom-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./classroom-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./classroom-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 classroom.proto
	@echo "Done"

proto: proto-api proto-classroom

build:
	mkdir -p ./out
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/apigw ./api-gw
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/classroom ./classroom-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/apigw-client ./apigw-client

run: build
	@echo "--> Starting servers"
	docker-compose build
	docker-compose up