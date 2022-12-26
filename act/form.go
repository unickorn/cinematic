package act

import (
	"github.com/df-mc/dragonfly/server/player"
	form "github.com/twistedasylummc/inline-forms"
)

// FormModal is an act that sends a question to the player and awaits a response.
type FormModal struct {
	Title string
	Body  string
	Yes   string
	No    string
	On    func(p *player.Player, r bool)
}

// NewFormModal creates a new form modal act.
func NewFormModal(title, body, yes, no string, on func(p *player.Player, r bool)) FormModal {
	return FormModal{
		Title: title,
		Body:  body,
		Yes:   yes,
		No:    no,
		On:    on,
	}
}

// Type ...
func (f FormModal) Type() string {
	return "form"
}

// Do ...
func (f FormModal) Do(p *player.Player, complete chan bool) {
	p.SendForm(&form.Modal{
		Title:   f.Title,
		Content: f.Body,
		Button1: form.Button{
			Text: f.Yes,
			Submit: func() {
				f.On(p, true)
				complete <- true
			},
		},
		Button2: form.Button{
			Text: f.No,
			Submit: func() {
				f.On(p, false)
				complete <- true
			},
		},
	})
}

// FromMap ...
func (f FormModal) FromMap(m map[string]interface{}) Act {
	f.Title = m["title"].(string)
	f.Body = m["body"].(string)
	f.Yes = m["yes"].(string)
	f.No = m["no"].(string)
	return f
}
