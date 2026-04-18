package engine

/*
#cgo LDFLAGS: -L${SRCDIR}/../../../lib -lengine
#cgo CPPFLAGS: -I${SRCDIR}/../../../engine
#include "bridge.h"
*/
import "C"
import (
	"fmt"
)

// ProcessBatch calls our C++ engine to process an inference batch.
func ProcessBatch(size int) error {
	res := C.ProcessBatch(C.int(size))
	if res != 0 {
		return fmt.Errorf("engine failed with error code: %d", res)
	}
	return nil
}
