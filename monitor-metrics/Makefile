ROOT_DIR = $(CURDIR)

install-node-exporter:
	cd $(ROOT_DIR) && curl -L https://github.com/prometheus/node_exporter/releases/download/v0.18.0/node_exporter-0.18.0.linux-amd64.tar.gz -o node_exporter.tar.gz \
	&& tar zxvf node_exporter.tar.gz -C . \
	&& rm -f node_exporter.tar.gz

run-node-exporter:
	cd $(ROOT_DIR) && ./node_exporter-0.18.0.linux-amd64/node_exporter \
	--web.listen-address=127.0.0.1:9100 \
	--collector.filesystem.ignored-fs-types="^(autofs|cifs|nfs4)$$"
