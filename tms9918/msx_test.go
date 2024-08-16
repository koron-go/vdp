package tms9918

import (
	"testing"
)

func TestMSXScreen0(t *testing.T) {
	vdp := New(make([]uint8, 16384))
	MSXScreen0(vdp)
	ba := vdp.baseAddresses()
	if want, got := 0x0000, ba.nameTable; got != want {
		t.Errorf("unexpected name table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x0800, ba.patternGenerator; got != want {
		t.Errorf("unexpected pattern generator table: want=%04X got=%04X", want, got)
	}
}

func TestMSXScreen1(t *testing.T) {
	vdp := New(make([]uint8, 16384))
	MSXScreen1(vdp)
	ba := vdp.baseAddresses()
	sp := vdp.spriteParams()
	if want, got := 0x0000, ba.patternGenerator; got != want {
		t.Errorf("unexpected pattern generator table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x1800, ba.nameTable; got != want {
		t.Errorf("unexpected name table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x2000, ba.colorTable; got != want {
		t.Errorf("unexpected color table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x1B00, sp.attrAddr; got != want {
		t.Errorf("unexpected sprite attribute table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x3800, sp.patternAddr; got != want {
		t.Errorf("unexpected sprite generator table: want=%04X got=%04X", want, got)
	}
}

func TestMSXScreen2(t *testing.T) {
	vdp := New(make([]uint8, 16384))
	MSXScreen2(vdp)
	ba := vdp.baseAddresses()
	sp := vdp.spriteParams()
	if want, got := 0x0000, ba.patternGenerator; got != want {
		t.Errorf("unexpected pattern generator table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x1800, ba.nameTable; got != want {
		t.Errorf("unexpected name table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x2000, ba.colorTable; got != want {
		t.Errorf("unexpected color table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x1B00, sp.attrAddr; got != want {
		t.Errorf("unexpected sprite attribute table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x3800, sp.patternAddr; got != want {
		t.Errorf("unexpected sprite generator table: want=%04X got=%04X", want, got)
	}
}

func TestMSXScreen3(t *testing.T) {
	vdp := New(make([]uint8, 16384))
	MSXScreen3(vdp)
	ba := vdp.baseAddresses()
	sp := vdp.spriteParams()
	if want, got := 0x0000, ba.patternGenerator; got != want {
		t.Errorf("unexpected pattern generator table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x0800, ba.nameTable; got != want {
		t.Errorf("unexpected name table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x1B00, sp.attrAddr; got != want {
		t.Errorf("unexpected sprite attribute table: want=%04X got=%04X", want, got)
	}
	if want, got := 0x3800, sp.patternAddr; got != want {
		t.Errorf("unexpected sprite generator table: want=%04X got=%04X", want, got)
	}
}
