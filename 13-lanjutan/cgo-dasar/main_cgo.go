//go:build cgo

package main

/*
#include <stdio.h>
static void halo_dari_c(void) { puts("Halo dari C (cgo)"); }
*/
import "C"

func main() {
	C.halo_dari_c()
}
