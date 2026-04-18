package engine

/*
#cgo LDFLAGS: -L${SRCDIR}/../../../lib -lengine
#cgo CPPFLAGS: -I${SRCDIR}/../../../engine
#include "bridge.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// ProcessBatch calls our C++ engine to process an inference batch.
// We use unsafe.SliceData to pass the Go slice pointer directly to C++
// with zero-copy overhead.
func ProcessBatch(promptLengths []int) error {
	if len(promptLengths) == 0 {
		return nil
	}

	// Extract raw pointer to the start of the Go slice via unsafe.Pointer bridge
	ptr := (*C.int)(unsafe.Pointer(unsafe.SliceData(promptLengths)))
	size := C.int(len(promptLengths))

	res := C.ProcessBatch(ptr, size)
	if res != 0 {
		return fmt.Errorf("engine failed with error code: %d", res)
	}
	return nil
}
