FROM alpine

WORKDIR /app

COPY ./post-svc/.env ./.env

RUN source .env

COPY ./post-svc/data/migrations /migrations

COPY ./out/post /app/

ENTRYPOINT ./post