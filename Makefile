postgres:
	docker exec -it $(service)-db psql -U postgres -d thesis_management_$(db)s -h thesis-management-backend-$(service)-db -p 5432

create_migration:
	migrate create -ext sql -dir $(name)-svc/data/migrations/ -seq $(filename)

migrate_up:
	docker run --rm -v $(PWD)/$(service)-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@thesis-management-backend-$(service)-db-service:5432/thesis_management_$(db)s?sslmode=disable" up 

migrate_down:
	docker run --rm -v $(PWD)/$(service)-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@thesis-management-backend-$(service)-db-service:5432/thesis_management_$(db)s?sslmode=disable" down $(ver)
	
migrate_force:
	docker run --rm -v $(PWD)/$(service)-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@thesis-management-backend-$(service)-db-service:5432/thesis_management_$(db)s?sslmode=disable" force $(ver) 

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
		 api_classroom.proto api_post.proto api_exercise.proto api_reporting_stage.proto api_submission.proto api_user.proto api_waiting_list.proto api_comment.proto api_attachment.proto api_topic.proto api_authorization.proto api_member.proto api_thesis_commitee.proto api_room.proto api_student_def.proto api_schedule.proto api_notification.proto api_point.proto
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

proto-reporting-stage:
	@echo "--> Generating gRPC clients for reporting stage API"
	@protoc -I ./reporting-stage-svc/api/v1 \
		--go_out ./reporting-stage-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./reporting-stage-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./reporting-stage-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./reporting-stage-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./reporting-stage-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 reporting_stage.proto
	@echo "Done"

proto-submission:
	@echo "--> Generating gRPC clients for submission API"
	@protoc -I ./submission-svc/api/v1 \
		--go_out ./submission-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./submission-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./submission-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./submission-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./submission-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 submission.proto
	@echo "Done"

proto-user:
	@echo "--> Generating gRPC clients for user API"
	@protoc -I ./user-svc/api/v1 \
		--go_out ./user-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./user-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./user-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./user-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./user-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 user.proto
	@echo "Done"

proto-waiting-list:
	@echo "--> Generating gRPC clients for classroom-waiting-list API"
	@protoc -I ./classroom-waiting-list-svc/api/v1 \
		--go_out ./classroom-waiting-list-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./classroom-waiting-list-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./classroom-waiting-list-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./classroom-waiting-list-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./classroom-waiting-list-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 waiting_list.proto
	@echo "Done"

proto-redis:
	@echo "--> Generating gRPC clients for redis API"
	@protoc -I ./redis-svc/api/v1 \
		--go_out ./redis-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./redis-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./redis-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./redis-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./redis-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 redis.proto
	@echo "Done"

proto-comment:
	@echo "--> Generating gRPC clients for comment API"
	@protoc -I ./comment-svc/api/v1 \
		--go_out ./comment-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./comment-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./comment-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./comment-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./comment-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 comment.proto
	@echo "Done"

proto-attachment:
	@echo "--> Generating gRPC clients for attachment API"
	@protoc -I ./attachment-svc/api/v1 \
		--go_out ./attachment-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./attachment-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./attachment-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./attachment-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./attachment-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 attachment.proto
	@echo "Done"

proto-authorization:
	@echo "--> Generating gRPC clients for authorization API"
	@protoc -I ./authorization-svc/api/v1 \
		--go_out ./authorization-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./authorization-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./authorization-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./authorization-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./authorization-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 authorization.proto
	@echo "Done"

proto-topic:
	@echo "--> Generating gRPC clients for topic API"
	@protoc -I ./topic-svc/api/v1 \
		--go_out ./topic-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./topic-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./topic-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./topic-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./topic-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 topic.proto
	@echo "Done"

proto-commitee:
	@echo "--> Generating gRPC clients for thesis-commitee API"
	@protoc -I ./thesis-commitee-svc/api/v1 \
		--go_out ./thesis-commitee-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./thesis-commitee-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./thesis-commitee-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./thesis-commitee-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./thesis-commitee-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 thesis_commitee.proto schedule_commitee.proto
	@echo "Done"


proto-sche:
	@echo "--> Generating gRPC clients for schedule API"
	@protoc -I ./schedule-svc/src/proto \
		--go_out ./schedule-svc/api/goclient/v1 --go_opt paths=source_relative \
	  	--go-grpc_out ./schedule-svc/api/goclient/v1 --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./schedule-svc/api/goclient/v1 \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
  		--openapiv2_out ./schedule-svc/api/goclient/v1 \
    	--openapiv2_opt logtostderr=true \
		--validate_out="lang=go,paths=source_relative:./schedule-svc/api/goclient/v1" \
		--experimental_allow_proto3_optional \
		 schedule.proto
	@echo "Done"





