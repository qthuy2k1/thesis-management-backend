FROM alpine

WORKDIR /app

COPY ./classroom-waiting-list-svc/.env ./.env

RUN source .env

COPY ./classroom-waiting-list-svc/data/migrations /migrations

COPY ./out/classroom-waiting-list /app/

ENTRYPOINT ./classroom-waiting-list