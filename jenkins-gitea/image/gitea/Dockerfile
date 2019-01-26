FROM gitea/gitea:1.7.0

ARG SSH_PORT

RUN sed -i -e "s/^Port 22$/Port ${SSH_PORT}/" /etc/ssh/sshd_config

#EXPOSE ${SSH_PORT} 3000

