package main

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -lm
#include "vector_math.h"
*/
import "C"
import (
	"unsafe"
)

func EuclideanDistance(a []float32, b []float32) float32 {
	if len(a) != len(b) || len(a) == 0 {
		return -1.0
	}
	n := C.int(len(a))
	ptrA := (*C.float)(unsafe.Pointer(&a[0]))
	ptrB := (*C.float)(unsafe.Pointer(&b[0]))
	result := C.euclidean_distance(ptrA, ptrB, n)
	return float32(result)
}
