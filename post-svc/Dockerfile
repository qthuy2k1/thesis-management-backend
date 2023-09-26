FROM alpine

WORKDIR /app

COPY ./post-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./post