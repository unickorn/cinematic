package act

import (
	"encoding/json"
	"github.com/df-mc/dragonfly/server/player"
)

// Message is an act that sends a message to the player.
type Message struct {
	Message string `json:"message"`
}

// NewMessage creates a new message act.
func NewMessage(message string) Message {
	return Message{Message: message}
}

// Type ...
func (m Message) Type() string {
	return "message"
}

// Do ...
func (m Message) Do(p *player.Player) {
	p.Message(m.Message)
}

// MarshalJSON ...
func (m Message) MarshalJSON() ([]byte, error) {
	type a Message
	return json.Marshal(struct {
		a
		Type string `json:"type"`
	}{
		a:    a(m),
		Type: m.Type(),
	})
}

// FromMap ...
func (m Message) FromMap(a map[string]interface{}) Act {
	m.Message = a["message"].(string)
	return m
}
