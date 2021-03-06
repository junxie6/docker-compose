# Set up MySQL Master Slave Replication on Docker Containers

##### Clone this repository and switch to the mysql-replication directory:

\# git clone https://github.com/junhsieh/docker-compose.git<br>
\# cd docker-compose/mysql-replication<br>

##### Create and start master and slave containers:

\# docker-compose up -d

##### Copy the custom-mysqld-master.cnf and the .my.cnf to the master server and change its permission:

\# (<br>
docker cp custom-mysqld-master.cnf mysql-replication_mysql-master_1:/etc/mysql/mysql.conf.d<br>
docker exec -it mysql-replication_mysql-master_1 chown root:root /etc/mysql/mysql.conf.d/custom-mysqld-master.cnf<br>
docker exec -it mysql-replication_mysql-master_1 chmod 644 /etc/mysql/mysql.conf.d/custom-mysqld-master.cnf<br>

docker cp .my.cnf mysql-replication_mysql-master_1:/root<br>
docker exec -it mysql-replication_mysql-master_1 chmod 400 /root/.my.cnf<br>
)<br>

##### Copy the custom-mysqld-slave.cnf and the .my.cnf to the slave server and change its permission:

\# (<br>
docker cp custom-mysqld-slave.cnf mysql-replication_mysql-slave_1:/etc/mysql/mysql.conf.d<br>
docker exec -it mysql-replication_mysql-slave_1 chown root:root /etc/mysql/mysql.conf.d/custom-mysqld-slave.cnf<br>
docker exec -it mysql-replication_mysql-slave_1 chmod 644 /etc/mysql/mysql.conf.d/custom-mysqld-slave.cnf<br>

docker cp .my.cnf mysql-replication_mysql-slave_1:/root<br>
docker exec -it mysql-replication_mysql-slave_1 chmod 400 /root/.my.cnf<br>
)<br>

##### Restart master and slave containers:

\# docker-compose restart mysql-master<br>
\# docker-compose restart mysql-slave<br>

##### Login into the master container:

\# docker exec -it mysql-replication_mysql-master_1 bash

##### Login into the slave container:

\# docker exec -it mysql-replication_mysql-slave_1 bash

##### On master, create a user account for replication:

root@mysql-master # mysql

mysql> GRANT REPLICATION SLAVE ON \*.\* TO 'slave_user'@'192.168.5.8/255.255.255.248' IDENTIFIED BY 'slave_pass';<br>
mysql> FLUSH PRIVILEGES;<br>
mysql> exit<br>

##### On master, show the grant user command:

root@mysql-master # mysql -e "SELECT CONCAT('mysql -e \\"SHOW GRANTS FOR ', '\\'', user, '\\'', '@', '\\'', host, '\\'',';\\"') AS userHost FROM mysql.user;"

##### On master, show the create user command:

root@mysql-master # mysql -e "SELECT CONCAT('mysql -e \\"SHOW CREATE USER ', '\\'', user, '\\'', '@', '\\'', host, '\\'',';\\"') AS userHost FROM mysql.user;"

##### Deletes all binary log files listed in the index file, resets the binary log index file to be empty, and creates a new binary log file:

root@mysql-master # mysql -e "RESET MASTER;"

**Note:** Use this statement with caution to ensure you do not lose binary log file data.

##### On master, show status information about the binary log files of the master:

root@mysql-master # mysql -e "SHOW MASTER STATUS \\G;"
```
             File: mysql-bin.000001
         Position: 154
```

##### On slave, stop the slave threads:

root@mysql-slave # mysql -e "STOP SLAVE;"

##### On slave, clear the master info and relay log info repositories, deletes all the relay log files, and starts a new relay log file:

root@mysql-slave # mysql -e "RESET SLAVE;"

**Note:** This statement is meant to be used for a clean start.

##### On slave, change the parameters that the slave server uses for connecting to the master server:

root@mysql-slave # mysql

mysql> CHANGE MASTER TO<br>
MASTER_HOST='mysql-master',<br>
MASTER_USER='slave_user',<br>
MASTER_PASSWORD='slave_pass',<br>
MASTER_CONNECT_RETRY=60,<br>
MASTER_LOG_FILE='mysql-bin.000001',<br>
MASTER_LOG_POS=154<br>
;<br>
mysql> exit<br>

##### On slave, start the slave threads:

root@mysql-slave # mysql -e "START SLAVE;"

**Note:** Because we have set **skip_slave_start** in custom-mysqld-slave.cnf. So, we need to start slave manually.

##### On slave, show status information on essential parameters of the slave threads:

root@mysql-slave # mysql -e "SHOW SLAVE STATUS \\G;" | grep -E 'Slave_|Master_|Seconds_' | sort -b

##### Benchmark the slave compressed protocol:

root@mysql-slave # iftop -P

root@mysql-master # bash bench.bash

##### Install server_audit plugin:

\# (<br>
docker cp custom-mysqld-audit.cnf mysql-replication_mysql-master_1:/etc/mysql/mysql.conf.d<br>
docker exec -it mysql-replication_mysql-master_1 chown root:root /etc/mysql/mysql.conf.d/custom-mysqld-audit.cnf<br>
docker exec -it mysql-replication_mysql-master_1 chmod 644 /etc/mysql/mysql.conf.d/custom-mysqld-audit.cnf<br>

docker cp custom-mysqld-audit.cnf mysql-replication_mysql-slave_1:/etc/mysql/mysql.conf.d<br>
docker exec -it mysql-replication_mysql-slave_1 chown root:root /etc/mysql/mysql.conf.d/custom-mysqld-audit.cnf<br>
docker exec -it mysql-replication_mysql-slave_1 chmod 644 /etc/mysql/mysql.conf.d/custom-mysqld-audit.cnf<br>
)<br>

\# docker exec -it mysql-replication_mysql-master_1 bash

root@mysql-master # mysql -e "SHOW GLOBAL variables WHERE variable_name REGEXP 'plugin_dir';"

root@mysql-master # curl -o server_audit.tar.gz https://downloads.mariadb.com/Audit-Plugin/MariaDB-Audit-Plugin/server_audit-1.4.0.tar.gz \\<br>
&& tar zxvf server_audit.tar.gz \\<br>
&& cp server_audit-1.4.0/linux-x86-64-glibc_214/server_audit.so /usr/lib/mysql/plugin/ \\<br>
&& chown root:root /usr/lib/mysql/plugin/server_audit.so \\<br>
&& chmod 644 /usr/lib/mysql/plugin/server_audit.so \\<br>
&& ls -l /usr/lib/mysql/plugin/server_audit.so \\<br>
&& rm -f server_audit.tar.gz \\<br>
&& rm -rf server_audit-1.4.0<br>

**Note:** Run the same command above for the slave.

\# docker-compose restart mysql-master<br>
\# docker-compose restart mysql-slave<br>

root@mysql-master # mysql -e "SHOW GLOBAL variables WHERE variable_name REGEXP 'server_audit' AND variable_name != 'server_audit_loc_info';"<br>
root@mysql-master # tail /var/log/mysql/audit.log<br>

##### Reference:

https://mariadb.com/resources/blog/introducing-mariadb-audit-plugin

https://www.percona.com/blog/2016/02/15/mysql-mariadb-with-mariadb-auditing-plugin/
