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
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <txt-file-name>")
		os.Exit(1)
	}

	s, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("-> Error! %s \n", err)
	}
	source := C.CString(string(s))
	defer C.free(unsafe.Pointer(source))

	cssClass := C.CString("example")
	defer C.free(unsafe.Pointer(cssClass))

	w, h := C.int(640), C.int(480)
	zOut := C.pikchr(source, cssClass, 0, &w, &h)

	svg := C.GoString(zOut)
	fmt.Println(svg)

	err = ioutil.WriteFile("example.svg", []byte(svg), 0644)
	if err != nil {
		fmt.Println("Unable to write result to SVG file")
	}
}
