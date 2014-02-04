package window

// #cgo LDFLAGS: -lcsfml-graphics
// #include <SFML/Graphics/Color.h>
import "C"

import (
	"image/color"
)

// sfmlColor returns a C sfColor based on the provided Go color.Color.
func sfmlColor(c color.Color) C.sfColor {
	r, g, b, a := c.RGBA()
	sfmlCol := C.sfColor{
		r: C.sfUint8(r),
		g: C.sfUint8(g),
		b: C.sfUint8(b),
		a: C.sfUint8(a),
	}
	return sfmlCol
}

// utf32 returns the UTF-32 representation of s.
func utf32(s string) *C.sfUint32 {
	s32 := make([]C.sfUint32, 0, len(s))
	for _, r := range s {
		s32 = append(s32, C.sfUint32(r))
	}
	return &s32[0]
}
