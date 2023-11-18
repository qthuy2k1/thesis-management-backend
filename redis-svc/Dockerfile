FROM alpine

WORKDIR /app

COPY ./redis-svc/.env ./.env

RUN source .env

COPY ./out /app/

ENTRYPOINT ./redis