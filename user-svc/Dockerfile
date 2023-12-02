FROM alpine

WORKDIR /app

COPY ./user-svc/.env ./.env

RUN source .env

COPY ./user-svc/data/migrations /migrations

COPY ./out/user /app/

ENTRYPOINT ./user