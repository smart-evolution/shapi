FROM debian:buster

# Install generictools
RUN apt-get update && apt-get upgrade -y && apt-get install -y \
wget \
python \
vim \
git \
build-essential

# Install nodejs
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash - && \
apt-get install -y nodejs

RUN wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz && \
tar -C /usr/local -xvf go1.12.7.linux-amd64.tar.gz
ENV PATH="${PATH}:/usr/local/go/bin"

# Set env variables
ENV GOPATH=/root/go \
GO111MODULE=on \
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
