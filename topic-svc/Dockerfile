FROM alpine

WORKDIR /app

COPY ./topic-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./topic