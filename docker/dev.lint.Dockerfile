FROM golangci/golangci-lint:v1.38

# Meta data
LABEL maintainer="email@mattglei.ch"
LABEL description="🗄  GitHub scraper for my GraphQL API"

# Copying over files
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing hadolint:
WORKDIR /usr/bin
RUN curl -sL -o hadolint "https://github.com/hadolint/hadolint/releases/download/v1.17.6/hadolint-$(uname -s)-$(uname -m)" \
    && chmod 700 hadolint

WORKDIR /usr/src/app

CMD ["make", "local-lint"]
