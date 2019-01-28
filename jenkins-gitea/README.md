# Jenkins, Gitea containers in docker-compose.yml example

## Getting started

### Prerequisites

#### Docker setting

```
$ vim /etc/docker/daemon.json

{
  "bridge": "none",
  "dns": ["8.8.8.8", "8.8.4.4"],
  "ipv6": false
}
```

```
$ sudo systemctl restart docker.service
```

### Start

```
$ git clone https://github.com/junxie6/docker-compose.git
$ cd docker-compose/jenkins-gitea
$ cp .env.dist .env
$ make rebuild-all
```

User a browser to access http://127.0.0.1:3000
