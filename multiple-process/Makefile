SHELL = /bin/bash -o pipefail

include .env

docker-rebuild-goapp:
	cd $(DOCKER_DIR) && docker-compose rm -sf "goapp"
	cd $(DOCKER_DIR) && docker-compose up -d --no-deps --build "goapp"

srv-build:
	cd $(SRV_DIR) && CGO_ENABLED=0 go build -o $(SRV_OUT)/goapp -v -x -mod vendor main.go