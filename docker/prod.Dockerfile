FROM golang:1.16-alpine AS builder

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
COPY --from=builder /usr/src/app/app .

CMD ["./app"]
