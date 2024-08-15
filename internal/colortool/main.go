package main

import (
	"fmt"
	"html/template"
	"image/color"
	"log"
	"os"
)

func toU8(v float32) uint8 {
	n := min(int(v*256), 255)
	return uint8(n)
}

type Color struct {
	Y  float32
	Cr float32
	Cb float32
}

func (c Color) toRGB() (r, g, b uint8) {
	return color.YCbCrToRGB(toU8(c.Y), toU8(c.Cb), toU8(c.Cr))
}

func (c Color) CSS() string {
	r, g, b := c.toRGB()
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// Colors provide TMS9918A's color pallet. These values are came from TMS9918
// datasheet.
//
// It is bit differ between TMS9918 datasheet and the wikipedia entry.
//
// TMS9918 datasheet: https://www.cs.columbia.edu/~sedwards/papers/TMS9918.pdf
// Wikipedia: https://en.wikipedia.org/wiki/List_of_8-bit_computer_hardware_graphics#Systems_based_on_the_Texas_Instruments_TMS9918_chip
var Colors = []Color{
	{Y: 0.00, Cr: 0.00, Cb: 0.00},
	{Y: 0.00, Cr: 0.47, Cb: 0.47},
	{Y: 0.53, Cr: 0.07, Cb: 0.20},
	{Y: 0.67, Cr: 0.17, Cb: 0.27},

	{Y: 0.40, Cr: 0.40, Cb: 1.00},
	{Y: 0.53, Cr: 0.43, Cb: 0.93},
	{Y: 0.47, Cr: 0.83, Cb: 0.30},
	{Y: 0.73, Cr: 0.00, Cb: 0.70},

	{Y: 0.53, Cr: 0.93, Cb: 0.27},
	{Y: 0.67, Cr: 0.93, Cb: 0.27},
	{Y: 0.73, Cr: 0.57, Cb: 0.07},
	{Y: 0.80, Cr: 0.57, Cb: 0.17},

	{Y: 0.47, Cr: 0.13, Cb: 0.23},
	{Y: 0.53, Cr: 0.73, Cb: 0.67},
	{Y: 0.80, Cr: 0.47, Cb: 0.47},
	{Y: 1.00, Cr: 0.47, Cb: 0.47},
}

var tmpl = template.Must(template.New("new").Parse(`<body>
<div class="color-list">{{ range $i, $c := . }}
  <div class="color" style="background-color: {{ $c.CSS }}">
    <span>{{ $i }}</span><br>
    <span>{{ $c.CSS }}</span><br>
  </div>{{ end }}
</div>
</body>`))

func main() {
	err := tmpl.Execute(os.Stdout, Colors)
	if err != nil {
		log.Fatal(err)
	}
}
