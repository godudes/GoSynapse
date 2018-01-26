package rak

import "testing"

func TestMagic(t *testing.T) {
	dst := make([]byte, 128)
	off := 2
	writeMagic(dst, &off)
	println(validateMagic(dst, &off))
}