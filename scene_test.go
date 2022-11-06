package cinematic

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sirupsen/logrus"
	"github.com/unickorn/cinematic/act"
	"testing"
	"time"
)

// TestScene ...
func TestScene(t *testing.T) {
	// Create a new path
	p := act.NewRotatingPath([][5]float64{{0, -60, 0, 0, 0}, {10, -50, 10, 180, -90}, {20, -59, 20, 0, 0}}, 10*time.Second, time.Second/20)
	// New scene using the path
	scene := NewScene("test").WithActions(map[int]act.Act{
		0:     p,
		1000:  act.NewMessage("Woooooahh"),
		5000:  act.NewMessage("We're halfway theeere"),
		7500:  act.NewMessage("Woooo hooo"),
		10000: act.NewMessage("Living on a prayer"),
	})

	// start a dragonfly server
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	c, _ := server.DefaultConfig().Config(logrus.StandardLogger())
	s := c.New()
	s.CloseOnProgramEnd()
	s.Listen()

	s.Accept(func(p *player.Player) {
		p.ShowCoordinates()
		go func() {
			for {
				scene.Play(p)
			}
		}()
	})
	select {}
}