proto: proto-api proto-classroom proto-post proto-exercise proto-reporting-stage proto-submission proto-user proto-waiting-list proto-redis proto-comment proto-attachment proto-topic proto-authorization proto-commitee proto-sche

clean:
	rm -rf ./out

build:
	mkdir -p ./out
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/classroom ./classroom-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/apigw ./api-gw
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/apigw-client ./apigw-client
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/post ./post-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/exercise ./exercise-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/user ./user-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/reporting-stage ./reporting-stage-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/submission ./submission-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/classroom-waiting-list ./classroom-waiting-list-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/comment ./comment-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/attachment ./attachment-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/topic ./topic-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/authorization ./authorization-svc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/thesis-commitee ./thesis-commitee-svc



build_and_run: clean build
	@echo "--> Starting servers"
	docker-compose build
	docker-compose up --remove-orphans

down:
	docker-compose down
	@echo "--> Server stopped"

docker-tag:
	# APP
	docker tag qthuy2k1/thesis-management-backend:$(tag) qthuy2k1/thesis-management-backend:$(tag)
	docker tag qthuy2k1/thesis-management-backend-apigw-client:$(tag) qthuy2k1/thesis-management-backend-apigw-client:$(tag)
	docker tag qthuy2k1/thesis-management-backend-classroom:$(tag) qthuy2k1/thesis-management-backend-classroom:$(tag)
	docker tag qthuy2k1/thesis-management-backend-post:$(tag) qthuy2k1/thesis-management-backend-post:$(tag)
	docker tag qthuy2k1/thesis-management-backend-exercise:$(tag) qthuy2k1/thesis-management-backend-exercise:$(tag)
	docker tag qthuy2k1/thesis-management-backend-user:$(tag) qthuy2k1/thesis-management-backend-user:$(tag)
	docker tag qthuy2k1/thesis-management-backend-reporting-stage:$(tag) qthuy2k1/thesis-management-backend-reporting-stage:$(tag)
	docker tag qthuy2k1/thesis-management-backend-submission:$(tag) qthuy2k1/thesis-management-backend-submisson:$(tag)
	docker tag qthuy2k1/thesis-management-backend-classroom-waiting-list:$(tag) qthuy2k1/thesis-management-backend-classroom-waiting-list:$(tag)
	docker tag qthuy2k1/thesis-management-backend-comment:$(tag) qthuy2k1/thesis-management-backend-comment:$(tag)
	docker tag qthuy2k1/thesis-management-backend-attachment:$(tag) qthuy2k1/thesis-management-backend-attachment:$(tag)
	docker tag qthuy2k1/thesis-management-backend-topic:$(tag) qthuy2k1/thesis-management-backend-topic:$(tag)
	docker tag qthuy2k1/thesis-management-backend-authorization:$(tag) qthuy2k1/thesis-management-backend-authorization:$(tag)
	docker tag qthuy2k1/thesis-management-backend-thesis-commitee:$(tag) qthuy2k1/thesis-management-backend-thesis-commitee:$(tag)
	
	# DB
	docker tag postgres qthuy2k1/thesis-management-backend-classroom-db:$(tag)
	docker tag postgres qthuy2k1/thesis-management-backend-post-db:$(tag)
	docker tag postgres qthuy2k1/thesis-management-backend-exercise-db:$(tag)
	docker tag postgres qthuy2k1/thesis-management-backend-user-db:$(tag)
	docker tag postgres qthuy2k1/thesis-management-backend-reporting-stage-db:$(tag)
	docker tag postgres qthuy2k1/thesis-management-backend-submission-db:$(tag)
	docker tag postgres qthuy2k1/thesis-management-backend-classroom-waiting-list-db:$(tag)
	docker tag postgres qthuy2k1/thesis-management-backend-comment-db:$(tag)
	docker tag postgres qthuy2k1/thesis-management-backend-attachment-db:$(tag)
	docker tag postgres qthuy2k1/thesis-management-backend-topic-db:$(tag)
	docker tag postgres qthuy2k1/thesis-management-backend-thesis-commitee-db:$(tag)


