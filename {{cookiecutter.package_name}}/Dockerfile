# build stage
FROM golang:1.19.0-alpine3.15 AS build-env

# RUN apk --no-cache add build-base

ADD . /src

RUN cd /src && go mod tidy && go build -o goapp

# final stage
FROM alpine

WORKDIR /app
COPY --from=build-env /src/goapp /app/
COPY --from=build-env /src/*.yaml /app/

# Tz ZONEINFO
COPY --from=build-env /usr/local/go/lib/time/zoneinfo.zip /app/
ENV ZONEINFO=/app/zoneinfo.zip

EXPOSE 8000

ENTRYPOINT ./goapp
