##########
# Building
##########

build-docker-prod:
	docker build -f docker/prod.Dockerfile -t mattgleich/github_scraper:latest .
build-docker-dev:
	docker build -f docker/dev.test.Dockerfile -t mattgleich/github_scraper:test .
build-docker-dev-lint:
	docker build -f docker/dev.lint.Dockerfile -t mattgleich/github_scraper:lint .
build-go:
	go get -v -t -d ./...
	go build -v ./cmd/github_scraper
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
	hadolint docker/prod.Dockerfile
	hadolint docker/dev.test.Dockerfile
	hadolint docker/dev.lint.Dockerfile
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

###################
# Local Development
###################

dev-start:
	docker-compose up -d postgres
	docker-compose up github_scraper

dev-reset:
	docker-compose down
	docker system prune -af

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
# Local development
dev-reboot: dev-reset dev-start
