FROM alpine

WORKDIR /app

COPY ./attachment-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./attachment