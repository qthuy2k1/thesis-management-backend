FROM alpine

WORKDIR /app

COPY ./topic-svc/.env ./.env

RUN source .env

COPY ./topic-svc/data/migrations /migrations

COPY ./out /app/

ENTRYPOINT ./topic