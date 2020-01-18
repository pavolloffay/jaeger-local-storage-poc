package model

import (
	"time"
)

// Event The event to be stored
type Event interface {

	// Tenant The optional tenant associated with the event
	Tenant() string

	// Type The optional type associated with the event
	Type() string

	// Timestamp The optional timestamp, will use current time if not defined
	Timestamp() time.Time

	// Fields The list of fields to be indexed for the event
	Fields() []*Field
}
