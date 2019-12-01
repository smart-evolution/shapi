FROM debian:buster

# Install generictools
RUN apt-get update && apt-get upgrade -y && apt-get install -y \
wget \
python \
vim \
git \
curl \
build-essential 

# Install influxdb
RUN curl -sL https://repos.influxdata.com/influxdb.key | apt-key add - && \
apt install apt-transport-https && \
echo "deb https://repos.influxdata.com/debian buster stable" | tee /etc/apt/sources.list.d/influxdb.list && \
apt update && \
apt install influxdb  

RUN mkdir /root/scripts
COPY ./docker/sh-influxdb/entrypoint-influxdb.sh /root/scripts

ENTRYPOINT /root/scripts/entrypoint-influxdb.sh
