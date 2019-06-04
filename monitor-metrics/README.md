# Prometheus

```
# vim /etc/sysctl.conf
```
```
vm.max_map_count=262144
```
```
# sysctl -p
```

NOTE: https://stackoverflow.com/questions/42300463/elasticsearch-bootstrap-checks-failing

NOTE: https://stackoverflow.com/questions/55956645/docker-compose-yml-for-elasticsearch-7-0-1-and-kibana-7-0-1

```
dexec elasticsearch bash
bin/elasticsearch-setup-passwords interactive
```

```
dexec elasticsearch bash

vi config/elasticsearch.yml

xpack.security.enabled: true

exit

drestart elasticsearch
```

```
dexec kibana bash

vi config/kibana.yml

elasticsearch.username: "kibana"
elasticsearch.password: "kibanapassword"

exit

drestart kibana
```

```
dexec logstash bash

vi config/logstash.yml
```

```
$ nc -zv -u 127.0.0.1 9000
```

```
$ logger --server 127.0.0.1 --port 9000 --udp --rfc3164 "my testing msg"
$ curl -u elastic:elastic -X GET "localhost:9200/filebeat-*/_search"
```
