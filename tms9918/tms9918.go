package tms9918

import (
	"errors"
	"image"
	"image/color"
)

type VDP struct {
	Register0 Register
	Register1 Register
	Register2 Register
	Register3 Register
	Register4 Register
	Register5 Register
	Register6 Register
	Register7 Register

	StatusRegster Register

	VRAM []uint8
}

type Register uint8

func (r Register) GetBit(n int) bool {
	return r&(0x80>>n) != 0
}

func (r *Register) SetBit(n int) {
	*r |= (0x80 >> n)
}

func (r *Register) ResetBit(n int) {
	*r &= ^(0x80 >> n)
}

func (r *Register) ResetSet(reset, set uint8) {
	*r = *r&^Register(reset) | Register(set)
}

func New(vram []uint8) *VDP {
	if len(vram) < 4090 {
		panic("VRAM sizes less than 4K (4096) are not supported")
	}
	vdp := &VDP{VRAM: vram}
	if len(vram) >= 16386 {
		vdp.Register1.SetBit(0)
	}
	return vdp
}

var Palette = color.Palette{
	color.RGBA{0x00, 0x00, 0x00, 0x00}, // 0: TRANSPARENT
	color.RGBA{0x00, 0x08, 0x00, 0xFF}, // 1: BLACK
	color.RGBA{0x00, 0xF1, 0x00, 0xFF}, // 2: MEDIUM GREEN
	color.RGBA{0x34, 0xFC, 0x43, 0xFF}, // 3: LIGHT GREEN
	color.RGBA{0x41, 0x4D, 0xFF, 0xFF}, // 4: DARK BLUE
	color.RGBA{0x6E, 0x6E, 0xFF, 0xFF}, // 5: LIGHT BLUE
	color.RGBA{0xEE, 0x4E, 0x1C, 0xFF}, // 6: DARK RED
	color.RGBA{0x07, 0xFF, 0xFF, 0xFF}, // 7: CYAN
	color.RGBA{0xFF, 0x4D, 0x1E, 0xFF}, // 8: MEDIUM RED
	color.RGBA{0xFF, 0x71, 0x43, 0xFF}, // 9: LIGHT RED
	color.RGBA{0xD2, 0xD4, 0x00, 0xFF}, // A: DARK YELLOW
	color.RGBA{0xE4, 0xDD, 0x36, 0xFF}, // B: LIGHT YELLOW
	color.RGBA{0x00, 0xD4, 0x00, 0xFF}, // C: DARK GREEN
	color.RGBA{0xD8, 0x4F, 0xD3, 0xFF}, // D: MAGENTA
	color.RGBA{0xC1, 0xD5, 0xBE, 0xFF}, // E: GRAY
	color.RGBA{0xF4, 0xFF, 0xF1, 0xFF}, // F: WHITE
}

func (v *VDP) NewPaletted(rect image.Rectangle) (*image.Paletted, error) {
	if rect.Dx() < 256 {
		return nil, errors.New("width must be at least 256")
	}
	if rect.Dy() < 192 {
		return nil, errors.New("height must be at least 192")
	}
	return image.NewPaletted(rect, Palette), nil
}

type Mode int

const (
	Graphics1Mode Mode = iota
	Graphics2Mode
	MulticolorMode
	TextMode
)

func (v *VDP) mode() Mode {
	if v.Register0.GetBit(0) {
		return Graphics2Mode
	}
	if v.Register1.GetBit(4) {
		return MulticolorMode
	}
	if v.Register1.GetBit(5) {
		return TextMode
	}
	return Graphics1Mode
}

type baseAddresses struct {
	nameTable        int
	patternGenerator int
	colorTable       int
}

func (v *VDP) backdropColorIndex() uint8 {
	return uint8(v.Register7 & 0x0F)
}

func (v *VDP) textColorIndex() uint8 {
	return uint8(v.Register7&0xF0) >> 4
}

func (v *VDP) nameTableBaseAddress() int {
	n := int(v.Register2 & 0x0F)
	return n * 0x0400
}

func (v *VDP) patternGeneratorBaseAddress() int {
	n := int(v.Register4 & 0x07)
	return n * 0x0800
}

func (v *VDP) colorTableBaseAddress() int {
	n := int(v.Register3)
	return n * 0x0040
}

