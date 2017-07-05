package go2c

/*
#cgo CFLAGS : -I. -m32
#cgo LDFLAGS:

#include "go2c.h"
*/
import "C"
import "unsafe"

// ConvertStringArrayToChars convert go string array into C char**;
// Directly return is int, unsafe.Pointer, should be cast to C type;
// Example:
// 		c, v := go2c.ConvertStringArrayToChars(os.Args)
// 		argc := C.int(c)
// 		argv := (**C.char)(v)
func ConvertStringArrayToChars(args []string) (argc int, argv unsafe.Pointer) {
	argc = len(args)
	var s string
	for _, v := range args {
		s += v
		s += "\n"
	}
	argv = C.pointerToArray(C.int(argc), C.CString(s))
	return
}

// GoConvertStringArrayToChars convert go string array into C char**;
// See "<ConvertStringArrayToChars>" also;
func GoConvertStringArrayToChars(args []string) (argc int, argv unsafe.Pointer) {
	argc = len(args)
	var s string
	for _, v := range args {
		s += v
		s += "\x00"
	}
	argv = unsafe.Pointer(&s)
	return
}
