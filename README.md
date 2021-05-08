# Kafka Broker Environment

## Set up kafka broker

`docker compose up -d`

## Producer

`go run main.go -p`

## Consumer

`go run main.go -c`

**Note:** _To start multiples consumers, run this command in different terminals concurrently_.

This repository has the code to run the example developed based on that:
[article](https://medium.com/@ronnansouza/setting-up-a-kafka-broker-using-docker-creating-a-producer-and-consumer-group-with-multiple-384b724cd324?sk=4f828cdc1adeec088e9e67f35dbb0c8f)
