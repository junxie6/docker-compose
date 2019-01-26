FROM php:7.3.1-fpm-stretch

ARG MODE

COPY image/php/resources/www.conf /usr/local/etc/php-fpm.d/
COPY image/php/resources/php.ini-$MODE /usr/local/etc/php/php.ini
COPY image/php/resources/demo.php /var/www/html/

RUN apt-get update \
	&& apt-get -y install procps tree less \
    && chown root:staff /usr/local/etc/php-fpm.d/www.conf \
    && chown root:staff /usr/local/etc/php/php.ini \
    && chown root:www-data /var/www/html/demo.php
