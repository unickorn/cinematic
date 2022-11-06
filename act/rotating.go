package act

import (
	"encoding/json"
	"github.com/cnkei/gospline"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"time"
)

// RotatingPath is a path that rotates the player as well, on top of moving them.
type RotatingPath struct {
	Points   [][5]float64  `json:"points"`
	Duration time.Duration `json:"duration"`
	Interval time.Duration `json:"interval"`

	splines [][]float64
}

// NewRotatingPath creates a new rotating path from the given points. The points are given as a slice of
// 5-arrays, where the first 3 values are the x, y and z coordinates, and the last 2 values are the yaw and the pitch.
func NewRotatingPath(points [][5]float64, duration, interval time.Duration) RotatingPath {
	return RotatingPath{
		Points:   points,
		Duration: duration,
		Interval: interval,
	}.initialize()
}

// initialize ...
func (p RotatingPath) initialize() RotatingPath {
	x := make([]float64, len(p.Points))
	y := make([]float64, len(p.Points))
	z := make([]float64, len(p.Points))
	yaw := make([]float64, len(p.Points))
	pitch := make([]float64, len(p.Points))
	t := make([]float64, len(p.Points))
	step := p.Duration / time.Duration(float64(len(p.Points)))
	for i, p := range p.Points {
		x[i] = p[0]
		y[i] = p[1]
		z[i] = p[2]
		yaw[i] = p[3]
		pitch[i] = p[4]
		t[i] = step.Seconds() * float64(i)
	}
	p.splines = [][]float64{
		gospline.NewCubicSpline(t, x).Range(0, p.Duration.Seconds(), p.Interval.Seconds()),
		gospline.NewCubicSpline(t, y).Range(0, p.Duration.Seconds(), p.Interval.Seconds()),
		gospline.NewCubicSpline(t, z).Range(0, p.Duration.Seconds(), p.Interval.Seconds()),
		gospline.NewCubicSpline(t, yaw).Range(0, p.Duration.Seconds(), p.Interval.Seconds()),
		gospline.NewCubicSpline(t, pitch).Range(0, p.Duration.Seconds(), p.Interval.Seconds()),
	}
	return p
}

// at returns the position at the given time.
func (p RotatingPath) at(i int) mgl32.Vec3 {
	return mgl32.Vec3{float32(p.splines[0][i]), float32(p.splines[1][i]), float32(p.splines[2][i])}
}

// rotationAt returns the rotation at the given time.
func (p RotatingPath) rotationAt(i int) (float32, float32) {
	return float32(p.splines[3][i]), float32(p.splines[4][i])
}

// Type ...
func (p RotatingPath) Type() string {
	return "rotating"
}

// MarshalJSON ...
func (p RotatingPath) MarshalJSON() ([]byte, error) {
	type a RotatingPath
	return json.Marshal(&struct {
		a
		Type string `json:"type"`
	}{
		a:    a(p),
		Type: p.Type(),
	})
}

// Do moves the player along the path.
func (p RotatingPath) Do(pl *player.Player) {
	s := player_session(pl)
	steps := int(p.Duration / p.Interval)
	for i := 0; i < steps; i++ {
		yaw, pitch := p.rotationAt(i)
		session_writePacket(s, &packet.MovePlayer{
			EntityRuntimeID: 1,
			Position:        p.at(i),
			Pitch:           pitch,
			Yaw:             yaw,
			HeadYaw:         yaw,
			Mode:            packet.MoveModeTeleport,
			OnGround:        pl.OnGround(),
		})
		time.Sleep(p.Interval)
	}
}

// FromMap ...
func (p RotatingPath) FromMap(m map[string]interface{}) Act {
	p.Points = m["points"].([][5]float64)
	p.Duration = m["duration"].(time.Duration)
	p.Interval = m["interval"].(time.Duration)
	return p.initialize()
}
