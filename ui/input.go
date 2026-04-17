package ui

import (
	"log"

	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"github.com/Setho0o/exho/audio/data"
)

func Input(gtx layout.Context, d *data.Data) {
	defer clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops).Pop()

	for {
		ev, ok := gtx.Event(key.Filter{
			Name: "Q",
		})

		if !ok {
			break
		}

		if k, ok := ev.(key.Event); ok {
			if k.Name == "Q" {
				log.Fatal("exit")
			}
		}
	}
}
