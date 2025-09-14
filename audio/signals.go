package audio

import (
	"fmt"
)

type Signal int

const (
	Skip Signal = iota
	Previous
	Pause

	VolUp
	VolDown

	Restart
)

func (p *Player) CheckSignals() {
	select {
	case x := <-p.ch:
		switch x {
		case Skip:
			p.Skip()
		case Previous:
			p.Previous()
		case Pause:
			p.Pause()

		case VolUp:
			p.VolUp()
		case VolDown:
			p.VolDown()

		case Restart:
			p.Restart()
		}
	default:
	}
}
func (p *Player) Skip() {
}
func (p *Player) Pause() {
}

func (p *Player) Previous() {
}
func (p *Player) VolUp() {
	p.player.SetVolume(p.player.Volume() + .1)
}

func (p *Player) VolDown() {
	p.player.SetVolume(p.player.Volume() - .1)
	fmt.Println("voldown")
}

func (p *Player) Restart() {
}
