ARG JENKINS_TAG
FROM ${JENKINS_TAG}

ARG JENKINS_USER
ARG GO_VERSION

ENV PATH ${PATH}:/usr/local/go/bin

USER root

RUN rm -rf /usr/local/go \
    && curl -L "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz" -o go.tar.gz \
    && tar -zxvf go.tar.gz -C /usr/local \
    && rm -f go.tar.gz

USER ${JENKINS_USER}
