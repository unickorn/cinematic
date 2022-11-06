package act

import (
	"encoding/json"
	"github.com/cnkei/gospline"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"time"
	_ "unsafe"
)

// Path is a path along which a player can move.
type Path interface {
	at(int) mgl32.Vec3
	initialize()
	Do(*player.Player)
}

// NormalPath is a path only concerning x, y and z coordinates.
type NormalPath struct {
	Points   []mgl64.Vec3  `json:"points"`
	Duration time.Duration `json:"duration"`
	Interval time.Duration `json:"interval"`

	splines [][]float64
}

// NewPath creates a new path from the given points, with the given duration and interval.
func NewPath(points []mgl64.Vec3, duration, interval time.Duration) NormalPath {
	return NormalPath{
		Points:   points,
		Duration: duration,
		Interval: interval,
	}.initialize()
}

// initialize creates the splines for the path. They are not stored in JSON files to save storage.
func (p NormalPath) initialize() NormalPath {
	x := make([]float64, len(p.Points))
	y := make([]float64, len(p.Points))
	z := make([]float64, len(p.Points))
	t := make([]float64, len(p.Points))
	step := p.Duration / time.Duration(float64(len(p.Points)-1))
	for i, p := range p.Points {
		x[i] = p.X()
		y[i] = p.Y()
		z[i] = p.Z()
		t[i] = step.Seconds() * float64(i)
	}
	p.splines = [][]float64{
		gospline.NewCubicSpline(t, x).Range(0, p.Duration.Seconds(), p.Interval.Seconds()),
		gospline.NewCubicSpline(t, y).Range(0, p.Duration.Seconds(), p.Interval.Seconds()),
		gospline.NewCubicSpline(t, z).Range(0, p.Duration.Seconds(), p.Interval.Seconds()),
	}
	return p
}

// at returns the position at the given time.
func (p NormalPath) at(i int) mgl32.Vec3 {
	return mgl32.Vec3{float32(p.splines[0][i]), float32(p.splines[1][i]), float32(p.splines[2][i])}
}

// Type ...
func (p NormalPath) Type() string {
	return "path"
}

// Do moves the player along the path.
func (p NormalPath) Do(pl *player.Player) {
	s := player_session(pl)
	steps := int(p.Duration / p.Interval)
	for i := 0; i < steps; i++ {
		session_writePacket(s, &packet.MovePlayer{
			EntityRuntimeID: 1,
			Position:        p.at(i),
			Mode:            packet.MoveModeNormal,
			OnGround:        pl.OnGround(),
		})
		time.Sleep(p.Interval)
	}
}

// MarshalJSON ...
func (p NormalPath) MarshalJSON() ([]byte, error) {
	type a NormalPath
	return json.Marshal(&struct {
		a
		Type string `json:"type"`
	}{
		a:    a(p),
		Type: p.Type(),
	})
}

// FromMap ...
func (p NormalPath) FromMap(m map[string]interface{}) Act {
	points := m["points"].([]interface{})
	p.Points = make([]mgl64.Vec3, len(points))
	for i, point := range points {
		p.Points[i] = mgl64.Vec3{point.([]interface{})[0].(float64), point.([]interface{})[1].(float64), point.([]interface{})[2].(float64)}
	}
	p.Duration = time.Duration(m["duration"].(float64))
	p.Interval = time.Duration(m["interval"].(float64))
	return p.initialize()
}

//go:linkname player_session github.com/df-mc/dragonfly/server/player.(*Player).session
//noinspection ALL
func player_session(*player.Player) *session.Session

//go:linkname session_writePacket github.com/df-mc/dragonfly/server/session.(*Session).writePacket
//noinspection ALL
func session_writePacket(*session.Session, packet.Packet)
