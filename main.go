package main

import (
	"github.com/Setho0o/exho/audio/data"
	"github.com/Setho0o/exho/ui"
)



func main() {
	d := data.Data{}
	ui.Gio(d)
	/*
		ch := make(chan audio.Signal)
		p := audio.PlayerInit(ch)

		p.Play(song)
	*/
}
