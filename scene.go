package cinematic

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/unickorn/cinematic/act"
)

// Scene is a collection of acts that are performed in sequence.
type Scene struct {
	// Name is the name of the scene.
	Name string `json:"name"`
	// Actions is a slice of actions in order.
	Actions []act.Act `json:"actions"`
}

var emptyScene = Scene{}

// NewScene creates a new scene with the given name.
func NewScene(name string) Scene {
	return Scene{
		Name:    name,
		Actions: []act.Act{},
	}
}

// WithActions returns a new scene with the given actions.
func (s Scene) WithActions(actions []act.Act) Scene {
	s.Actions = actions
	return s
}

// AddAction adds an action to the scene.
func (s Scene) AddAction(a act.Act) Scene {
	s.Actions = append(s.Actions, a)
	return s
}

// RemoveAction removes the action from the scene.
func (s Scene) RemoveAction(a act.Act) Scene {
	for i, action := range s.Actions {
		if action == a {
			s.Actions = append(s.Actions[:i], s.Actions[i+1:]...)
		}
	}
	return s
}

// Play plays the scene for the given player.
func (s Scene) Play(p *player.Player) {
	// play actions with blocking in between
	c := make(chan bool)
	for _, a := range s.Actions {
		go a.Do(p, c)
		if <-c == false {
			return
		}
	}
}
