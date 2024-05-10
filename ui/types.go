package ui

import (
	"fyne.io/fyne/v2"
	"github.com/badonix/pixelapp/apptype"
	"github.com/badonix/pixelapp/pxcanvas"
	"github.com/badonix/pixelapp/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