docker-google-cloud-tag:
	# APP
	docker tag qthuy2k1/thesis-management-backend:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend:$(tag)
	docker tag qthuy2k1/thesis-management-backend-apigw-client:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-apigw-client:$(tag)
	docker tag qthuy2k1/thesis-management-backend-classroom:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-classroom:$(tag)
	docker tag qthuy2k1/thesis-management-backend-post:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-post:$(tag)
	docker tag qthuy2k1/thesis-management-backend-exercise:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-exercise:$(tag)
	docker tag qthuy2k1/thesis-management-backend-reporting-stage:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-reporting-stage:$(tag)
	docker tag qthuy2k1/thesis-management-backend-submission:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-submission:$(tag)
	docker tag qthuy2k1/thesis-management-backend-classroom-waiting-list:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-waiting-list:$(tag)
	docker tag qthuy2k1/thesis-management-backend-comment:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-comment:$(tag)
	docker tag qthuy2k1/thesis-management-backend-attachment:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-attachment:$(tag)
	docker tag qthuy2k1/thesis-management-backend-topic:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-topic:$(tag)
	docker tag qthuy2k1/thesis-management-backend-authorization:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-authorization:$(tag)
	docker tag qthuy2k1/thesis-management-backend-user:$(tag) asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-user:$(tag)
	
	# DB
	docker tag postgres asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-classroom-db:$(tag)
	docker tag postgres asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-post-db:$(tag)
	docker tag postgres asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-exercise-db:$(tag)
	docker tag postgres asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-user-db:$(tag)
	docker tag postgres asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-reporting-stage-db:$(tag)
	docker tag postgres asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-submission-db:$(tag)
	docker tag postgres asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-waiting-list-db:$(tag)
	docker tag postgres asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-comment-db:$(tag)
	docker tag postgres asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-attachment-db:$(tag)
	docker tag postgres asia-southeast1-docker.pkg.dev/${PROJECT_ID}/thesis-course-registration/thesis-management-backend-topic-db:$(tag)

docker-push:
	# APP
	docker push qthuy2k1/thesis-management-backend:latest
	docker push qthuy2k1/thesis-management-backend-apigw-client:latest
	docker push qthuy2k1/thesis-management-backend-classroom:latest
	docker push qthuy2k1/thesis-management-backend-post:latest
	docker push qthuy2k1/thesis-management-backend-exercise:latest
	docker push qthuy2k1/thesis-management-backend-user:latest
	docker push qthuy2k1/thesis-management-backend-reporting-stage:latest
	docker push qthuy2k1/thesis-management-backend-submission:latest
	docker push qthuy2k1/thesis-management-backend-classroom-waiting-list:latest
	docker push qthuy2k1/thesis-management-backend-comment:latest
	docker push qthuy2k1/thesis-management-backend-attachment:latest
	docker push qthuy2k1/thesis-management-backend-topic:latest
	docker push qthuy2k1/thesis-management-backend-authorization:latest
	docker push qthuy2k1/thesis-management-backend-thesis-commitee:latest

	# DB
	docker push qthuy2k1/thesis-management-backend-classroom-db:latest
	docker push qthuy2k1/thesis-management-backend-post-db:latest
	docker push qthuy2k1/thesis-management-backend-exercise-db:latest
	docker push qthuy2k1/thesis-management-backend-user-db:latest
	docker push qthuy2k1/thesis-management-backend-reporting-stage-db:latest
	docker push qthuy2k1/thesis-management-backend-submission-db:latest
	docker push qthuy2k1/thesis-management-backend-classroom-waiting-list-db:latest
	docker push qthuy2k1/thesis-management-backend-comment-db:latest
	docker push qthuy2k1/thesis-management-backend-attachment-db:latest
	docker push qthuy2k1/thesis-management-backend-topic-db:latest
	docker push qthuy2k1/thesis-management-backend-thesis-commitee-db:latest


