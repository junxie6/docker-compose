PACKAGES := \
	github.com/junxie6/util
DEPENDENCIES := \
	github.com/gin-gonic/gin

all: build silent-test

start: build docker_build_image

build: build_app docker_build_image docker_build_container

build_app:
	go build -o main main.go

docker_build_image:
	docker build --network=host -t exp/main:0.0.0 .

docker_build_container:
	docker-compose rm -sf "$(CONTAINER_NAME)" && docker-compose up -d --no-deps --build "$(CONTAINER_NAME)"

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)
