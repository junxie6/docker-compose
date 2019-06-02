ARG FILEBEAT_URL
ARG FILEBEAT_TAG
FROM ${FILEBEAT_URL}${FILEBEAT_TAG}

USER root

# Copy our custom configuration file
COPY filebeat.yml /usr/share/filebeat/filebeat.yml
COPY image/filebeat/resource/filebeat.yml /etc/prometheus/prometheus.yml

# Create a directory to map volume with all docker log files
#RUN mkdir -p /usr/share/filebeat/dockerlogs

RUN chown -R filebeat:filebeat /usr/share/filebeat/
#RUN chmod -R go-w /usr/share/filebeat/

USER filebeat

