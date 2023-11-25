FROM alpine

WORKDIR /app

COPY ./upload-svc/credentials.json ./credentials.json
COPY ./upload-svc/token.json ./token.json

# RUN source .env

COPY ./out /app/

ENTRYPOINT ./upload