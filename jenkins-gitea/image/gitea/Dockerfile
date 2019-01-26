FROM gitea/gitea:1.7.0

ARG GITEA_SSH_PORT

RUN sed -i -e "s/^Port 22$/Port ${GITEA_SSH_PORT}/" /etc/ssh/sshd_config

#EXPOSE ${GITEA_SSH_PORT} 3000

