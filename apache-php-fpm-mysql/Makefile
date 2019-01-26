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
	cd $(ROOT_DIR) && docker-compose rm -sf "apache" && docker-compose up -d --no-deps --build "apache"
	cd $(ROOT_DIR) && docker cp resources/apache/httpd.conf apache-php-fpm-mysql_apache_1:/usr/local/apache2/conf/httpd.conf
	cd $(ROOT_DIR) && docker-compose exec apache chown root:www-data /usr/local/apache2/conf/httpd.conf
	cd $(ROOT_DIR) && docker cp resources/apache/php7.2-fpm.conf apache-php-fpm-mysql_apache_1:/usr/local/apache2/conf/extra/php7.2-fpm.conf
	cd $(ROOT_DIR) && docker-compose exec apache chown root:www-data /usr/local/apache2/conf/extra/php7.2-fpm.conf
	cd $(ROOT_DIR) && docker cp resources/apache/httpd-vhosts.conf apache-php-fpm-mysql_apache_1:/usr/local/apache2/conf/extra/httpd-vhosts.conf
	cd $(ROOT_DIR) && docker-compose exec apache chown root:www-data /usr/local/apache2/conf/extra/httpd-vhosts.conf
	cd $(ROOT_DIR) && docker-compose restart "apache"

