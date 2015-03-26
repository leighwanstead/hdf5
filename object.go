package hdf5

// #include <stdlib.h>
// #include <hdf5.h>
import "C"

import (
	"unsafe"
)

type object struct {
	data unsafe.Pointer
	sid  C.hid_t
	tid  C.hid_t

	inner []*object
}

func newObject() *object {
	return &object{
		sid: -1,
		tid: -1,
	}
}

func (o *object) free() {
	for i := range o.inner {
		o.inner[i].free()
	}
	if o.tid >= 0 {
		_ = C.H5Tclose(o.tid)
	}
	if o.sid >= 0 {
		_ = C.H5Sclose(o.sid)
	}
}
