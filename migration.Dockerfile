FROM alpine:latest

RUN apk update && \
  apk upgrade && \
  apk add bash && \
  rm -rf /var/chace/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.24.3/goose_linux_x86_64 bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/*.sql migrations/
ADD migrations.sh .
ADD .env .

RUN chmod +x migrations.sh

ENTRYPOINT ["bash","migrations.sh"]