func (v *VDP) baseAddresses() baseAddresses {
	return baseAddresses{
		nameTable:        v.nameTableBaseAddress(),
		patternGenerator: v.patternGeneratorBaseAddress(),
		colorTable:       v.colorTableBaseAddress(),
	}
}

type spriteParams struct {
	size        int
	psize       int
	mag         int
	attrAddr    int
	patternAddr int
}

func (v *VDP) spriteSize() int {
	if v.Register1.GetBit(6) {
		return 16
	}
	return 8
}

func (v *VDP) spritePatterSize() int {
	if v.Register1.GetBit(6) {
		return 32
	}
	return 8
}

func (v *VDP) spriteMag() int {
	if v.Register1.GetBit(6) {
		return 2
	}
	return 1
}

func (v *VDP) spriteParams() spriteParams {
	return spriteParams{
		size:        v.spriteSize(),
		psize:       v.spritePatterSize(),
		mag:         v.spriteMag(),
		attrAddr:    int(v.Register5&0x7F) * 0x0080,
		patternAddr: int(v.Register6&0x07) * 0x0800,
	}
}

func (v *VDP) Render(scr *image.Paletted) {
	v.renderBackdropPlane(scr)
	switch v.mode() {
	case Graphics1Mode:
		v.renderGraphics1Plane(scr)
	case Graphics2Mode:
		v.renderGraphics2Plane(scr)
	case MulticolorMode:
		v.renderMulticolorPlane(scr)
	case TextMode:
		v.renderTextPlane(scr)
	}
	v.renderSpritesAll(scr)
}

func (v *VDP) calcRenderArea(scr *image.Paletted) image.Rectangle {
	h := scr.Rect.Dy()
	w := scr.Rect.Dx()
	startY := (h - 192) / 2
	endY := startY + 192
	startX := (w - 256) / 2
	endX := startX + 256
	return image.Rect(startX, startY, endX, endY)
}

func (v *VDP) renderBackdropPlane(scr *image.Paletted) {
	r := v.calcRenderArea(scr)
	w := scr.Rect.Dx()
	h := scr.Rect.Dy()
	b := v.backdropColorIndex()
	for y := 0; y < r.Min.Y; y++ {
		for x := 0; x < w; x++ {
			scr.SetColorIndex(x, y, b)
		}
	}
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := 0; x < r.Min.X; x++ {
			scr.SetColorIndex(x, y, b)
		}
		for x := r.Max.X; x < w; x++ {
			scr.SetColorIndex(x, y, b)
		}
	}
	for y := r.Max.Y; y <= h; y++ {
		for x := 0; x < w; x++ {
			scr.SetColorIndex(x, y, b)
		}
	}
}

func (v *VDP) renderGraphics1Plane(scr *image.Paletted) {
	r := v.calcRenderArea(scr)
	offX, offY := r.Min.X, r.Min.Y
	addr := v.baseAddresses()
	for cy := 0; cy < 24; cy++ {
		baseY := offY + cy*8
		for cx := 0; cx < 32; cx++ {
			baseX := offX + cx*8
			m := v.VRAM[addr.nameTable+cx+cy*32]
			p := addr.patternGenerator + int(m)*8
			pattern := v.VRAM[p : p+8]
			cv := v.VRAM[addr.colorTable+int(m)/8]
			color0, color1 := uint8(cv&0x0F), uint8(cv&0xF0>>4)
			for y := range 8 {
				p := pattern[y]
				for x := range 8 {
					c := color0
					if p&(0x80>>x) != 0 {
						c = color1
					}
					scr.SetColorIndex(baseX+x, baseY+y, c)
				}
			}
		}
	}
}

func (v *VDP) renderGraphics2Plane(scr *image.Paletted) {
	r := v.calcRenderArea(scr)
	offX, offY := r.Min.X, r.Min.Y
	addr := v.baseAddresses()
	for i := range 3 {
		for j := range 256 {
			m := v.VRAM[addr.nameTable+i*256+j]
			p := addr.patternGenerator + i*256*8 + int(m)*8
			pattern := v.VRAM[p : p+8]
			cp := addr.colorTable + i*256*8 + int(m)*8
			colors := v.VRAM[cp : cp+8]
			baseX := offX + (j%32)*8
			baseY := offY + (i*8+j/32)*8
			for y := range 8 {
				p := pattern[y]
				color0, color1 := uint8(colors[y]&0x0F), uint8(colors[y]&0xF0>>4)
				for x := range 8 {
					c := color0
					if p&(0x80>>x) != 0 {
						c = color1
					}
					scr.SetColorIndex(baseX+x, baseY+y, c)
				}
			}
		}
	}
}

