FROM gitlab/gitlab-ce:latest

ARG GITLAB_SSH_PORT

RUN sed -i -e "s/^Port 22$/Port ${GITLAB_SSH_PORT}/" /etc/ssh/sshd_config
