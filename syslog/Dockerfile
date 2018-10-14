#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
FROM ubuntu:16.04

#
ENV TZ Asia/Shanghai

RUN echo $TZ > /etc/timezone \
  && apt-get -y update && apt-get install -y tzdata \
  && rm /etc/localtime \
  && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
  && dpkg-reconfigure -f noninteractive tzdata \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

#
RUN apt-get -y update \
  && apt-get -y install vim tree git curl htop less wget procps iproute dnsutils netcat-openbsd ca-certificates \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

COPY main /usr/local/bin

RUN chmod 755 /usr/local/bin/main \
  && groupadd -r -g 600 webadmin \
  && groupadd -r -g 601 jboss \
  && useradd -r -u 601 -g 601 -s /sbin/nologin -c 'web user' jboss \
  && usermod -aG webadmin jboss \
  && mkdir /ASDF \
  && touch /ASDF/a.txt \
  && echo hihi /ASDF/b.txt

USER jboss

CMD /usr/local/bin/main

