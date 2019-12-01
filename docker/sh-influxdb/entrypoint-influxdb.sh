#!/bin/sh

main() {
  create_sh_db & \
  influxd
}

create_sh_db() {
  while ! curl -sl -I http://localhost:8086/ping 
  do
    sleep 1
  done

  echo "Influxdb is running"

  curl "http://localhost:8086/query" --data-urlencode "q=CREATE DATABASE smarthome"
}

main
