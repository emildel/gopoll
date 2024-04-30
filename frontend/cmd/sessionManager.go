package main

import (
	"errors"
	"fmt"
	"sync"
)

type ChannelManager struct {
	mu           sync.Mutex
	ChannelSlice map[string]*Poll
}

type Poll struct { // Manager subscription was created from
	channel chan string
	id      string
	once    sync.Once // ensures we can only close channel c once
}

func NewChannelManager() *ChannelManager {
	return &ChannelManager{
		ChannelSlice: make(map[string]*Poll),
	}
}

func (c *ChannelManager) CreateSubscription(pollId string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	sub := &Poll{
		// Channel of size 1, as long as there is a value still in the channel,
		// there is updating to be done by the client.
		channel: make(chan string, 1),
		id:      pollId,
		once:    sync.Once{},
	}

	c.ChannelSlice[pollId] = sub
}

func (c *ChannelManager) SendMessage(pollId, message string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	channel, ok := c.ChannelSlice[pollId]
	if !ok {
		fmt.Println("Session not found for channel")
		return
	}

	select {
	// If channel is not full (i.e. empty in our case), send a message to signal we should do an update
	case channel.channel <- message:
		// If the channel is full, we can just return since we're using the channel as a
		// signal for the client to do an update to the database, and a "signal" is
		// already there
	default:
		return
	}
}

func (c *ChannelManager) WaitOnChannel(pollId string) error {

	channel, ok := c.ChannelSlice[pollId]
	if !ok {
		return ErrChannelNotFound
	}

	select {

	// Block until we have a signal, then return
	case msg := <-channel.channel:
		fmt.Printf("Received msg: %s\n", msg)
	}
	return nil
}

func (c *ChannelManager) Unsubscribe(poll *Poll) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.unsubscribe(poll)
}

func (c *ChannelManager) unsubscribe(poll *Poll) {
	poll.once.Do(func() {
		close(poll.channel)
	})

	// Find Poll. Exit early if one doesn't exist
	poll, ok := c.ChannelSlice[poll.id]
	if !ok {
		return
	}

	// Remove Poll from ChannelManager map
	delete(c.ChannelSlice, poll.id)
}

var (
	ErrBlockedChannel  = errors.New("channel is blocked")
	ErrChannelNotFound = errors.New("channel was not found")
)
