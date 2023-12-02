FROM alpine

WORKDIR /app

COPY ./upload-svc/credentials.json ./credentials.json
COPY ./upload-svc/token.json ./token.json

# RUN source .env

COPY ./out/upload /app/

ENTRYPOINT ./upload