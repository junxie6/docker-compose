# Set up MySQL Master Slave Replication

##### Create and start master and slave containers:

\# docker-compose up -d

##### Copy the custom-mysqld-master.cnf and the .my.cnf to the master server and change its permission:

\# docker cp custom-mysqld-master.cnf mysqlreplication_mysql-master_1:/etc/mysql/mysql.conf.d<br>
\# docker exec -it mysqlreplication_mysql-master_1 chown root:root /etc/mysql/mysql.conf.d/custom-mysqld-master.cnf<br>
\# docker exec -it mysqlreplication_mysql-master_1 chmod 644 /etc/mysql/mysql.conf.d/custom-mysqld-master.cnf<br>

\# docker cp .my.cnf mysqlreplication_mysql-master_1:/root<br>
\# docker exec -it mysqlreplication_mysql-master_1 chmod 400 /root/.my.cnf<br>

##### Copy the custom-mysqld-slave.cnf and the .my.cnf to the slave server and change its permission:

\# docker cp custom-mysqld-slave.cnf mysqlreplication_mysql-slave_1:/etc/mysql/mysql.conf.d<br>
\# docker exec -it mysqlreplication_mysql-slave_1 chown root:root /etc/mysql/mysql.conf.d/custom-mysqld-slave.cnf<br>
\# docker exec -it mysqlreplication_mysql-slave_1 chmod 644 /etc/mysql/mysql.conf.d/custom-mysqld-slave.cnf<br>

\# docker cp .my.cnf mysqlreplication_mysql-slave_1:/root<br>
\# docker exec -it mysqlreplication_mysql-slave_1 chmod 400 /root/.my.cnf<br>

##### Restart master and slave containers:

\# docker-compose restart mysql-master<br>
\# docker-compose restart mysql-slave<br>

##### Login into the master container:

\# docker exec -it mysqlreplication_mysql-master_1 bash

##### Login into the slave container:

\# docker exec -it mysqlreplication_mysql-slave_1 bash

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

##### On slave, show status information on essential parameters of the slave threads:

root@mysql-slave # mysql -e "SHOW SLAVE STATUS \\G;" | grep -E 'Slave_|Master_'


