# run BuildAndCleanup to build the project and clean up the your 

FROM golang:alpine AS builder

WORKDIR /src

COPY . .

RUN apk add  bash

RUN go build -o main . 

LABEL maintainer="Aymen <aymen@example.com>, Abderahman <abderahman@example.com>, Mohamed <mohamed@example.com>"
LABEL version="2.0"
LABEL description="Adding multi-staging support for better performance"

EXPOSE 8080

FROM alpine:latest

WORKDIR /app

COPY --from=builder /src/main /src/shadow.txt /src/standard.txt /src/thinkertoy.txt /app/

COPY --from=builder /src/templates /app/templates

COPY --from=builder /src/static /app/static

CMD ["./main"]