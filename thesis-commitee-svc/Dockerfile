FROM alpine

WORKDIR /app

COPY ./thesis-commitee-svc/.env ./.env

RUN source .env

COPY ./thesis-commitee-svc/data/migrations /migrations

COPY ./out /app/

ENTRYPOINT ./thesis-commitee