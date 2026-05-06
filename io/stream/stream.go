// Package stream provides .NET System.IO.Stream-like abstractions for Go.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.io.stream?view=netframework-4.7.2
package stream

import "io"

// Stream represents a sequence of bytes, equivalent to System.IO.Stream in .NET.
// Composes the standard Go io interfaces with additional stream metadata methods.
type Stream interface {
	io.Reader
	io.Writer
	io.Seeker
	io.Closer

	// Length gets the length in bytes of the stream.
	Length() int64

	// SetLength sets the length of the current stream.
	SetLength(value int64) error

	// Position gets the current position within the stream.
	// Equivalent to Seek(0, io.SeekCurrent).
	Position() int64

	// SetPosition sets the position within the stream.
	// Equivalent to Seek(offset, io.SeekStart).
	SetPosition(value int64) (int64, error)

	// Flush clears all buffers for this stream and causes any buffered data
	// to be written to the underlying device.
	Flush() error

	// CanRead gets a value indicating whether the current stream supports reading.
	CanRead() bool

	// CanWrite gets a value indicating whether the current stream supports writing.
	CanWrite() bool

	// CanSeek gets a value indicating whether the current stream supports seeking.
	CanSeek() bool
}

// SeekOrigin specifies the reference point used to obtain the new position in a stream.
// Equivalent to System.IO.SeekOrigin in .NET.
type SeekOrigin int

const (
	// SeekBegin specifies the beginning of the stream.
	SeekBegin SeekOrigin = iota
	// SeekCurrent specifies the current position within the stream.
	SeekCurrent
	// SeekEnd specifies the end of the stream.
	SeekEnd
)

// seekOriginToWhence converts SeekOrigin to Go's io.SeekMode value.
func seekOriginToWhence(origin SeekOrigin) int {
	switch origin {
	case SeekBegin:
		return io.SeekStart
	case SeekCurrent:
		return io.SeekCurrent
	case SeekEnd:
		return io.SeekEnd
	default:
		return io.SeekStart
	}
}
