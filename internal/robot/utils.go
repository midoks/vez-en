package robot

import (
	"math/rand"
	"unsafe"
)

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = LETTER_BYTES[rand.Intn(len(LETTER_BYTES))]
	}
	return BytesToString(b)
}
