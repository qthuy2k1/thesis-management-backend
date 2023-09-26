FROM alpine

WORKDIR /app

COPY ./classroom-waiting-list-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./classroom-waiting-list