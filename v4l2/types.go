package v4l2

import (
	"context"
)

// Device is the base interface for a v4l2 device
type Device interface {
	Name() string
	Fd() uintptr
	Capability() Capability
	MemIOType() IOType
	GetOutput() <-chan []byte
	SetInput(<-chan []byte)
	Close() error
}

// StreamingDevice represents device that supports streaming IO
// via mapped buffer sharing.
type StreamingDevice interface {
	Device
	Buffers() [][]byte
	BufferType() BufType
	BufferCount() uint32
	Start(context.Context) error
	Stop() error
}

//// CaptureDevice represents a device that captures video from an underlying device
//type CaptureDevice interface {
//	StreamingDevice
//	StartCapture(context.Context) (<-chan []byte, error)
//}
//
//// OutputDevice represents a device that can output video data to an underlying device driver
//type OutputDevice interface {
//	StreamingDevice
//	StartOutput(context.Context, chan<- []byte) error
//}