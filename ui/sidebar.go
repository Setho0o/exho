package ui

import (
	"image"
	"log"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
)

var lameList = [3]string{"yo", "why", "damn"}

type SideBar struct {
	X        int
	Hovered  bool
	theme    *material.Theme
	SongList layout.List
}

func (sb *SideBar) Draw(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(0.9, func(gtx layout.Context) layout.Dimensions { // Top
			return sb.DisplaySongs(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions { // Bottom
			return sb.Widgets(gtx)
		}),
	)
}

func (sb *SideBar) Events(gtx layout.Context) {
	event.Op(gtx.Ops, sb.Hovered)
	for {
		ev, ok := gtx.Source.Event(pointer.Filter{
			Target: sb.Hovered,
			Kinds:  pointer.Enter | pointer.Leave,
		})
		if !ok {
			break
		}

		if e, ok := ev.(pointer.Event); ok {
			switch e.Kind {
			case pointer.Enter:
				log.Println("enter")
			case pointer.Leave:
				log.Println("leave")
			}
		}
	}
}

func (sb *SideBar) DisplaySongs(gtx layout.Context) layout.Dimensions {
	return sb.SongList.Layout(gtx, 10, func(gtx layout.Context, i int) layout.Dimensions {
		if i >= len(lameList) {
			i = 0
		}
		return material.H5(sb.theme, lameList[i]).Layout(gtx)
	})
}

func (sb *SideBar) Widgets(gtx layout.Context) layout.Dimensions {
	defer clip.Rect{Max: image.Pt(IntDp(sb.X), gtx.Constraints.Max.Y)}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: red}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: image.Pt(IntDp(sb.X), 0)}
}
