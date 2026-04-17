package ui

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Setho0o/exho/audio/data"
)

var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

// just spawning the ui in a go func to run it concurently with our audio
func Gio(d *data.Data) {
	go func(d *data.Data) {
		window := new(app.Window)
		err := run(window, d)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}(d)
	app.Main()
}

func run(window *app.Window, d *data.Data) error {
	var ops op.Ops
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	c := Center{}
	bb := BottomBar{}
	sb := SideBar{
		theme: th,
	}
	sb.X = 200

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e) //graphics ctx for rendering state

			Flex(gtx, sb.Draw, c.Draw, bb.Draw) // all the visual rendering
			Input(gtx, d)

			//wave(&ops)
			//	WaveForm(&ops, wave)

			// Pass the drawing operations to the GPU.

			e.Frame(gtx.Ops)
		}
	}
}

func Flex(gtx layout.Context, leftbar, center, bottombar func(layout.Context) layout.Dimensions) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions { // Center
					return center(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions { // Left
					return leftbar(gtx)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { // Bottom
			return bottombar(gtx)
		}),
	)
}

// need to just make a dedicated utils folder for these, and add a lot of helpfull audio funcs
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func IntDp(d int) int {
	return int(unit.Dp(d))
}
