package act

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
)

// Teleport is an act that teleports the player to a position.
type Teleport struct {
	Pos mgl64.Vec3 `json:"position"`
	Rot [2]float64 `json:"rotation"`
}

// NewTeleport creates a new teleport act.
func NewTeleport(pos mgl64.Vec3, rot [2]float64) Teleport {
	return Teleport{Pos: pos, Rot: rot}
}

// Type ...
func (t Teleport) Type() string {
	return "teleport"
}

// Do ...
func (t Teleport) Do(p *player.Player) {
	p.Teleport(t.Pos)
	yaw, pitch := p.Rotation()
	deltaYaw := t.Rot[0] - yaw
	deltaPitch := t.Rot[1] - pitch
	p.Move(mgl64.Vec3{0, 0, 0}, deltaYaw, deltaPitch)
}

// FromMap ...
func (t Teleport) FromMap(a map[string]interface{}) Act {
	t.Pos = mgl64.Vec3{a["position"].([]interface{})[0].(float64), a["position"].([]interface{})[1].(float64), a["position"].([]interface{})[2].(float64)}
	t.Rot = [2]float64{a["rotation"].([]interface{})[0].(float64), a["rotation"].([]interface{})[1].(float64)}
	return t
}
