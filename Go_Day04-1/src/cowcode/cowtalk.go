package cowcode

// #include <cowtalk.h>
//#include <stdio.h>
//#include <stdlib.h>
//#include <string.h>
import "C"
import "unsafe"

func CowTalk(str string) string {
	cstr := C.CString(str)
	p := C.ask_cow(cstr)
	defer C.free(unsafe.Pointer(p))
	s := C.GoString(p)

	return s
}
