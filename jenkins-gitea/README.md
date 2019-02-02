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

### Start the services

```
$ git clone https://github.com/junxie6/docker-compose.git
$ cd docker-compose/jenkins-gitea
$ cp .env.dist .env
$ make rebuild-all
```

User a browser to access http://127.0.0.1:3000

### Install Generic Webhook Trigger Plugin

1. Manage Jenkins &gt; Manage Plugin &gt; Generic Webhook Trigger Plugin

2. Check "Generic Webhook Trigger" checkbox

3. Token: TOKEN_HERE

### Set up webhook on Gitea's repository setting

1. Repoistory &gt; Settings &gt; Webhooks

2. Target URL

```
http://gitea.local:8080/generic-webhook-trigger/invoke?token=TOKEN_HERE
```

