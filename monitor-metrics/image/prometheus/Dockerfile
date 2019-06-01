ARG PROMETHEUS_TAG
FROM ${PROMETHEUS_TAG}

CMD apt-get update \
  && apt-get -y install iproute

COPY image/prometheus/resource/prometheus.yml /etc/prometheus/prometheus.yml

