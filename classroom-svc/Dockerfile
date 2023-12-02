FROM alpine

WORKDIR /app

COPY ./classroom-svc/.env ./.env

RUN source .env

COPY ./classroom-svc/data/migrations /migrations

COPY ./out/classroom /app/

ENTRYPOINT ./classroom