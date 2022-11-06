package act

import "github.com/df-mc/dragonfly/server/player"

// Act is an interface for the possible actions that make up a cinematic scene.
type Act interface {
	// Type returns the act type.
	Type() string
	// Do performs the action on the given player.
	Do(p *player.Player)
	// FromMap sets the act's fields from the given map.
	FromMap(m map[string]interface{}) Act
}

var registeredActs = map[string]func() Act{}

// Register registers an act with the given type. The function passed should return a new instance of the
// act.
func Register(act func() Act) {
	registeredActs[act().Type()] = act
}

// New returns a new instance of the act with the given type.
func New(t string) Act {
	return registeredActs[t]()
}

func init() {
	Register(func() Act { return RotatingPath{} })
	Register(func() Act { return NormalPath{} })
	Register(func() Act { return Message{} })
	Register(func() Act { return Teleport{} })
}
