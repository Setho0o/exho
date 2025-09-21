package ui

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func Gio(wave []int) {
	go func(wave []int) {
		window := new(app.Window)
		err := run(window, wave)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}(wave)
	app.Main()
}

func run(window *app.Window, wave []int) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e) //graphics ctx for rendering state
			//wave(&ops)
			WaveForm(&ops, wave)
			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}

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
func drawRedRect(ops *op.Ops, p1, p2 image.Point) {
	defer clip.Rect{
		Min: p1,
		Max: p2,
	}.Push(ops).Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func IntDp(d int) int {
	return int(unit.Dp(d))
}
