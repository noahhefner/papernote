REPO ?= nhefner/notesserver
GITSHA = $(shell git rev-parse --short HEAD)
TAG_COMMIT = $(REPO):$(GITSHA)
TAG_LATEST=$(REPO):latest

all:dev

.PHONY:dev
dev:
	go run .

.PHONY:build
build:
	docker build --no-cache -t $(TAG_LATEST) .

.PHONY:run
run:
	docker run --rm -p 8080:8080 $(TAG_LATEST)

.PHONY:clean
clean:
	@echo "Deleting Docker images with tag $(TAG_LATEST)"
	@docker images -q $(TAG_LATEST) | xargs -r docker rmi -f

.PHONY:publish
publish:
	docker push $(TAG_LATEST)
	@docker tag $(TAG_LATEST) $(TAG_COMMIT)
	docker push $(TAG_COMMIT)