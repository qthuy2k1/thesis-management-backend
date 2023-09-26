FROM alpine

WORKDIR /app

COPY ./comment-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./comment