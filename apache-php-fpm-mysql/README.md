# Apache, PHP-FPM, and MySQL containers docker-compose.yml

## Getting started

### Prerequisites

#### Docker setting

```
$ vim /etc/docker/daemon.json

{
  "bridge": "none",
  "dns": ["8.8.8.8"],
  "ipv6": false
}
```

#### Hosts setting

```
$ sudo vim /etc/hosts

127.0.0.1 dummy-host2.example.com
```

### Start

```
$ git clone https://github.com/junxie6/docker-compose.git
$ cd docker-compose/apache-php-fpm-mysql
$ make rebuild-all
$ curl http://dummy-host2.example.com/demo.php
```
