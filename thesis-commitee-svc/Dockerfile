FROM alpine

WORKDIR /app

COPY ./thesis-commitee-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./thesis-commitee