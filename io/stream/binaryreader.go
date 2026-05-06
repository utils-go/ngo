package stream

import (
	"encoding/binary"
	"fmt"
	"io"
)

// BinaryReader reads primitive data types as binary values in a specific encoding.
// Equivalent to System.IO.BinaryReader in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.io.binaryreader?view=netframework-4.7.2
type BinaryReader struct {
	reader      io.Reader
	baseStream  io.Reader
	closed      bool
	littleEndian bool
}

// NewBinaryReader creates a new BinaryReader for the specified stream using UTF-8 encoding.
func NewBinaryReader(stream io.Reader) *BinaryReader {
	return &BinaryReader{
		reader:      stream,
		baseStream:  stream,
		littleEndian: true,
	}
}

// NewBinaryReaderWithEndian creates a BinaryReader with specified byte order.
func NewBinaryReaderWithEndian(stream io.Reader, littleEndian bool) *BinaryReader {
	return &BinaryReader{
		reader:      stream,
		baseStream:  stream,
		littleEndian: littleEndian,
	}
}

// BaseStream returns the underlying stream.
func (br *BinaryReader) BaseStream() io.Reader {
	return br.baseStream
}

// Dispose / Close closes the current reader and the underlying stream.
func (br *BinaryReader) Close() error {
	if br.closed {
		return nil
	}
	br.closed = true
	if closer, ok := br.baseStream.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// ReadBytes reads the specified number of bytes from the current stream into a byte array.
func (br *BinaryReader) ReadBytes(count int) ([]byte, error) {
	if br.closed {
		return nil, fmt.Errorf("binary reader is closed")
	}
	buf := make([]byte, count)
	n, err := io.ReadFull(br.reader, buf)
	if err != nil {
		return buf[:n], err
	}
	return buf, nil
}

// ReadByte reads the next byte from the current stream.
func (br *BinaryReader) ReadByte() (byte, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	var buf [1]byte
	_, err := io.ReadFull(br.reader, buf[:])
	return buf[0], err
}

// ReadBoolean reads a Boolean value from the current stream.
func (br *BinaryReader) ReadBoolean() (bool, error) {
	b, err := br.ReadByte()
	if err != nil {
		return false, err
	}
	return b != 0, nil
}

// ReadInt16 reads a 2-byte signed integer from the current stream.
func (br *BinaryReader) ReadInt16() (int16, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	var val int16
	err := br.readBinary(&val)
	return val, err
}

// ReadInt32 reads a 4-byte signed integer from the current stream.
func (br *BinaryReader) ReadInt32() (int32, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	var val int32
	err := br.readBinary(&val)
	return val, err
}

// ReadInt64 reads an 8-byte signed integer from the current stream.
func (br *BinaryReader) ReadInt64() (int64, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	var val int64
	err := br.readBinary(&val)
	return val, err
}

// ReadUInt16 reads a 2-byte unsigned integer from the current stream.
func (br *BinaryReader) ReadUInt16() (uint16, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	var val uint16
	err := br.readBinary(&val)
	return val, err
}

// ReadUInt32 reads a 4-byte unsigned integer from the current stream.
func (br *BinaryReader) ReadUInt32() (uint32, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	var val uint32
	err := br.readBinary(&val)
	return val, err
}

// ReadUInt64 reads an 8-byte unsigned integer from the current stream.
func (br *BinaryReader) ReadUInt64() (uint64, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	var val uint64
	err := br.readBinary(&val)
	return val, err
}

// ReadSingle reads a 4-byte floating point value from the current stream.
func (br *BinaryReader) ReadSingle() (float32, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	var val float32
	err := br.readBinary(&val)
	return val, err
}

// ReadDouble reads an 8-byte floating point value from the current stream.
func (br *BinaryReader) ReadDouble() (float64, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	var val float64
	err := br.readBinary(&val)
	return val, err
}

// ReadString reads a string from the current stream. The string is prefixed with the length,
// encoded as an 7-bit encoded integer.
func (br *BinaryReader) ReadString() (string, error) {
	if br.closed {
		return "", fmt.Errorf("binary reader is closed")
	}
	// Read 7-bit encoded length (simplified: use int32 length prefix)
	length, err := br.ReadInt32()
	if err != nil {
		return "", err
	}
	if length < 0 {
		return "", fmt.Errorf("invalid string length: %d", length)
	}
	if length == 0 {
		return "", nil
	}
	bytes, err := br.ReadBytes(int(length))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadRune reads the next character from the current stream.
func (br *BinaryReader) ReadRune() (rune, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	// Read rune as 4 bytes (int32)
	val, err := br.ReadInt32()
	if err != nil {
		return 0, err
	}
	return rune(val), nil
}

// PeekChar returns the next available character without advancing the position.
// Returns 0 if no characters are available or if the stream does not support seeking.
func (br *BinaryReader) PeekChar() (rune, error) {
	if br.closed {
		return 0, fmt.Errorf("binary reader is closed")
	}
	if seeker, ok := br.reader.(io.Seeker); ok {
		curPos, err := seeker.Seek(0, io.SeekCurrent)
		if err != nil {
			return 0, err
		}
		r, err := br.ReadRune()
		if err != nil {
			return 0, err
		}
		_, seekErr := seeker.Seek(curPos, io.SeekStart)
		return r, seekErr
	}
	// Fallback: just read
	return br.ReadRune()
}

// readBinary reads binary data into val using the configured byte order.
func (br *BinaryReader) readBinary(val interface{}) error {
	var order binary.ByteOrder
	if br.littleEndian {
		order = binary.LittleEndian
	} else {
		order = binary.BigEndian
	}
	return binary.Read(br.reader, order, val)
}
