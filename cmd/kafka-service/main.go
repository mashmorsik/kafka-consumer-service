package main

import "github.com/mashmorsik/kafka-consumer-service/infrastructure/kfk"

func main() {
	err := kfk.Produce()
	if err != nil {
		panic(err)
	}

	err = kfk.Consume()
	if err != nil {
		panic(err)
	}
}