migrate_all_up:
	docker run --rm -v $(PWD)/classroom-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@classroom-db:5432/thesis_management_classrooms?sslmode=disable" up

	docker run --rm -v $(PWD)/post-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@post-db:5432/thesis_management_posts?sslmode=disable" up

	docker run --rm -v $(PWD)/exercise-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@exercise-db:5432/thesis_management_exercises?sslmode=disable" up

	docker run --rm -v $(PWD)/user-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@user-db:5432/thesis_management_users?sslmode=disable" up

	docker run --rm -v $(PWD)/reporting-stage-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@reporting-stage-db:5432/thesis_management_reporting_stages?sslmode=disable" up

	docker run --rm -v $(PWD)/submission-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@submission-db:5432/thesis_management_submissions?sslmode=disable" up

	docker run --rm -v $(PWD)/classroom-waiting-list-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@classroom-waiting-list-db:5432/thesis_management_waiting_lists?sslmode=disable" up

	docker run --rm -v $(PWD)/comment-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@comment-db:5432/thesis_management_comments?sslmode=disable" up

	docker run --rm -v $(PWD)/attachment-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@attachment-db:5432/thesis_management_attachments?sslmode=disable" up

	docker run --rm -v $(PWD)/topic-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@topic-db:5432/thesis_management_topics?sslmode=disable" up

	docker run --rm -v $(PWD)/thesis-commitee-svc/data/migrations/:/migrations --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@thesis-commitee-db:5432/thesis_management_thesis_commitees?sslmode=disable" up

partner_migrate_all_up:
	docker run --rm -v "D:/Web Dev/thesis-management-backend/classroom-svc/data/migrations/:/migrations" --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@classroom-db:5432/thesis_management_classrooms?sslmode=disable" up

	docker run --rm -v "D:/Web Dev/thesis-management-backend/post-svc/data/migrations/:/migrations" --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@post-db:5432/thesis_management_posts?sslmode=disable" up

	docker run --rm -v "D:/Web Dev/thesis-management-backend/exercise-svc/data/migrations/:/migrations" --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@exercise-db:5432/thesis_management_exercises?sslmode=disable" up

	docker run --rm -v "D:/Web Dev/thesis-management-backend/user-svc/data/migrations/:/migrations" --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@user-db:5432/thesis_management_users?sslmode=disable" up

	docker run --rm -v "D:/Web Dev/thesis-management-backend/reporting-stage-svc/data/migrations/:/migrations" --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@reporting-stage-db:5432/thesis_management_reporting_stages?sslmode=disable" up

	docker run --rm -v "D:/Web Dev/thesis-management-backend  /submission-svc/data/migrations/:/migrations" --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@submission-db:5432/thesis_management_submissions?sslmode=disable" up

	docker run --rm -v "D:/Web Dev/thesis-management-backend/classroom-waiting-list-svc/data/migrations/:/migrations" --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@classroom-waiting-list-db:5432/thesis_management_waiting_lists?sslmode=disable" up

	docker run --rm -v "D:/Web Dev/thesis-management-backend/comment-svc/data/migrations/:/migrations" --network thesis-management-backend_mynet migrate/migrate -path=/migrations/ -database "postgres://postgres:root@comment-db:5432/thesis_management_comments?sslmode=disable" up


gcloud-ssh:
	gcloud compute ssh --project=thesis-course-registration --zone=asia-southeast1-b instance-1

build_and_push_single_image:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/$(name) ./$(folder)
	docker build -f $(name)/Dockerfile -t qthuy2k1/thesis-management-backend-$(name) .
	docker push qthuy2k1/thesis-management-backend-$(name):$(tag)

docker-login:
	docker login --username qthuy2k1 --password-stdin

build_and_run_image:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/$(name) ./$(folder)
	docker build -f $(folder)/Dockerfile -t qthuy2k1/thesis-management-backend-$(name)s .
	docker compose up

docker-pull-db:
	# DB
	docker pull qthuy2k1/thesis-management-backend-classroom-db:$(tag)
	docker pull qthuy2k1/thesis-management-backend-post-db:$(tag)
	docker pull qthuy2k1/thesis-management-backend-exercise-db:$(tag)
	docker pull qthuy2k1/thesis-management-backend-user-db:$(tag)
	docker pull qthuy2k1/thesis-management-backend-reporting-stage-db:$(tag)
	docker pull qthuy2k1/thesis-management-backend-submission-db:$(tag)
	docker pull qthuy2k1/thesis-management-backend-classroom-waiting-list-db:$(tag)
	docker pull qthuy2k1/thesis-management-backend-comment-db:$(tag)
	docker pull qthuy2k1/thesis-management-backend-attachment-db:$(tag)
	docker pull qthuy2k1/thesis-management-backend-topic-db:$(tag)

kuber-exec:
	kubectl exec -it thesis-management-backend$(name) -n thesis-management-backend -- bash

kuber-delete:
	kubectl delete -f $(file).yaml --cascade=orphan

