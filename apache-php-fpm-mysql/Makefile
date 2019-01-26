SHELL = /bin/bash -o pipefail

include .env

ROOT_DIR = $(CURDIR)
DOCKER_IMAGE_DIR = $(ROOT_DIR)/image

rebuild-all: rebuild-mysql rebuild-php rebuild-apache

rebuild-mysql:
	cd $(ROOT_DIR) && docker-compose rm -sf "mysql" && docker-compose up -d --no-deps --build "mysql"

rebuild-php:
	cd $(ROOT_DIR) && docker-compose build php
	cd $(ROOT_DIR) && docker-compose rm -sf "php" && docker-compose up -d --no-deps --build "php"

rebuild-apache:
	cd $(ROOT_DIR) && docker-compose build apache
	cd $(ROOT_DIR) && docker-compose rm -sf "apache" && docker-compose up -d --no-deps --build "apache"

