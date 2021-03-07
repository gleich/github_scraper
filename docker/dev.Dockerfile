FROM golang:1.16-alpine

# Meta data
LABEL maintainer="email@mattglei.ch"
LABEL description="🗄 GitHub scraper for my GraphQL API"

# Copying over files
WORKDIR /usr/src/app
COPY . /usr/src/app

# Installing dependencies/
RUN go get -v -t -d ./...

# Setting up air
RUN go get -u github.com/cosmtrek/air
ENV air_wd /usr/src/app/

CMD [ "air", "-c", ".air.toml" ]
