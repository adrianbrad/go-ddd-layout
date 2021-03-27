package domain

import (
	"io"
)

// PubSub represents a service for managing message dispatch and event
// listeners (aka subscriptions).
type PubSub interface {
	// Subscribes an user to receive updates from the system
	Subscribe(id int64, w io.Writer) error

	// Publishes updates to all the users that subscribed.
	Publish(message string) error
}
