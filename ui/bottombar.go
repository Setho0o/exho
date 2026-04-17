package ui

import (
	"image"

	"gioui.org/layout"
)

type BottomBar struct{}

func (b *BottomBar) Draw(gtx layout.Context) layout.Dimensions {
	return ColorBox(gtx, image.Pt(gtx.Constraints.Max.X, IntDp(200)), green)
}
