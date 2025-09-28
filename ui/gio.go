package ui

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/Setho0o/exho/audio/data"
)

func Gio(d data.Data) {
	go func(d data.Data) {
		window := new(app.Window)
		err := run(window, d)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}(d)
	app.Main()
}

func run(window *app.Window, d data.Data) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e) //graphics ctx for rendering state

			Flex(gtx)

			//wave(&ops)
			//	WaveForm(&ops, wave)
			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}

func Flex(gtx layout.Context) layout.Dimensions {
	return layout.Flex{}.Layout(gtx, // left side bar
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(IntDp(200), gtx.Constraints.Max.Y), red)
		}),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions { // main scene
			return ColorBox(gtx, gtx.Constraints.Min, blue)
		}),
	)
}
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
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
