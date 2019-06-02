SHELL = /bin/bash -o pipefail

ROOT_DIR = $(CURDIR)
SRC_DIR = $(ROOT_DIR)/src

go-run:
	cd $(SRC_DIR) && go mod vendor -v
	cd $(SRC_DIR) && go run -v -mod vendor -race main.go

