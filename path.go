package cinematic

import (
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
type Path struct {
	points   []mgl64.Vec3
	duration time.Duration
	interval time.Duration

	splines [][]float64
}

// NewPath creates a new path from the given points, with the given duration and interval.
func NewPath(points []mgl64.Vec3, duration, interval time.Duration) *Path {
	x := make([]float64, len(points))
	y := make([]float64, len(points))
	z := make([]float64, len(points))
	t := make([]float64, len(points))
	step := duration / time.Duration(float64(len(points)))
	for i, p := range points {
		x[i] = p.X()
		y[i] = p.Y()
		z[i] = p.Z()
		t[i] = step.Seconds() * float64(i)
	}
	return &Path{
		points:   points,
		duration: duration,
		interval: interval,
		splines: [][]float64{
			gospline.NewCubicSpline(t, x).Range(0, duration.Seconds(), interval.Seconds()),
			gospline.NewCubicSpline(t, y).Range(0, duration.Seconds(), interval.Seconds()),
			gospline.NewCubicSpline(t, z).Range(0, duration.Seconds(), interval.Seconds()),
		},
	}
}

// at returns the position at the given time.
func (p *Path) at(i int) mgl32.Vec3 {
	return mgl32.Vec3{float32(p.splines[0][i]), float32(p.splines[1][i]), float32(p.splines[2][i])}
}

// Move moves the player along the path.
func (p *Path) Move(pl *player.Player) {
	s := player_session(pl)
	steps := int(p.duration / p.interval)
	for i := 0; i < steps; i++ {
		session_writePacket(s, &packet.MovePlayer{
			EntityRuntimeID: 1,
			Position:        p.at(i),
			Mode:            packet.MoveModeNormal,
			OnGround:        pl.OnGround(),
		})
		time.Sleep(p.interval)
	}
}

//go:linkname player_session github.com/df-mc/dragonfly/server/player.(*Player).session
//noinspection ALL
func player_session(*player.Player) *session.Session

//go:linkname session_writePacket github.com/df-mc/dragonfly/server/session.(*Session).writePacket
//noinspection ALL
func session_writePacket(*session.Session, packet.Packet)
