package main

import (
	"fmt"
	"sync"
)

type ChannelManager struct {
	ChannelSlice map[string]chan string
}

// var channels = map[string]ChannelManager{}
var mutex = sync.Mutex{}

func NewChannelManager() *ChannelManager {
	mutex.Lock()
	defer mutex.Unlock()

	c := &ChannelManager{
		ChannelSlice: make(map[string]chan string),
	}

	return c
}

func (c *ChannelManager) createChannel(pollId string) {
	mutex.Lock()
	defer mutex.Unlock()

	channel := make(chan string)
	c.ChannelSlice[pollId] = channel
}

func (c *ChannelManager) sendMessage(pollId, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	channel, ok := c.ChannelSlice[pollId]
	if ok {
		channel <- message
	} else {
		fmt.Println("Session not found for channel")
	}
}

func (c *ChannelManager) waitOnChannel(pollId string) {

	channel, ok := c.ChannelSlice[pollId]
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

func (c *ChannelManager) closeChannel(pollId string) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(c.ChannelSlice, pollId)
}
