package ui

import "gioui.org/layout"

type Center struct{}

func (c *Center) Draw(gtx layout.Context) layout.Dimensions {
	return ColorBox(gtx, gtx.Constraints.Min, blue)
}
