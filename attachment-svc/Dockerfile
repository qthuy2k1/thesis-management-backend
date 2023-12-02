FROM alpine

WORKDIR /app

COPY ./attachment-svc/.env ./.env

RUN source .env

COPY ./attachment-svc/data/migrations /migrations

COPY ./out/attachment /app/

ENTRYPOINT ./attachment