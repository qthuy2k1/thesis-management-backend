FROM alpine

WORKDIR /app

COPY ./thesis-commitee-svc/.env ./.env

RUN source .env

COPY ./thesis-commitee-svc/data/migrations /migrations

COPY ./out/thesis-commitee /app/

ENTRYPOINT ./thesis-commitee