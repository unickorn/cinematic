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
	// New scene using the path
	scene := NewScene("test").WithActions([]act.Act{
		act.NewMessage("§bWelcome to the cinematic test scene"),
		act.NewDelay(time.Second * 2),
		act.NewMessage("§eThis is a test scene"),
		act.NewDelay(time.Second),
		act.NewMessage("Your feedback will be appreciated in 3 seconds."),
		act.NewDelay(time.Second * 3),
		act.NewFormModal("Feedback form", "How were things?", "Great", "Terrible", func(p *player.Player, r bool) {
			if r {
				p.Message("We are glad you liked it!")
			} else {
				p.Message("It's okay, we will try to improve.")
			}
		}),
		act.NewDelay(time.Second),
		act.NewMessage("Thank you for your input!"),
	})

	// start a dragonfly server
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	c, err := server.DefaultConfig().Config(logrus.StandardLogger())
	if err != nil {
		panic(err)
	}
	s := c.New()
	s.CloseOnProgramEnd()
	s.Listen()

	s.Accept(func(p *player.Player) {
		p.ShowCoordinates()
		go func() {
			scene.Play(p)
			p.Message("§cScene finished")
		}()
	})
	select {}
}
