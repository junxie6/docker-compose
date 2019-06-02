SHELL = /bin/bash -o pipefail

ROOT_DIR = $(CURDIR)
SRC_DIR = $(ROOT_DIR)/src

install-node-exporter:
	cd $(ROOT_DIR) && curl -L https://github.com/prometheus/node_exporter/releases/download/v0.18.0/node_exporter-0.18.0.linux-amd64.tar.gz -o node_exporter.tar.gz \
	&& tar zxvf node_exporter.tar.gz -C . \
	&& rm -f node_exporter.tar.gz

run-node-exporter:
	cd $(ROOT_DIR) && ./node_exporter-0.18.0.linux-amd64/node_exporter \
	--web.listen-address=127.0.0.1:9100 \
	--collector.filesystem.ignored-fs-types="^(autofs|cifs|nfs4)$$"

demo-sql-create-table:
	bash -c "for i in {1..1}; do docker exec monitor-metrics_mssql_1 /opt/mssql-tools/bin/sqlcmd -S 127.0.0.1 -U SA -P Change_Admin_Pass6 -Q 'USE [exp]; CREATE TABLE [dbo].[Bug]([BugID] [bigint] IDENTITY(1,1) NOT NULL, [Subject] [nchar](32) NOT NULL, [CreatedDT] [datetime] NOT NULL) ON [PRIMARY]' && sleep 1 ; done"

demo-sql-select-version:
	bash -c "for i in {1..1}; do docker exec monitor-metrics_mssql_1 /opt/mssql-tools/bin/sqlcmd -S 127.0.0.1 -U SA -P Change_Admin_Pass6 -Q 'SELECT @@VERSION, GETDATE()' && sleep 1 ; done"

demo-sql-select-count:
	bash -c "for i in {1..1}; do docker exec monitor-metrics_mssql_1 /opt/mssql-tools/bin/sqlcmd -S 127.0.0.1 -U SA -P Change_Admin_Pass6 -Q 'SELECT @@VERSION, GETDATE()' && sleep 1 ; done"

demo-sql-insert:
	for i in {1..10}; do docker exec monitor-metrics_mssql_1 /opt/mssql-tools/bin/sqlcmd -S 127.0.0.1 -U SA -P Change_Admin_Pass6 -Q "INSERT INTO [exp].[dbo].[Bug] ([Subject], [CreatedDT]) VALUES ('test', GETDATE());"  && sleep 1 ; done \

go-run:
	cd $(SRC_DIR) && go mod vendor -v
	cd $(SRC_DIR) && go run -v -mod vendor -race main.go
