package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int64
	fmt.Println("Sizeof int64:", unsafe.Sizeof(x))
	// unsafe.Pointer dipakai hanya untuk interoperabilitas rendah level (FFI, optimasi).
	// Salah pakai = undefined behavior / crash.
}
