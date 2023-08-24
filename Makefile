postgres:
	docker exec -it $(db)-db psql -U postgres -d thesis_management_$(db)s -h $(db)-db -p $(port)

create_migration:
	migrate create -ext sql -dir data/migrations/ -seq $(filename)

migrate_up:
	docker run --rm -v $(PWD)/$(db)-svc/data/migrations/:/migrations --network api_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@$(db)-db:5432/thesis_management_$(db)s?sslmode=disable" up 

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
		 api_classroom.proto api_post.proto api_exercise.proto
	@echo "Done"

proto-classroom:
	@echo "--> Generating gRPC clients for classroom API"
	@protoc -I ./classroom-svc/api/v1 \
		--go_out ./classroom-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./classroom-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
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

proto-post:
	@echo "--> Generating gRPC clients for post API"
	@protoc -I ./post-svc/api/v1 \
		--go_out ./post-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./post-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./post-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./post-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./post-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 post.proto
	@echo "Done"

proto-exercise:
	@echo "--> Generating gRPC clients for exercise API"
	@protoc -I ./exercise-svc/api/v1 \
		--go_out ./exercise-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./exercise-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./exercise-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./exercise-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./exercise-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 exercise.proto
	@echo "Done"

proto: proto-api proto-classroom proto-post proto-exercise

clean:
	rm -rf ./out

build:
	mkdir -p ./out
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/apigw ./api-gw
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/apigw-client ./apigw-client
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/classroom ./classroom-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/post ./post-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/exercise ./exercise-svc

run: clean build
	@echo "--> Starting servers"
	docker-compose build
	docker-compose up

down:
	docker-compose down
	@echo "--> Server stopped"

docker-tag:
	docker tag qthuy2k1/thesis-management-backend:latest qthuy2k1/thesis-management-backend:latest
	docker tag qthuy2k1/thesis-management-backend-apigw-client:latest qthuy2k1/thesis-management-backend-apigw-client:latest
	docker tag qthuy2k1/thesis-management-backend-classroom:latest qthuy2k1/thesis-management-backend-classroom:latest
	docker tag qthuy2k1/thesis-management-backend-post:latest qthuy2k1/thesis-management-backend-post:latest
	docker tag qthuy2k1/thesis-management-backend-exercise:latest qthuy2k1/thesis-management-backend-exercise:latest

docker-push:
	docker push qthuy2k1/thesis-management-backend:latest
	docker push qthuy2k1/thesis-management-backend-apigw-client:latest
	docker push qthuy2k1/thesis-management-backend-classroom:latest
	docker push qthuy2k1/thesis-management-backend-post:latest
	docker push qthuy2k1/thesis-management-backend-exercise:latest