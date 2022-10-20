package cinematic

import (
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
func NewRotatingPath(points [][5]float64, duration, interval time.Duration) *RotatingPath {
	x := make([]float64, len(points))
	y := make([]float64, len(points))
	z := make([]float64, len(points))
	yaw := make([]float64, len(points))
	pitch := make([]float64, len(points))
	t := make([]float64, len(points))
	step := duration / time.Duration(float64(len(points)))
	for i, p := range points {
		x[i] = p[0]
		y[i] = p[1]
		z[i] = p[2]
		yaw[i] = p[3]
		pitch[i] = p[4]
		t[i] = step.Seconds() * float64(i)
	}
	return &RotatingPath{
		Points:   points,
		Duration: duration,
		Interval: interval,
		splines: [][]float64{
			gospline.NewCubicSpline(t, x).Range(0, duration.Seconds(), interval.Seconds()),
			gospline.NewCubicSpline(t, y).Range(0, duration.Seconds(), interval.Seconds()),
			gospline.NewCubicSpline(t, z).Range(0, duration.Seconds(), interval.Seconds()),
			gospline.NewCubicSpline(t, yaw).Range(0, duration.Seconds(), interval.Seconds()),
			gospline.NewCubicSpline(t, pitch).Range(0, duration.Seconds(), interval.Seconds()),
		},
	}
}

// at returns the position at the given time.
func (p *RotatingPath) at(i int) (mgl32.Vec3, float32, float32) {
	return mgl32.Vec3{float32(p.splines[0][i]), float32(p.splines[1][i]), float32(p.splines[2][i])}, float32(p.splines[3][i]), float32(p.splines[4][i])
}

// Move moves the player along the path.
func (p *RotatingPath) Move(pl *player.Player) {
	s := player_session(pl)
	steps := int(p.Duration / p.Interval)
	for i := 0; i < steps; i++ {
		pos, yaw, pitch := p.at(i)
		session_writePacket(s, &packet.MovePlayer{
			EntityRuntimeID: 1,
			Position:        pos,
			Pitch:           pitch,
			Yaw:             yaw,
			HeadYaw:         yaw,
			Mode:            packet.MoveModeTeleport,
			OnGround:        pl.OnGround(),
		})
		time.Sleep(p.Interval)
	}
}
