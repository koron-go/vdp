package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/koron-go/vdp/tms9918"
)

func main() {
	vdp := tms9918.New(make([]uint8, 16384))
	tms9918.MSXScreen1(vdp)
	tms9918.SetupReferencePattern(vdp)
	// setup color
	vdp.Register7 = 0x07
	for i := range 32 {
		vdp.VRAM[0x2000+i] = 0xF4
	}
	// setup name table
	for i := 0x20; i <= 0x7F; i++ {
		vdp.VRAM[0x1800+i] = uint8(i)
	}

	// render
	img, err := vdp.NewPaletted(image.Rect(0, 0, 256+32, 192+32))
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	vdp.Render(img)
	dur := time.Since(start)
	log.Printf("rendered in %d", dur)
	log.Printf("start=%s", start)

	// output as PNG
	f, err := os.Create("test.png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(f, img)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
