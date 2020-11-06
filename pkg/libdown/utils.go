package libdown

import (
	"encoding/hex"
	"fmt"
	"github.com/muesli/gamut"
	"github.com/muesli/termenv"
	"image/color"
	"log"
)

var (
	S            = fmt.Sprintf
	colorProfile = termenv.ColorProfile()
)

func Cry(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ColorToHex(c color.Color) string {
	r, g, b, a := c.RGBA()
	return "#" + hex.EncodeToString([]byte{uint8(r), uint8(g), uint8(b), uint8(a)})
}

func Darker(color string, percent float64) string {
	return ColorToHex(gamut.Darker(gamut.Hex(color), percent))
}

func Lighter(color string, percent float64) string {
	return ColorToHex(gamut.Lighter(gamut.Hex(color), percent))
}

func Colorize(foreground, background string, s ...string) termenv.Style {
	out := termenv.String(s...)
	out = out.Foreground(colorProfile.Color(foreground))
	out = out.Background(colorProfile.Color(background))
	// retrieve color profile supported by terminal
	return out
}

type Stringer interface {
	String() string
}
