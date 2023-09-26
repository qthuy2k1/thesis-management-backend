FROM alpine


WORKDIR /app

COPY ./reporting-stage-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./reporting-stage