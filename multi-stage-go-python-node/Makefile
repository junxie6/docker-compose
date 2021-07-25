SHELL = /bin/bash -o pipefail

include .env

docker-rebuild-$(APP_NAME):
	cd $(DOCKER_DIR) && docker-compose rm -sf "$(APP_NAME)"
	cd $(DOCKER_DIR) && docker-compose up -d --no-deps --build "$(APP_NAME)"

docker-exec-$(APP_NAME):
	cd $(DOCKER_DIR) && docker-compose --env-file .env -f docker-compose.yml exec -e TERM=$$TERM -e LINES=$$LINES -e COLUMNS=$$COLUMNS "$(APP_NAME)" bash -l

