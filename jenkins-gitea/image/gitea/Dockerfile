ARG GITEA_TAG
ARG GITEA_SSH_PORT

FROM ${GITEA_TAG}

RUN sed -i -e "s/^Port 22$/Port ${GITEA_SSH_PORT}/" /etc/ssh/sshd_config

#EXPOSE ${GITEA_SSH_PORT} 3000

