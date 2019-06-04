ARG FILEBEAT_TAG
FROM ${FILEBEAT_TAG}

USER root

# Copy our custom configuration file
COPY image/filebeat/resource/filebeat.yml /usr/share/filebeat/filebeat.yml

# Create a directory to map volume with all docker log files
#RUN mkdir -p /usr/share/filebeat/dockerlogs

RUN chown -R filebeat:filebeat /usr/share/filebeat/
#RUN chmod -R go-w /usr/share/filebeat/

USER filebeat

