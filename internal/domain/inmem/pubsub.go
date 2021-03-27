package inmem

import (
	"io"
)

// PubSub represents a service for managing message dispatch and event
// listeners (aka subscriptions).
type PubSub struct {
	subscribers map[int64]io.Writer
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[int64]io.Writer),
	}
}

// Subscribes an user to receive updates from the system
func (p *PubSub) Subscribe(userID int64, w io.Writer) error {
	// TODO: in memory implementation for subscribing users
	return nil
}

// Publishes updates to all the users that subscribed.
func (p * PubSub) Publish(message string) error {
	// TODO: in memory implementation for publishing messages
	return nil
}