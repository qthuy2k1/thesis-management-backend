FROM alpine

WORKDIR /app

COPY ./exercise-svc/.env ./.env

RUN source .env

COPY ./exercise-svc/data/migrations /migrations

COPY ./out/exercise /app/

ENTRYPOINT ./exercise