FROM debian:buster

# Install generictools
RUN apt-get update && apt-get upgrade -y && apt-get install -y \
wget \
curl

# Set env variables
ENV GOPATH=/root/go \
PATH="${PATH}:${GOPATH}/bin" \
SH_HTTP_PORT=3222 \
SH_CLI_TCP_PORT=3333 \
SH_MAILER_EMAIL_NAME="" \
SH_MAILER_EMAIL_PASS="" \
SH_MAILER_SMTP_PORT=smtp.gmail.com:587 \
SH_MAILER_SMTP_AUTHURL=smtp.gmail.com \
SH_MONGO_URI=mongodb://localhost:27017 \
SH_MONGO_DB=smarthome \
SH_INFLUX_URI=http://localhost:8086

# Install smarthome project
RUN mkdir -p /root/go/src/github.com/smart-evolution/shapi/public

# Add mocks
RUN mkdir /root/mocks && \
mkdir /root/scripts
COPY ./shapi /root/go/src/github.com/smart-evolution/shapi/shapi
COPY ./docker/sh-api/prod/mock_data.sh /root/scripts/
COPY ./docker/sh-api/prod/start_shapi.sh /root/scripts/
COPY ./docker/sh-api/prod/entrypoint-prod.sh /root/scripts/

ENTRYPOINT /root/scripts/entrypoint-prod.sh
