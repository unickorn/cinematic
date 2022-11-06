package cinematic

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/unickorn/cinematic/act"
	"sort"
	"time"
)

// Scene is a collection of acts that are performed in sequence.
type Scene struct {
	// Name is the name of the scene.
	Name string `json:"name"`
	// Actions is a map of actions with their timestamps in milliseconds.
	Actions map[int]act.Act `json:"actions"`
}

var emptyScene = Scene{}

// NewScene creates a new scene with the given name.
func NewScene(name string) Scene {
	return Scene{
		Name:    name,
		Actions: map[int]act.Act{},
	}
}

// WithActions returns a new scene with the given actions.
func (s Scene) WithActions(actions map[int]act.Act) Scene {
	s.Actions = actions
	return s
}

// AddAction adds an action to the scene at the given timestamp.
func (s Scene) AddAction(timestamp int, a act.Act) Scene {
	s.Actions[timestamp] = a
	return s
}

// RemoveAction removes the action at the given timestamp from the scene.
func (s Scene) RemoveAction(timestamp int) Scene {
	delete(s.Actions, timestamp)
	return s
}

// Play plays the scene for the given player.
func (s Scene) Play(p *player.Player) {
	// sort actions by timestamp
	var timestamps []int
	for timestamp := range s.Actions {
		timestamps = append(timestamps, timestamp)
	}
	sort.Ints(timestamps)

	lastTimestamp := 0
	// play actions
	for _, timestamp := range timestamps {
		time.Sleep(time.Duration(timestamp-lastTimestamp) * time.Millisecond)
		go s.Actions[timestamp].Do(p)
		lastTimestamp = timestamp
	}
}
