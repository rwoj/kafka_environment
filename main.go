package main

import (
	"flag"
	"fmt"

	"github.com/rwoj/kafka_environment/consumer"
	"github.com/rwoj/kafka_environment/producer"
)

func main() {

	stringFlagValue := flag.String("o", "", "    Use this flag with either \"consumer\" or \"producer\"")

	// Flag Processing
	flag.Parse()

	// Decision Time
	if *stringFlagValue == "consumer" {
		consumer.StartConsumer()
	} else if *stringFlagValue == "producer" {
		producer.StartProducer()
	} else {
		fmt.Print("Usage: \n -o     and either \"consumer\" or \"producer\"\n")
	}
}
