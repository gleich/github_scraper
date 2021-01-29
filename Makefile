##########
# Building
##########

build-docker-prod:
	docker build -t mattgleich/github_scraper:latest .
build-docker-dev:
	docker build -f dev.Dockerfile -t mattgleich/github_scraper:test .
build-docker-dev-lint:
	docker build -f dev.lint.Dockerfile -t mattgleich/github_scraper:lint .
build-go:
	go get -v -t -d ./...
	go build -v .
	rm github_scraper

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
lint-hadolint:
	hadolint Dockerfile
	hadolint dev.Dockerfile
	hadolint dev.lint.Dockerfile
lint-in-docker: build-docker-dev-lint
	docker run mattgleich/github_scraper:lint

#########
# Testing
#########

test-go:
	go get -v -t -d ./...
	go test -v ./...
test-in-docker: build-docker-dev
	docker run mattgleich/github_scraper:test

##########
# Grouping
##########

# Testing
local-test: test-go
docker-test: test-in-docker
# Linting
local-lint: lint-golangci lint-hadolint lint-gomod
docker-lint: lint-in-docker
# Build
local-build: build-docker-prod build-docker-dev build-docker-dev-lint
