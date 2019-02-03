ARG GITEA_TAG
FROM ${GITEA_TAG}

ARG GITEA_SSH_PORT

RUN sed -i -e "s/^Port 22$/Port ${GITEA_SSH_PORT}/" /etc/ssh/sshd_config

