package cinematic

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"time"
	_ "unsafe"
)

// Move moves the player with the Splines given.
func Move(p *player.Player, splines *Splines, duration time.Duration, interval time.Duration) {
	s := player_session(p)
	steps := int(duration / interval)
	for i := 0; i < steps; i++ {
		session_writePacket(s, &packet.MovePlayer{
			EntityRuntimeID: 1,
			Position:        splines.At(i),
			Mode:            packet.MoveModeNormal,
			OnGround:        p.OnGround(),
		})
		time.Sleep(interval)
	}
}

// MoveRotation moves the player with the RotationSplines given.
func MoveRotation(p *player.Player, splines *RotationSplines, duration time.Duration, interval time.Duration) {
	s := player_session(p)
	steps := int(duration / interval)
	for i := 0; i < steps; i++ {
		session_writePacket(s, &packet.MovePlayer{
			EntityRuntimeID: 1,
			Position:        splines.At(i),
			Pitch:           splines.Pitch(i),
			Yaw:             splines.Yaw(i),
			HeadYaw:         splines.Yaw(i),
			Mode:            packet.MoveModeTeleport,
			OnGround:        p.OnGround(),
		})
		time.Sleep(interval)
	}
}

//go:linkname player_session github.com/df-mc/dragonfly/server/player.(*Player).session
//noinspection ALL
func player_session(*player.Player) *session.Session

//go:linkname session_writePacket github.com/df-mc/dragonfly/server/session.(*Session).writePacket
//noinspection ALL
func session_writePacket(*session.Session, packet.Packet)
