FROM alpine

WORKDIR /app

COPY ./comment-svc/.env ./.env

RUN source .env

COPY ./comment-svc/data/migrations /migrations

COPY ./out/comment /app/

ENTRYPOINT ./comment