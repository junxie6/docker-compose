FROM docker.elastic.co/beats/filebeat-oss:6.4.2

USER root

# Copy our custom configuration file
COPY filebeat.yml /usr/share/filebeat/filebeat.yml

# Create a directory to map volume with all docker log files
#RUN mkdir -p /usr/share/filebeat/dockerlogs

RUN chown -R filebeat:filebeat /usr/share/filebeat/
#RUN chmod -R go-w /usr/share/filebeat/

USER filebeat

