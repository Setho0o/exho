package ui

import (
	"fmt"
	"image/color"

	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func Line(ops *op.Ops, offset, val float32) {
	var c clip.Path
	c.Begin(ops)
	c.MoveTo(f32.Pt(500+offset, 810))
	p1 := f32.Pt(500+offset, 810+val)
	c.LineTo(p1)
	c.Close()

	paint.FillShape(ops, color.NRGBA{R: 0x80, A: 0xFF}, clip.Stroke{
		Path:  c.End(),
		Width: 3,
	}.Op())
}

func WaveForm(ops *op.Ops, wave []int) {
	fmt.Print("i")
	for i, e := range wave {
		Line(ops, float32(i+3), float32(e))
	}
}
