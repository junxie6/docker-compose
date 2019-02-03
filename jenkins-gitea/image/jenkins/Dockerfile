ARG GO_VERSION
ARG JENKINS_TAG
ARG user=jenkins

FROM jenkins/jenkins:2.161

ENV PATH ${PATH}:/usr/local/go/bin

USER root

RUN rm -rf /usr/local/go \
    && curl -L "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz" -o go.tar.gz \
    && tar -zxvf go.tar.gz -C /usr/local \
    && rm -f go.tar.gz

USER ${user}
