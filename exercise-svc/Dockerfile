FROM alpine

WORKDIR /app

COPY ./exercise-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./exercise