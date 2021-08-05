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

const (
	numProducers int = 3
	numConsumers int = 3
)

func produce(msgs chan<- string, wg *sync.WaitGroup, mu *sync.Mutex, msgIdx *int, pId int) {
	msgsLen := len(messages)
	for {
		mu.Lock()
		idx := *msgIdx
		if idx >= msgsLen {
			mu.Unlock()
			break
		}
		*msgIdx = idx + 1
		mu.Unlock()
		msg := messages[idx]
		fmt.Printf("producer #%d message: %s\n", pId, msg)
		msgs <- msg
	}
	wg.Done()
}

func consume(msgs <-chan string, wg *sync.WaitGroup, cId int) {
	for msg := range msgs {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf(" > consumer #%d message: %s\n", cId, msg)
	}
	wg.Done()
}

func main() {
	msgs := make(chan string)
	prodWg := sync.WaitGroup{}
	consWg := sync.WaitGroup{}
	mu := sync.Mutex{}
	msgIdx := 0
	for pId := 1; pId <= numProducers; pId++ {
		prodWg.Add(1)
		go produce(msgs, &prodWg, &mu, &msgIdx, pId)
	}
	for cId := 1; cId <= numConsumers; cId++ {
		consWg.Add(1)
		go consume(msgs, &consWg, cId)
	}
	prodWg.Wait()
	close(msgs)
	consWg.Wait()
}
