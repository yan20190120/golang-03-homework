package main

import (
	"log"
	"unsafe"
)

func main() {
	// init1
	var a1 = [3]int{1, 2, 3}
	log.Printf("a1: %v, len: %d, size: %v\n", a1, len(a1), unsafe.Sizeof(a1))
	// init2
	var a2 = [3]int{1, 2}
	log.Printf("a2: %v, len: %d, size: %v\n", a2, len(a2), unsafe.Sizeof(a2))
	// init3
	var a3 = [...]int{1, 2, 3}
	log.Printf("a3: %v, len: %d, size: %v\n", a3, len(a3), unsafe.Sizeof(a3))
	// init4
	var a4 = [...]int{0: 1, 2: 3}
	log.Printf("a4: %v, len: %d, size: %v\n", a4, len(a4), unsafe.Sizeof(a4))
}
