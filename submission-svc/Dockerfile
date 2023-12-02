FROM alpine


WORKDIR /app

COPY ./submission-svc/.env ./.env

RUN source .env

COPY ./submission-svc/data/migrations /migrations

COPY ./out/submission /app/

ENTRYPOINT ./submission