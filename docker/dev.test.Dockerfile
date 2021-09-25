FROM golang:1.16

# Meta data:
LABEL maintainer="email@mattglei.ch"
LABEL description="🗄 GitHub scraper for my GraphQL API"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing make
RUN apt-get update && apt-get install make=4.3-4.1 -y --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

CMD ["make", "local-test"]
