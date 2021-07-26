ARG ALPINE_tag
FROM ${ALPINE_tag}

ARG APP_NAME
ARG USER_ID
ARG USERNAME

#
# Important! Update this no-op ENV variable when this Dockerfile
# is updated with the current date. It will force refresh of all
# of the base images.
ENV REFRESHED_AT=2021-07-25 \
    PATH=/usr/local/go/bin:$PATH

# Install some packages
RUN apk upgrade --update --no-cache \
    && apk add --no-cache make \
    # install ca-certificates so that HTTPS works consistently
    && apk add --no-cache ca-certificates \
    # and tzdata for PEP 615 (https://www.python.org/dev/peps/pep-0615/)
    && apk add --no-cache tzdata \
    && apk add --no-cache bash bash-completion curl util-linux \
    && apk add --no-cache logrotate \
    && apk add --no-cache iproute2 tcptraceroute util-linux bind-tools \
    && rm -rf /var/cache/apk/* \
    && adduser -u ${USER_ID} -s /bin/bash -D ${USERNAME}

COPY --chown=${USERNAME}:${USERNAME} image/${APP_NAME}/appdev/ /home/${USERNAME}

# Install Node, Python, and Go
COPY --from=node:12.22.3-alpine3.14 /usr/local /usr/local
COPY --from=node:12.22.3-alpine3.14 /usr/lib /usr/lib
COPY --from=python:3.9.6-alpine3.14 /usr/local /usr/local
COPY --from=python:3.9.6-alpine3.14 /usr/lib /usr/lib
COPY --from=golang:1.16.6-alpine3.14 /usr/local/go /usr/local/go

# Install Angular
RUN npm install -g --silent @angular/cli@12.1.3

# Switch to a regular user
USER ${USERNAME}

# Install Python packages
RUN python -m pip install --user --upgrade pip wheel setuptools

# Define working directory
WORKDIR /${APP_NAME}
