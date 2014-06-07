FROM ubuntu:precise
ENV DEBIAN_FRONTEND noninteractive


VOLUME ["/conf"]


WORKDIR /app

EXPOSE 9000

ADD stinger /app/stinger
ADD stinger.sh /app/stinger.sh
ADD config_example.json /conf/config.json
CMD ["./stinger.sh"]
