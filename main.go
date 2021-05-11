package main

/*
#include <stdlib.h>
#include <pikchr.h>
*/
import "C"
import "unsafe"

import (
	"fmt"
	"io/ioutil"
)

func main() {
	s, err := ioutil.ReadFile("prova.txt")
	if err != nil {
		fmt.Printf("-> Error! %s \n", err)
	}
	source := C.CString(string(s))
	defer C.free(unsafe.Pointer(source))

	cssClass := C.CString("prova")
	defer C.free(unsafe.Pointer(cssClass))

	w, h := C.int(640), C.int(480)
	zOut := C.pikchr(source, cssClass, 0, &w, &h)

	svg := C.GoString(zOut)
	fmt.Println(svg)

	err = ioutil.WriteFile("prova.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Println("Non posso scrivere svg su file")
	}
}