func (v *VDP) renderMulticolorPlane(scr *image.Paletted) {
	r := v.calcRenderArea(scr)
	offX, offY := r.Min.X, r.Min.Y
	addr := v.baseAddresses()
	for cy := 0; cy < 4; cy += 6 {
		baseY := offY + cy*6*8
		for cx := 0; cx < 32; cx++ {
			baseX := offX + cx*8
			m := v.VRAM[addr.nameTable+cx+cy*32]
			p := addr.patternGenerator + int(m)*8
			pattern := v.VRAM[p : p+8]
			for row := range 8 {
				c := pattern[row]
				baseY2 := baseY + row*4
				for y := range 4 {
					colorA, colorB := c&0xF0>>4, c&0x0F
					for x := 0; x < 4; x++ {
						scr.SetColorIndex(baseX+x, baseY2+y, colorA)
					}
					for x := 4; x < 8; x++ {
						scr.SetColorIndex(baseX+x, baseY2+y, colorB)
					}
				}
			}
		}
	}
}

func (v *VDP) renderTextPlane(scr *image.Paletted) {
	r := v.calcRenderArea(scr)
	offX, offY := r.Min.X, r.Min.Y
	addr := v.baseAddresses()
	color0, color1 := v.textColorIndex(), v.backdropColorIndex()
	for cy := 0; cy < 24; cy++ {
		baseY := offY + cy*8
		for cx := 0; cx < 40; cx++ {
			baseX := offX + cx*6
			m := v.VRAM[addr.nameTable+cx+cy*40]
			p := addr.patternGenerator + int(m)*8
			pattern := v.VRAM[p : p+8]
			for y := range 8 {
				p := pattern[y]
				for x := range 6 {
					c := color0
					if p&(0x80>>x) != 0 {
						c = color1
					}
					scr.SetColorIndex(baseX+x, baseY+y, c)
				}
			}
		}
	}
}

type sprite struct {
	y       int
	x       int
	color   uint8
	pattern []uint8
}

func (v *VDP) spritesAll(params spriteParams, sprites []sprite) []sprite {
	sprites = sprites[:0]
	for i := range 32 {
		ap := params.attrAddr + i*4
		attr := v.VRAM[ap : ap+4]
		n := int(attr[2])
		pattern := params.patternAddr + n*8
		sp := sprite{
			y:       int(attr[0]),
			x:       int(attr[1]),
			color:   attr[3] & 0x0F,
			pattern: v.VRAM[pattern : pattern+params.psize],
		}
		if sp.y >= 224 {
			sp.y -= 256
		}
		if attr[3]&0x80 != 0 {
			sp.x -= 32
		}
		sprites = append(sprites, sp)
	}
	return sprites
}

func isIn(r image.Rectangle, x, y int) bool {
	return x >= r.Min.X && x < r.Max.X && y >= r.Min.Y && y < r.Max.Y
}

func (v *VDP) renderSpritesAll(scr *image.Paletted) {
	params := v.spriteParams()
	sprites := v.spritesAll(params, make([]sprite, 32))
	r := v.calcRenderArea(scr)
	for i := len(sprites) - 1; i >= 0; i-- {
		sp := sprites[i]
		baseX := r.Min.X + sp.x
		baseY := r.Min.Y + sp.y
		for j := range params.psize {
			p := sp.pattern[j]
			baseX2 := baseX + j/16*8*params.mag
			y := baseY + j%16*params.mag
			for dx := range 8 {
				if p&(0x80>>dx) != 0 {
					x := baseX2 + dx*params.mag
					if isIn(r, x, y) {
						scr.SetColorIndex(x, y, sp.color)
					}
					if params.mag == 2 {
						if isIn(r, x+1, y) {
							scr.SetColorIndex(x+1, y, sp.color)
						}
						if isIn(r, x, y+1) {
							scr.SetColorIndex(x, y+1, sp.color)
						}
						if isIn(r, x+1, y+1) {
							scr.SetColorIndex(x+1, y+1, sp.color)
						}
					}
				}
			}
		}
	}
}
