package stream

import (
	"encoding/binary"
	"fmt"
	"io"
)

// BinaryWriter writes primitive types in binary to a stream and supports
// writing strings in a specific encoding.
// Equivalent to System.IO.BinaryWriter in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.io.binarywriter?view=netframework-4.7.2
type BinaryWriter struct {
	writer       io.Writer
	baseStream   io.Writer
	closed       bool
	littleEndian bool
}

// NewBinaryWriter creates a new BinaryWriter for the specified stream.
func NewBinaryWriter(stream io.Writer) *BinaryWriter {
	return &BinaryWriter{
		writer:      stream,
		baseStream:  stream,
		littleEndian: true,
	}
}

// NewBinaryWriterWithEndian creates a BinaryWriter with specified byte order.
func NewBinaryWriterWithEndian(stream io.Writer, littleEndian bool) *BinaryWriter {
	return &BinaryWriter{
		writer:      stream,
		baseStream:  stream,
		littleEndian: littleEndian,
	}
}

// BaseStream returns the underlying stream.
func (bw *BinaryWriter) BaseStream() io.Writer {
	return bw.baseStream
}

// Close closes the current BinaryWriter and the underlying stream.
func (bw *BinaryWriter) Close() error {
	if bw.closed {
		return nil
	}
	bw.closed = true
	if closer, ok := bw.baseStream.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// Flush clears all buffers for the current writer.
func (bw *BinaryWriter) Flush() error {
	if bw.closed {
		return fmt.Errorf("binary writer is closed")
	}
	if flusher, ok := bw.writer.(interface{ Flush() error }); ok {
		return flusher.Flush()
	}
	return nil
}

// Seek sets the position within the current stream.
func (bw *BinaryWriter) Seek(offset int64, origin SeekOrigin) (int64, error) {
	if bw.closed {
		return 0, fmt.Errorf("binary writer is closed")
	}
	if seeker, ok := bw.baseStream.(io.Seeker); ok {
		return seeker.Seek(offset, seekOriginToWhence(origin))
	}
	return 0, fmt.Errorf("stream does not support seeking")
}

// Write writes raw bytes to the stream.
func (bw *BinaryWriter) Write(data []byte) error {
	if bw.closed {
		return fmt.Errorf("binary writer is closed")
	}
	_, err := bw.writer.Write(data)
	return err
}

// WriteByte writes an unsigned byte to the current stream.
func (bw *BinaryWriter) WriteByte(b byte) error {
	if bw.closed {
		return fmt.Errorf("binary writer is closed")
	}
	_, err := bw.writer.Write([]byte{b})
	return err
}

// WriteBytes writes a byte array to the underlying stream.
func (bw *BinaryWriter) WriteBytes(data []byte) error {
	return bw.Write(data)
}

// WriteBoolean writes a one-byte Boolean value to the current stream.
func (bw *BinaryWriter) WriteBoolean(value bool) error {
	if value {
		return bw.WriteByte(1)
	}
	return bw.WriteByte(0)
}

// WriteInt16 writes a 2-byte signed integer to the current stream.
func (bw *BinaryWriter) WriteInt16(val int16) error {
	return bw.writeBinary(val)
}

// WriteInt32 writes a 4-byte signed integer to the current stream.
func (bw *BinaryWriter) WriteInt32(val int32) error {
	return bw.writeBinary(val)
}

// WriteInt64 writes an 8-byte signed integer to the current stream.
func (bw *BinaryWriter) WriteInt64(val int64) error {
	return bw.writeBinary(val)
}

// WriteUInt16 writes a 2-byte unsigned integer to the current stream.
func (bw *BinaryWriter) WriteUInt16(val uint16) error {
	return bw.writeBinary(val)
}

// WriteUInt32 writes a 4-byte unsigned integer to the current stream.
func (bw *BinaryWriter) WriteUInt32(val uint32) error {
	return bw.writeBinary(val)
}

// WriteUInt64 writes an 8-byte unsigned integer to the current stream.
func (bw *BinaryWriter) WriteUInt64(val uint64) error {
	return bw.writeBinary(val)
}

// WriteSingle writes a 4-byte floating point value to the current stream.
func (bw *BinaryWriter) WriteSingle(val float32) error {
	return bw.writeBinary(val)
}

// WriteDouble writes an 8-byte floating point value to the current stream.
func (bw *BinaryWriter) WriteDouble(val float64) error {
	return bw.writeBinary(val)
}

// WriteString writes a length-prefixed string to this stream.
func (bw *BinaryWriter) WriteString(s string) error {
	if bw.closed {
		return fmt.Errorf("binary writer is closed")
	}
	// Write length as int32 then the string bytes
	if err := bw.WriteInt32(int32(len(s))); err != nil {
		return err
	}
	return bw.Write([]byte(s))
}

// WriteRune writes a Unicode character to the current stream.
func (bw *BinaryWriter) WriteRune(r rune) error {
	return bw.writeBinary(int32(r))
}

// writeBinary writes a value using the configured byte order.
func (bw *BinaryWriter) writeBinary(val interface{}) error {
	if bw.closed {
		return fmt.Errorf("binary writer is closed")
	}
	var order binary.ByteOrder
	if bw.littleEndian {
		order = binary.LittleEndian
	} else {
		order = binary.BigEndian
	}
	return binary.Write(bw.writer, order, val)
}
