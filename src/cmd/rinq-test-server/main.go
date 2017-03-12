package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/rinq/rinq-go/src/rinq/amqp"
	"github.com/rinq/test/src/commands"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// TODO: this env var will be handled by rinq-go
	// https://github.com/rinq/rinq-go/issues/94
	peer, err := amqp.Dial(os.Getenv("RINQ_AMQP_DSN"))
	if err != nil {
		panic(err)
	}

	if err = peer.Listen("rinq.test", commands.NewHandler(peer)); err != nil {
		panic(err)
	}

	<-peer.Done()

	if err = peer.Err(); err != nil {
		panic(err)
	}
}
