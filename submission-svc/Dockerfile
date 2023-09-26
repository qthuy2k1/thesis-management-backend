FROM alpine


WORKDIR /app

COPY ./submission-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./submission