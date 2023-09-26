FROM alpine

WORKDIR /app

COPY ./classroom-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./classroom