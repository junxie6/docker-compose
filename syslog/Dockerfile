# Use latest jboss/base-jdk:8 image as the base
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY main /usr/local/bin

CMD /usr/local/bin/main

