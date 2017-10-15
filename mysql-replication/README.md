Create and start master and slave containers:

\# docker-compose up -d

Copy the custom-mysqld-master.cnf and the .my.cnf to the master server and change its permission:

\# docker cp custom-mysqld-master.cnf mysqlreplication_mysql-master_1:/etc/mysql/mysql.conf.d<br>
\# docker cp .my.cnf mysqlreplication_mysql-master_1:/root<br>
\# docker exec -it mysqlreplication_mysql-master_1 chmod 400 /root/.my.cnf<br>

Copy the custom-mysqld-slave.cnf and the .my.cnf to the slave server and change its permission:

\# docker cp custom-mysqld-slave.cnf mysqlreplication_mysql-slave_1:/etc/mysql/mysql.conf.d<br>
\# docker cp .my.cnf mysqlreplication_mysql-slave_1:/root<br>
\# docker exec -it mysqlreplication_mysql-slave_1 chmod 400 /root/.my.cnf<br>

Login into the master container:

\# docker exec -it mysqlreplication_mysql-master_1 bash

Login into the slave container:

\# docker exec -it mysqlreplication_mysql-slave_1 bash

Create a user account for replication on master:

root@mysql-master # mysql

mysql> GRANT REPLICATION SLAVE ON \*.\* TO 'slave_user'@'192.168.5.8/255.255.255.248' IDENTIFIED BY 'slave_pass';<br>
mysql> FLUSH PRIVILEGES;<br>
mysql> exit<br>

Show the grant user command on master:

root@mysql-master # mysql -e "SELECT CONCAT('mysql -e \\"SHOW GRANTS FOR ', '\\'', user, '\\'', '@', '\\'', host, '\\'',';\\"') AS userHost FROM mysql.user;"

Show the create user command on master:

root@mysql-master # mysql -e "SELECT CONCAT('mysql -e \\"SHOW CREATE USER ', '\\'', user, '\\'', '@', '\\'', host, '\\'',';\\"') AS userHost FROM mysql.user;"

Deletes all binary log files listed in the index file, resets the binary log index file to be empty, and creates a new binary log file:

root@mysql-master # mysql -e "RESET MASTER;"

Note: Use this statement with caution to ensure you do not lose binary log file data.

Show status information about the binary log files of the master:

root@mysql-master # mysql -e "SHOW MASTER STATUS \G;"
```
             File: mysql-bin.000001
         Position: 154
```

Stop the slave threads:

\# mysql -e "STOP SLAVE;"

Clear the master info and relay log info repositories, deletes all the relay log files, and starts a new relay log file:

\# mysql -e "RESET SLAVE;"

Note: This statement is meant to be used for a clean start.

mysql> CHANGE MASTER TO<br>
MASTER_HOST='mysql-master',<br>
MASTER_USER='slave_user',<br>
MASTER_PASSWORD='slave_pass',<br>
MASTER_CONNECT_RETRY=60,<br>
MASTER_LOG_FILE='mysql-bin.000001',<br>
MASTER_LOG_POS=154<br>
;<br>

\# mysql -e "START SLAVE;"

\# mysql -e "SHOW SLAVE STATUS \G;"
