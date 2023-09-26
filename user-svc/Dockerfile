FROM alpine

WORKDIR /app

COPY ./user-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./user