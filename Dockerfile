# Dev Image
FROM golang AS build
WORKDIR /build
COPY ubuntu-mirror.go .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /build/ubuntu-mirror

# Production Image
FROM alpine:latest
LABEL maintainer="Tobias Wiese [git@twiese99.de]"

WORKDIR /ubuntu-mirror

RUN apk add --no-cache rsync bash
COPY --from=build /build/ubuntu-mirror ./
ADD *.sh ./

VOLUME /data

ENV MIRROR_TYPE="releases"
ENV INTERVAL="360"
ENV COUNTRY_CODE=""
ENV RSYNC_SOURCE=""

CMD ["/bin/bash", "./ubuntu-mirror.sh"]
