SHELL = /bin/bash -o pipefail

include .env

rebuild-nats:
	cd $(CURDIR) && docker-compose rm -sf "nats" && docker-compose up -d --no-deps --build "nats"

exec-nats:
	cd $(CURDIR) && docker-compose exec -e TERM=$$TERM -e LINES=$$LINES -e COLUMNS=$$COLUMNS "nats" sh

restart-nats:
	cd $(CURDIR) && docker-compose restart "nats"

