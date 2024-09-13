#!/bin/bash

sleep 2
/usr/bin/kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 3 --topic sms
/usr/bin/kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 3 --topic email
/usr/bin/kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 3 --topic slack

echo "Kafka topics created successfully."
