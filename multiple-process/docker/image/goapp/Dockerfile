# Reference:
# https://github.com/just-containers/base-alpine
# https://github.com/just-containers/s6-overlay/blob/master/builder/overlay-rootfs/usr/bin/with-contenv
FROM alpine:3.12.0

ARG S6_OVERLAY_RELEASE=https://github.com/just-containers/s6-overlay/releases/latest/download/s6-overlay-amd64.tar.gz
ENV S6_OVERLAY_RELEASE=${S6_OVERLAY_RELEASE}

ADD image/goapp/rootfs /

# Add logrotate command and its config
ADD image/goapp/logrotate.sample.conf /goapp/logrotate.conf
ADD image/goapp/logrotate /etc/periodic/15min

# Add services
ADD image/goapp/services.d/cron /etc/services.d/cron
ADD image/goapp/services.d/goapp /etc/services.d/goapp

# Add srv_out
ADD image/goapp/srv_out /goapp

# s6 overlay Download
ADD ${S6_OVERLAY_RELEASE} /tmp/s6overlay.tar.gz

# Build and some of image configuration
RUN apk upgrade --update --no-cache \
    && apk add --no-cache logrotate \
    && apk add --no-cache iproute2 tcptraceroute util-linux bind-tools \
    && rm -rf /var/cache/apk/* \
    && tar xzf /tmp/s6overlay.tar.gz -C / \
    && rm /tmp/s6overlay.tar.gz

# Init
ENTRYPOINT [ "/init" ]
