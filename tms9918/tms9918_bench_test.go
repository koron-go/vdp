package tms9918

import (
	"image"
	"testing"
)

func BenchmarkRenderScreen1(b *testing.B) {
	vdp := New(make([]uint8, 16384))
	MSXScreen1(vdp)
	SetupReferencePattern(vdp)
	// setup color
	vdp.Register7 = 0x07
	for i := range 32 {
		vdp.VRAM[0x2000+i] = 0xF4
	}
	// setup name table
	for i := 0x20; i <= 0x7F; i++ {
		vdp.VRAM[0x1800+i] = uint8(i)
	}
	// prepare render target image buffer
	img, err := vdp.NewPaletted(image.Rect(0, 0, 256+16, 192+16))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		vdp.Render(img)
	}
}
