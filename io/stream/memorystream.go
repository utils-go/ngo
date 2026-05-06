package stream

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

// MemoryStream creates a stream whose backing store is memory.
// Equivalent to System.IO.MemoryStream in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.io.memorystream?view=netframework-4.7.2
type MemoryStream struct {
	buffer   *bytes.Buffer
	closed   bool
	mu       sync.RWMutex
}

// NewMemoryStream creates a new MemoryStream with an empty buffer.
func NewMemoryStream() *MemoryStream {
	return &MemoryStream{
		buffer: new(bytes.Buffer),
	}
}

// NewMemoryStreamWithCapacity creates a new MemoryStream with the specified initial capacity.
func NewMemoryStreamWithCapacity(capacity int) *MemoryStream {
	buf := new(bytes.Buffer)
	buf.Grow(capacity)
	return &MemoryStream{
		buffer: buf,
	}
}

// NewMemoryStreamFromBytes creates a new MemoryStream initialized with the specified byte array.
func NewMemoryStreamFromBytes(data []byte) *MemoryStream {
	return &MemoryStream{
		buffer: bytes.NewBuffer(data),
	}
}

// NewMemoryStreamFromBytesRange creates a new MemoryStream from a range of bytes.
func NewMemoryStreamFromBytesRange(data []byte, index, count int) *MemoryStream {
	if index < 0 || count < 0 || index+count > len(data) {
		return nil
	}
	return &MemoryStream{
		buffer: bytes.NewBuffer(data[index : index+count]),
	}
}

// Read reads a block of bytes from the current stream and writes the data to buffer.
// Equivalent to MemoryStream.Read in .NET.
func (ms *MemoryStream) Read(buffer []byte) (int, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	if ms.closed {
		return 0, fmt.Errorf("memory stream is closed")
	}
	return ms.buffer.Read(buffer)
}

// Write writes a block of bytes to the current stream.
// Equivalent to MemoryStream.Write in .NET.
func (ms *MemoryStream) Write(buffer []byte) (int, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if ms.closed {
		return 0, fmt.Errorf("memory stream is closed")
	}
	return ms.buffer.Write(buffer)
}

// Seek sets the position within the current stream.
// Equivalent to MemoryStream.Seek in .NET.
func (ms *MemoryStream) Seek(offset int64, whence int) (int64, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	if ms.closed {
		return 0, fmt.Errorf("memory stream is closed")
	}
	// bytes.Buffer doesn't support Seek, so we use bytes.Reader for that
	// But since Buffer is not a Seeker, we implement a basic version
	reader := bytes.NewReader(ms.buffer.Bytes())
	return reader.Seek(offset, whence)
}

// SeekOrigin sets the position using the .NET SeekOrigin enum.
func (ms *MemoryStream) SeekOrigin(offset int64, origin SeekOrigin) (int64, error) {
	return ms.Seek(offset, seekOriginToWhence(origin))
}

// Length gets the length of the stream in bytes.
func (ms *MemoryStream) Length() int64 {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return int64(ms.buffer.Len())
}

// SetLength sets the length of the current stream.
func (ms *MemoryStream) SetLength(value int64) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if ms.closed {
		return fmt.Errorf("memory stream is closed")
	}
	current := ms.buffer.Bytes()
	if int64(len(current)) > value {
		ms.buffer = bytes.NewBuffer(current[:value])
	} else if int64(len(current)) < value {
		padding := make([]byte, value-int64(len(current)))
		ms.buffer.Write(padding)
	}
	return nil
}

// Position gets the current position within the stream.
func (ms *MemoryStream) Position() int64 {
	pos, _ := ms.Seek(0, io.SeekCurrent)
	return pos
}

// SetPosition sets the position within the stream.
func (ms *MemoryStream) SetPosition(value int64) (int64, error) {
	return ms.Seek(value, io.SeekStart)
}

// Flush clears any buffers — no-op for MemoryStream.
func (ms *MemoryStream) Flush() error {
	return nil
}

// Close closes the current stream.
func (ms *MemoryStream) Close() error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.closed = true
	return nil
}

// CanRead indicates whether the stream supports reading.
func (ms *MemoryStream) CanRead() bool {
	return true
}

// CanWrite indicates whether the stream supports writing.
func (ms *MemoryStream) CanWrite() bool {
	return !ms.closed
}

// CanSeek indicates whether the stream supports seeking.
func (ms *MemoryStream) CanSeek() bool {
	return true
}

// ToArray writes the stream contents to a byte array, regardless of the Position property.
func (ms *MemoryStream) ToArray() []byte {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return append([]byte(nil), ms.buffer.Bytes()...)
}

// GetBuffer returns the array of unsigned bytes from which this stream was created.
func (ms *MemoryStream) GetBuffer() []byte {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.buffer.Bytes()
}

// Capacity gets or sets the number of bytes allocated for this stream.
func (ms *MemoryStream) Capacity() int {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.buffer.Cap()
}

// SetCapacity sets the capacity.
func (ms *MemoryStream) SetCapacity(value int) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if value > ms.buffer.Cap() {
		ms.buffer.Grow(value - ms.buffer.Cap())
	}
}

// ReadByte reads a byte from the stream and advances the position by one byte.
func (ms *MemoryStream) ReadByte() (byte, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	if ms.closed {
		return 0, fmt.Errorf("memory stream is closed")
	}
	return ms.buffer.ReadByte()
}

// WriteByte writes a byte to the current stream.
func (ms *MemoryStream) WriteByte(b byte) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if ms.closed {
		return fmt.Errorf("memory stream is closed")
	}
	return ms.buffer.WriteByte(b)
}

// WriteTo writes the entire contents of this memory stream to another stream.
func (ms *MemoryStream) WriteTo(dst io.Writer) (int64, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	if ms.closed {
		return 0, fmt.Errorf("memory stream is closed")
	}
	reader := bytes.NewReader(ms.buffer.Bytes())
	return io.Copy(dst, reader)
}

// ReadFrom reads all data from another stream and writes it to this memory stream.
func (ms *MemoryStream) ReadFrom(src io.Reader) (int64, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if ms.closed {
		return 0, fmt.Errorf("memory stream is closed")
	}
	return io.Copy(ms.buffer, src)
}
