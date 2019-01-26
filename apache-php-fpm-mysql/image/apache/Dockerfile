FROM httpd:2.4.37

COPY image/apache/resources/httpd.conf /usr/local/apache2/conf/
COPY image/apache/resources/php7.3-fpm.conf /usr/local/apache2/conf/extra
COPY image/apache/resources/httpd-vhosts.conf /usr/local/apache2/conf/extra

RUN apt-get update \
	&& apt-get -y install procps tree less \
    && chown root:www-data /usr/local/apache2/conf/httpd.conf \
    && chown root:www-data /usr/local/apache2/conf/extra/php7.3-fpm.conf \
    && chown root:www-data /usr/local/apache2/conf/extra/httpd-vhosts.conf
