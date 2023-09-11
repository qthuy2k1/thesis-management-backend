FROM alpine

ARG image_name
ENV SVC_ENV=$image_name

WORKDIR /app

COPY ./${SVC_ENV}-svc/.env ./.env
RUN source .env

COPY ./out /app/

ENTRYPOINT ./${SVC_ENV}