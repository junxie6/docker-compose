Copy the custom-mysqld-master.cnf:

\# docker cp custom-mysqld-master.cnf mysqlreplication_mysql-master_1:/etc/mysql/mysql.conf.d

Create a user account for replication:

mysql> GRANT REPLICATION SLAVE ON *.* TO 'slave_user'@'192.168.5.8/255.255.255.248' IDENTIFIED BY 'slave_pass';<br>
mysql> FLUSH PRIVILEGES;<br>

Show the grant user command:

\# mysql -e "SELECT CONCAT('mysql -e \"SHOW GRANTS FOR ', '\'', user, '\'', '@', '\'', host, '\'',';\"') AS userHost FROM mysql.user;"

Show the create user command:

\# mysql -e "SELECT CONCAT('mysql -e \"SHOW CREATE USER ', '\'', user, '\'', '@', '\'', host, '\'',';\"') AS userHost FROM mysql.user;"


\# mysql -e "RESET MASTER;"

\# mysql -e "SHOW MASTER STATUS \G;"
```
             File: mysql-bin.000001
         Position: 154
```

\# mysql -e "STOP SLAVE;"
\# mysql -e "RESET SLAVE;"

mysql> CHANGE MASTER TO
MASTER_HOST='mysql-master',
MASTER_USER='slave_user',
MASTER_PASSWORD='slave_pass',
MASTER_CONNECT_RETRY=60,
MASTER_LOG_FILE='mysql-bin.000001',
MASTER_LOG_POS=154
;

\# mysql -e "START SLAVE;"

\# mysql -e "SHOW SLAVE STATUS \G;"
