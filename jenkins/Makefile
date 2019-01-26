SHELL = /bin/bash -o pipefail

include .env

ROOT_DIR = $(CURDIR)
DOCKER_IMAGE_DIR = $(ROOT_DIR)/image

rebuild-php:
	cd $(ROOT_DIR) && docker-compose rm -sf "php" && docker-compose up -d --no-deps --build "php"
