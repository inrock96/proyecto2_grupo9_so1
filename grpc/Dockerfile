FROM ubuntu:18.04
MAINTAINER 201403504
RUN apt-get update
RUN add-apt-repository -y ppa:longsleep/golang-backports
RUN apt update
RUN apt -y install golang
RUN apt -y install golang-goprotobuf-dev
RUN rm /var/www/html/*
COPY index.html /var/www/html/
CMD ["/usr/sbin/apache2ctl", "-D", "FOREGROUND"]
EXPOSE 80