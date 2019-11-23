FROM debian:jessie

RUN apt-get update && apt-get install -y vim

ADD bin /opt
RUN chmod +x /opt/pengha-api
RUN ls -l /opt
WORKDIR /opt

CMD ["./pengha-api", "-config", "./config.yml"]