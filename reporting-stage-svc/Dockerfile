FROM alpine


WORKDIR /app

COPY ./reporting-stage-svc/.env ./.env

RUN source .env

COPY ./reporting-stage-svc/data/migrations /migrations

COPY ./out/reporting-stage /app/

ENTRYPOINT ./reporting-stage