kuber-all: clean build
	docker compose build

	# ===================================== PUSHING TO DOCKER =====================================
	# APP
	docker push qthuy2k1/thesis-management-backend:latest
	docker push qthuy2k1/thesis-management-backend-apigw-client:latest
	docker push qthuy2k1/thesis-management-backend-classroom:latest
	docker push qthuy2k1/thesis-management-backend-post:latest
	docker push qthuy2k1/thesis-management-backend-exercise:latest
	docker push qthuy2k1/thesis-management-backend-user:latest
	docker push qthuy2k1/thesis-management-backend-reporting-stage:latest
	docker push qthuy2k1/thesis-management-backend-submission:latest
	docker push qthuy2k1/thesis-management-backend-classroom-waiting-list:latest
	docker push qthuy2k1/thesis-management-backend-comment:latest
	docker push qthuy2k1/thesis-management-backend-attachment:latest
	docker push qthuy2k1/thesis-management-backend-topic:latest
	docker push qthuy2k1/thesis-management-backend-authorization:latest
	docker push qthuy2k1/thesis-management-backend-thesis-commitee:latest


	# ===================================== DELETING KUBERNETES =================================
	kubectl delete -f kubernetes/attachment-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/classroom-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/classroom-waiting-list-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/comment-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/exercise-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/post-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/reporting-stage-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/submission-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/topic-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/user-deployment.yaml --namespace thesis-management-backend

	kubectl delete -f kubernetes/attachment-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/classroom-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/classroom-waiting-list-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/comment-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/exercise-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/post-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/reporting-stage-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/submission-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/topic-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/user-db-deployment.yaml --namespace thesis-management-backend

	kubectl delete -f kubernetes/api-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/apigw-client-deployment.yaml --namespace thesis-management-backend


	# ===================================== APPLYING KUBERNETES =================================
	kubectl apply -f kubernetes/attachment-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/classroom-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/classroom-waiting-list-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/comment-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/exercise-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/post-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/reporting-stage-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/submission-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/topic-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/user-deployment.yaml --namespace thesis-management-backend

	kubectl apply -f kubernetes/attachment-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/classroom-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/classroom-waiting-list-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/comment-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/exercise-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/post-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/reporting-stage-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/submission-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/topic-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/user-db-deployment.yaml --namespace thesis-management-backend

	kubectl apply -f kubernetes/api-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/apigw-client-deployment.yaml --namespace thesis-management-backend

kuber-serve-gw-client:
	minikube service thesis-management-backend-apigw-client-service --url -n thesis-management-backend

kuber-apply:
	kubectl apply -f kubernetes/attachment-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/classroom-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/classroom-waiting-list-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/comment-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/exercise-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/post-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/reporting-stage-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/submission-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/topic-db-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/user-db-deployment.yaml --namespace thesis-management-backend

	kubectl apply -f kubernetes/attachment-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/classroom-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/classroom-waiting-list-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/comment-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/exercise-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/post-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/reporting-stage-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/submission-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/topic-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/user-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/schedule-deployment.yaml --namespace thesis-management-backend

	kubectl apply -f kubernetes/api-deployment.yaml --namespace thesis-management-backend
	kubectl apply -f kubernetes/apigw-client-deployment.yaml --namespace thesis-management-backend


kuber-del:
	kubectl delete -f kubernetes/attachment-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/classroom-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/classroom-waiting-list-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/comment-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/exercise-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/post-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/reporting-stage-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/submission-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/topic-db-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/user-db-deployment.yaml --namespace thesis-management-backend

	kubectl delete -f kubernetes/attachment-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/classroom-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/classroom-waiting-list-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/comment-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/exercise-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/post-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/reporting-stage-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/submission-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/topic-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/user-deployment.yaml --namespace thesis-management-backend

	kubectl delete -f kubernetes/api-deployment.yaml --namespace thesis-management-backend
	kubectl delete -f kubernetes/apigw-client-deployment.yaml --namespace thesis-management-backend


rebuild-kuber:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/$(file-out) ./$(folder-svc)
	docker build -f $(folder-svc)/Dockerfile -t qthuy2k1/thesis-management-backend$(svc) .
	docker push qthuy2k1/thesis-management-backend$(svc):latest
	kubectl delete -f kubernetes/$(kuber-name)-deployment.yaml
	kubectl apply -f kubernetes/$(kuber-name)-deployment.yaml

build-a-svc:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/$(name) ./$(folder)
	docker build -f $(folder)/Dockerfile -t qthuy2k1/thesis-management-backend$(svc) .
	docker compose up