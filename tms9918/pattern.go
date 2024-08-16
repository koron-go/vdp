package tms9918

var ReferencePattern = []uint8{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // (SPACE)
	0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x20, 0x00, // !
	0x50, 0x50, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, // "
	0x50, 0x50, 0xF8, 0x50, 0xF8, 0x50, 0x50, 0x00, // #
	0x20, 0x78, 0xA0, 0x70, 0x28, 0xF0, 0x20, 0x00, // $
	0xC0, 0xC8, 0x10, 0x20, 0x40, 0x98, 0x18, 0x00, // %
	0x40, 0xA0, 0xA0, 0x40, 0xA8, 0x90, 0x68, 0x00, // &
	0x20, 0x20, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, // '
	0x20, 0x40, 0x80, 0x80, 0x80, 0x40, 0x20, 0x00, // (
	0x20, 0x10, 0x08, 0x08, 0x08, 0x10, 0x20, 0x00, // )
	0x20, 0xA8, 0x70, 0x20, 0x70, 0xA8, 0x20, 0x00, // *
	0x00, 0x20, 0x20, 0xF8, 0x20, 0x20, 0x00, 0x00, // +
	0x00, 0x00, 0x00, 0x00, 0x20, 0x20, 0x40, 0x00, // ,
	0x00, 0x00, 0x00, 0xF8, 0x00, 0x00, 0x00, 0x00, // -
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00, // .
	0x00, 0x08, 0x10, 0x20, 0x40, 0x80, 0x00, 0x00, // /
	0x70, 0x88, 0x98, 0xA8, 0xC8, 0x88, 0x70, 0x00, // 0
	0x20, 0x60, 0x20, 0x20, 0x20, 0x20, 0x70, 0x00, // 1
	0x70, 0x88, 0x08, 0x30, 0x40, 0x80, 0xF8, 0x00, // 2
	0xF8, 0x08, 0x10, 0x30, 0x08, 0x88, 0x70, 0x00, // 3
	0x10, 0x30, 0x50, 0x90, 0xF8, 0x10, 0x10, 0x00, // 4
	0xF8, 0x80, 0xF0, 0x08, 0x08, 0x88, 0x70, 0x00, // 5
	0x38, 0x40, 0x80, 0xF0, 0x88, 0x88, 0x70, 0x00, // 6
	0xF8, 0x08, 0x10, 0x20, 0x40, 0x40, 0x40, 0x00, // 7
	0x70, 0x88, 0x88, 0x70, 0x88, 0x88, 0x70, 0x00, // 8
	0x70, 0x88, 0x88, 0x78, 0x08, 0x10, 0xE0, 0x00, // 9
	0x00, 0x00, 0x20, 0x00, 0x20, 0x00, 0x00, 0x00, // :
	0x00, 0x00, 0x20, 0x00, 0x20, 0x20, 0x40, 0x00, // ;
	0x10, 0x20, 0x40, 0x80, 0x40, 0x20, 0x10, 0x00, // <
	0x00, 0x00, 0xF8, 0x00, 0xF8, 0x00, 0x00, 0x00, // =
	0x40, 0x20, 0x10, 0x08, 0x10, 0x20, 0x40, 0x00, // >
	0x70, 0x88, 0x10, 0x20, 0x20, 0x00, 0x20, 0x00, // ?

	0x70, 0x88, 0xA8, 0xB8, 0xB0, 0x80, 0x78, 0x00, // @
	0x20, 0x50, 0x88, 0x88, 0xF8, 0x88, 0x88, 0x00, // A
	0xF0, 0x88, 0x88, 0xF0, 0x88, 0x88, 0xF0, 0x00, // B
	0x70, 0x88, 0x80, 0x80, 0x80, 0x88, 0x70, 0x00, // C
	0xF0, 0x88, 0x88, 0x88, 0x88, 0x88, 0xF0, 0x00, // D
	0xF8, 0x80, 0x80, 0xF0, 0x80, 0x80, 0xF0, 0x00, // E
	0xF8, 0x80, 0x80, 0xF0, 0x80, 0x80, 0x80, 0x00, // F
	0x78, 0x80, 0x80, 0x80, 0x98, 0x88, 0x78, 0x00, // G
	0x88, 0x88, 0x88, 0xF8, 0x88, 0x88, 0x88, 0x00, // H
	0x70, 0x20, 0x20, 0x20, 0x20, 0x20, 0x70, 0x00, // I
	0x08, 0x08, 0x08, 0x08, 0x08, 0x88, 0x70, 0x00, // J
	0x88, 0x90, 0xA0, 0xC0, 0xA0, 0x90, 0x88, 0x00, // K
	0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0xF8, 0x00, // L
	0x88, 0xD8, 0xA8, 0xA8, 0x88, 0x88, 0x88, 0x00, // M
	0x88, 0x88, 0xC8, 0xA8, 0x98, 0x88, 0x88, 0x00, // N
	0x70, 0x88, 0x88, 0x88, 0x88, 0x88, 0x70, 0x00, // O
	0xF0, 0x88, 0x88, 0xF0, 0x80, 0x80, 0x80, 0x00, // P
	0x70, 0x88, 0x88, 0x88, 0xA8, 0x90, 0x68, 0x00, // Q
	0xF0, 0x88, 0x88, 0xF0, 0xA0, 0x90, 0x88, 0x00, // R
	0x70, 0x88, 0x80, 0x70, 0x08, 0x88, 0x70, 0x00, // S
	0xF8, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, // T
	0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x70, 0x00, // U
	0x88, 0x88, 0x88, 0x88, 0x88, 0x50, 0x20, 0x00, // V
	0x88, 0x88, 0x88, 0xA8, 0xA8, 0xD8, 0x88, 0x00, // W
	0x88, 0x88, 0x50, 0x20, 0x50, 0x88, 0x88, 0x00, // X
	0x88, 0x88, 0x50, 0x20, 0x20, 0x20, 0x20, 0x00, // Y
	0xF8, 0x08, 0x10, 0x20, 0x40, 0x80, 0xF8, 0x00, // Z
	0xF8, 0xC0, 0xC0, 0xC0, 0xC0, 0xC0, 0xF8, 0x00, // [
	0x00, 0x80, 0x40, 0x20, 0x10, 0x08, 0x00, 0x00, // \
	0xF8, 0x18, 0x18, 0x18, 0x18, 0x18, 0xF8, 0x00, // ]
	0x00, 0x00, 0x20, 0x50, 0x88, 0x00, 0x00, 0x00, // ^
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF8, 0x00, // _

	0x40, 0x20, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, // `
	0x00, 0x00, 0x70, 0x88, 0xF8, 0x88, 0x88, 0x00, // a
	0x00, 0x00, 0xF0, 0x48, 0x70, 0x48, 0xF0, 0x00, // b
	0x00, 0x00, 0x78, 0x80, 0x80, 0x80, 0x78, 0x00, // c
	0x00, 0x00, 0xF0, 0x48, 0x48, 0x48, 0xF0, 0x00, // d
	0x00, 0x00, 0xF0, 0x80, 0xE0, 0x80, 0xF0, 0x00, // e
	0x00, 0x00, 0xF0, 0x80, 0xE0, 0x80, 0x80, 0x00, // f
	0x00, 0x00, 0x78, 0x80, 0xB8, 0x88, 0x70, 0x00, // g
	0x00, 0x00, 0x88, 0x88, 0xF8, 0x88, 0x88, 0x00, // h
	0x00, 0x00, 0xF8, 0x20, 0x20, 0x20, 0xF8, 0x00, // i
	0x00, 0x00, 0x70, 0x20, 0x20, 0xA0, 0xE0, 0x00, // j
	0x00, 0x00, 0x90, 0xA0, 0xC0, 0xA0, 0x90, 0x00, // k
	0x00, 0x00, 0x80, 0x80, 0x80, 0x80, 0xF8, 0x00, // l
	0x00, 0x00, 0x88, 0xD8, 0xA8, 0x88, 0x88, 0x00, // m
	0x00, 0x00, 0x88, 0xC8, 0xA8, 0x98, 0x88, 0x00, // n
	0x00, 0x00, 0xF8, 0x88, 0x88, 0x88, 0xF8, 0x00, // o
	0x00, 0x00, 0xF0, 0x88, 0xF0, 0x80, 0x80, 0x00, // p
	0x00, 0x00, 0xF8, 0x88, 0xA8, 0x90, 0xE0, 0x00, // q
	0x00, 0x00, 0xF8, 0x88, 0xF8, 0xA0, 0x90, 0x00, // r
	0x00, 0x00, 0x78, 0x80, 0x70, 0x08, 0xF0, 0x00, // s
	0x00, 0x00, 0xF8, 0x20, 0x20, 0x20, 0x20, 0x00, // t
	0x00, 0x00, 0x88, 0x88, 0x88, 0x88, 0x70, 0x00, // u
	0x00, 0x00, 0x88, 0x88, 0x90, 0xA0, 0x40, 0x00, // v
	0x00, 0x00, 0x88, 0x88, 0xA8, 0xD8, 0x88, 0x00, // w
	0x00, 0x00, 0x88, 0x50, 0x20, 0x50, 0x88, 0x00, // x
	0x00, 0x00, 0x88, 0x50, 0x20, 0x20, 0x20, 0x00, // y
	0x00, 0x00, 0xF8, 0x10, 0x20, 0x40, 0xF8, 0x00, // z
	0x38, 0x40, 0x20, 0xC0, 0x20, 0x40, 0x38, 0x00, // {
	0x40, 0x20, 0x10, 0x08, 0x10, 0x20, 0x40, 0x00, // |
	0xE0, 0x10, 0x20, 0x18, 0x20, 0x10, 0xE0, 0x00, // }
	0x40, 0xA8, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, // ~
	0xA8, 0x50, 0xA8, 0x50, 0xA8, 0x50, 0xA8, 0x00, // (0x7F)
}

func SetupReferencePattern(vdp *VDP) {
	start := vdp.patternGeneratorBaseAddress() + 8*0x20
	copy(vdp.VRAM[start:start+len(ReferencePattern)], ReferencePattern)
}
