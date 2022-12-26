package act

import (
	"encoding/json"
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
func (t Teleport) Do(p *player.Player, complete chan bool) {
	p.Teleport(t.Pos)
	rot := p.Rotation()
	deltaYaw := t.Rot[0] - rot.Yaw()
	deltaPitch := t.Rot[1] - rot.Pitch()
	p.Move(mgl64.Vec3{0, 0, 0}, deltaYaw, deltaPitch)
	complete <- true
}

// FromMap ...
func (t Teleport) FromMap(a map[string]interface{}) Act {
	t.Pos = mgl64.Vec3{a["position"].([]interface{})[0].(float64), a["position"].([]interface{})[1].(float64), a["position"].([]interface{})[2].(float64)}
	t.Rot = [2]float64{a["rotation"].([]interface{})[0].(float64), a["rotation"].([]interface{})[1].(float64)}
	return t
}

// MarshalJSON ...
func (t Teleport) MarshalJSON() ([]byte, error) {
	type a Teleport
	return json.Marshal(struct {
		a
		Type string `json:"type"`
	}{
		a:    a(t),
		Type: t.Type(),
	})
}
