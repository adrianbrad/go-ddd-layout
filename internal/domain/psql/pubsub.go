package psql

import (
	"github.com/adrianbrad/ddd-layout/internal/domain"
	"github.com/lib/pq"
	"io"
)

// Ensure the PubSub implemented in this package satisfies de PubSub interface from the domain.
var _ domain.PubSub = (*PubSub)(nil)

var conninfo string = "dbname=exampledb user=webapp password=webapp"

// PubSub represents a service for managing message dispatch and event
// listeners (aka subscriptions).
type PubSub struct {
	pq.Listener
}

func NewPubSub() *PubSub {
	return &PubSub{}
}

// Publishes updates to all the users that subscribed.
func (p *PubSub) Publish(message string) error {
	// TODO: Postgres implementation for publishing messages
	return nil
}

// Subscribes an user to receive updates from the system
func (p *PubSub) Subscribe(userID int64, sub io.Writer) error {
	// TODO: PostgreSQL implementation for subscribing a user
	return nil
}