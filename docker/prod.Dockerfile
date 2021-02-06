FROM golang:1.15-alpine3.12 AS builder

# Meta data:
LABEL maintainer="email@mattglei.ch"
LABEL description="🗄 GitHub scraper for my GraphQL API"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing dependencies/
RUN go get -v -t -d ./...

# Build the app
RUN go build -o app ./cmd/github_scraper

# hadolint ignore=DL3006,DL3007
FROM alpine:latest

# Set timezone to EST
ENV TZ=America/New_York
# hadolint ignore=DL3018
RUN apk update && apk add tzdata --no-cache
RUN echo "America/New_York" > /etc/timezone

COPY --from=builder /usr/src/app/app .

CMD ["./app"]
