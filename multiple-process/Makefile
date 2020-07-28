SHELL = /bin/bash -o pipefail

include .env

docker-rebuild-goapp:
	cd $(DOCKER_DIR) && docker-compose rm -sf "goapp"
	cd $(DOCKER_DIR) && docker-compose up -d --no-deps --build "goapp"

docker-exec-goapp:
	cd $(DOCKER_DIR) && docker-compose --env-file .env -f docker-compose.yml exec -e TERM=$$TERM -e LINES=$$LINES -e COLUMNS=$$COLUMNS "goapp" sh

docker-log-goapp:
	cd $(DOCKER_DIR) && docker-compose logs "goapp"

srv-build:
	cd $(SRV_DIR) && CGO_ENABLED=0 go build -o $(SRV_OUT)/goapp -v -x -mod vendor main.go
