SHELL = /bin/bash -o pipefail

include .env

rebuild-qpid-cpp-broker:
	cd $(CURDIR) && docker-compose rm -sf "qpid-cpp-broker" && docker-compose up -d --no-deps --build "qpid-cpp-broker"

exec-qpid-cpp-broker:
	cd $(CURDIR) && docker-compose exec -e TERM=$$TERM -e LINES=$$LINES -e COLUMNS=$$COLUMNS "qpid-cpp-broker" bash

restart-qpid-cpp-broker:
	cd $(CURDIR) && docker-compose restart "qpid-cpp-broker"
