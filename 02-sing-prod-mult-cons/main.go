package main

import (
	"fmt"
	"sync"
	"time"
)

var messages = []string{
	"The world itself's",
	"just one big hoax.",
	"Spamming each other with our",
	"running commentary of bullshit,",
	"masquerading as insight, our social media",
	"faking as intimacy.",
	"Or is it that we voted for this?",
	"Not with our rigged elections,",
	"but with our things, our property, our money.",
	"I'm not saying anything new.",
	"We all know why we do this,",
	"not because Hunger Games",
	"books make us happy,",
	"but because we wanna be sedated.",
	"Because it's painful not to pretend,",
	"because we're cowards.",
}

const numConsumers int = 3

func produce(msgs chan<- string) {
	for _, msg := range messages {
		fmt.Printf("producer: message: %s\n", msg)
		msgs <- msg
	}
	close(msgs)
}

func consume(msgs <-chan string, wg *sync.WaitGroup, cId int) {
	for msg := range msgs {
		time.Sleep(time.Millisecond * 200)
		fmt.Printf("consumer #%d: message: %s\n", cId, msg)
	}
	wg.Done()
}

func main() {
	msgs := make(chan string)
	wg := sync.WaitGroup{}
	go produce(msgs)
	for cId := 1; cId <= numConsumers; cId++ {
		wg.Add(1)
		go consume(msgs, &wg, cId)
	}
	wg.Wait()
}
