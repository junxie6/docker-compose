SHELL = /bin/bash -o pipefail

rebuild-all: rebuild-mysql rebuild-php rebuild-apache

rebuild-mysql:
	docker-compose rm -sf "mysql" && docker-compose up -d --no-deps --build "mysql"

rebuild-php:
	docker-compose rm -sf "php" && docker-compose up -d --no-deps --build "php"
	docker cp resources/www.conf apache-php-fpm-mysql_php_1:/usr/local/etc/php-fpm.d/www.conf

rebuild-apache:
	docker-compose rm -sf "apache" && docker-compose up -d --no-deps --build "apache"
	docker cp resources/httpd.conf apache-php-fpm-mysql_apache_1:/usr/local/apache2/conf/httpd.conf
	docker cp resources/php7.2-fpm.conf apache-php-fpm-mysql_apache_1:/usr/local/apache2/conf/extra/php7.2-fpm.conf
	docker cp resources/httpd-vhosts.conf apache-php-fpm-mysql_apache_1:/usr/local/apache2/conf/extra/httpd-vhosts.conf
