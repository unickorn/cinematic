package act

import (
	"encoding/json"
	"github.com/df-mc/dragonfly/server/player"
	"time"
)

// Delay is an act that delays the scene for a specified amount of time.
type Delay struct {
	Length time.Duration `json:"length"`
}

// NewDelay creates a new delay act.
func NewDelay(length time.Duration) Delay {
	return Delay{Length: length}
}

// Type ...
func (d Delay) Type() string {
	return "delay"
}

// Do ...
func (d Delay) Do(_ *player.Player, complete chan bool) {
	time.Sleep(d.Length)
	complete <- true
}

// FromMap ...
func (d Delay) FromMap(m map[string]interface{}) Act {
	d.Length = time.Duration(m["length"].(float64))
	return d
}

// MarshalJSON ...
func (d Delay) MarshalJSON() ([]byte, error) {
	type a Delay
	return json.Marshal(struct {
		a
		Type string `json:"type"`
	}{
		a:    a(d),
		Type: d.Type(),
	})
}
