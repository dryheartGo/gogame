package strings

import (
	"reflect"
	"unsafe"
)

func IsBlank(str string) bool {
	return len(str) <= 0
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func String2Byte(s string) (b []byte) {
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pString := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pBytes.Data = pString.Data
	pBytes.Len = pString.Len
	pBytes.Cap = pString.Len
	return
}
