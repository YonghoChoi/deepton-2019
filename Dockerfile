FROM debian:jessie

RUN apt-get update && apt-get install -y vim

COPY bin/pengha-api /opt/
RUN chmod +x /opt/pengha-api

CMD ["/opt/pengha-api"]