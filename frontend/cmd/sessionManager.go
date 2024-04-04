package main

import (
	"fmt"
	"sync"
)

type SessionChannel chan string

var channels = map[string]SessionChannel{}
var mutex = sync.Mutex{}

func (SessionChannel) createChannel(pollId string) {
	mutex.Lock()
	defer mutex.Unlock()

	channel := make(SessionChannel)
	channels[pollId] = channel
	//return channel
}

func (SessionChannel) sendMessage(pollId, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	channel, ok := channels[pollId]
	if ok {
		channel <- message
	} else {
		fmt.Println("Session not found for channel")
	}
}

func (SessionChannel) waitOnChannel(pollId string) {

	channel, ok := channels[pollId]
	if !ok {
		fmt.Println("Cannot wait on channel since not found")
	}

loop:
	for {
		select {
		case msg := <-channel:
			fmt.Printf("Received msg: %s\n", msg)
			break loop
		}
	}
}

func (SessionChannel) closeChannel(pollId string) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(channels, pollId)
}
