SHELL = /bin/bash -o pipefail

ROOT_DIR = $(CURDIR)

rebuild-all: rebuild-mysql rebuild-php rebuild-apache

rebuild-mysql:
	cd $(ROOT_DIR) && docker-compose rm -sf "mysql" && docker-compose up -d --no-deps --build "mysql"

rebuild-php:
	cd $(ROOT_DIR) && docker-compose rm -sf "php" && docker-compose up -d --no-deps --build "php"
	cd $(ROOT_DIR) && docker cp resources/www.conf apache-php-fpm-mysql_php_1:/usr/local/etc/php-fpm.d/www.conf
	cd $(ROOT_DIR) && docker-compose restart "php"

rebuild-apache:
	cd $(ROOT_DIR) && docker-compose rm -sf "apache" && docker-compose up -d --no-deps --build "apache"
	cd $(ROOT_DIR) && docker cp resources/httpd.conf apache-php-fpm-mysql_apache_1:/usr/local/apache2/conf/httpd.conf
	cd $(ROOT_DIR) && docker cp resources/php7.2-fpm.conf apache-php-fpm-mysql_apache_1:/usr/local/apache2/conf/extra/php7.2-fpm.conf
	cd $(ROOT_DIR) && docker cp resources/httpd-vhosts.conf apache-php-fpm-mysql_apache_1:/usr/local/apache2/conf/extra/httpd-vhosts.conf
	cd $(ROOT_DIR) && docker-compose restart "apache"
