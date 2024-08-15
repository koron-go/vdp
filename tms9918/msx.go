package tsm9918

func MSXScreen0(vdp *VDP) {
	vdp.Register2.ResetSet(0x0F, 0x00) // 0x0000: name table
	vdp.Register4.ResetSet(0x07, 0x01) // 0x0800: pattern generator table
}

func MSXScreen1(vdp *VDP) {
	vdp.Register4.ResetSet(0x07, 0x00) // 0x0000: pattern generator table
	vdp.Register2.ResetSet(0x0F, 0x06) // 0x1800: name table
	vdp.Register3.ResetSet(0xFF, 0x80) // 0x2000: color table
	vdp.Register5.ResetSet(0x7F, 0x36) // 0x1B00: sprite attribute table
	vdp.Register6.ResetSet(0x07, 0x07) // 0x3800: sprite generator table
}

func MSXScreen2(vdp *VDP) {
	vdp.Register4.ResetSet(0x07, 0x00) // 0x0000: pattern generator table
	vdp.Register2.ResetSet(0x0F, 0x06) // 0x1800: name table
	vdp.Register3.ResetSet(0xFF, 0x80) // 0x2000: color table
	vdp.Register5.ResetSet(0x7F, 0x36) // 0x1B00: sprite attribute table
	vdp.Register6.ResetSet(0x07, 0x07) // 0x3800: sprite generator table
}

func MSXScreen3(vdp *VDP) {
	vdp.Register4.ResetSet(0x07, 0x00) // 0x0000: pattern generator table
	vdp.Register2.ResetSet(0x0F, 0x02) // 0x0800: name table
	vdp.Register5.ResetSet(0x7F, 0x36) // 0x1B00: sprite attribute table
	vdp.Register6.ResetSet(0x07, 0x07) // 0x3800: sprite generator table
}
