FROM alpine


WORKDIR /app

COPY ./reporting-stage-svc/.env ./.env

RUN source .env

COPY ./reporting-stage-svc/data/migrations /migrations

COPY ./out /app/

ENTRYPOINT ./reporting-stage