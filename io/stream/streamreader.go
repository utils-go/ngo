package stream

import (
	"bufio"
	"fmt"
	"io"
	"unicode/utf8"
)

// StreamReader implements a TextReader that reads characters from a byte stream
// in a particular encoding. Equivalent to System.IO.StreamReader in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.io.streamreader?view=netframework-4.7.2
type StreamReader struct {
	reader    *bufio.Reader
	base      io.Reader
	closed    bool
	endOfStream bool
}

// NewStreamReader creates a new StreamReader for the specified stream.
// The reader uses UTF-8 encoding by default.
func NewStreamReader(stream io.Reader) *StreamReader {
	return &StreamReader{
		reader: bufio.NewReader(stream),
		base:   stream,
	}
}

// NewStreamReaderWithBuffer creates a new StreamReader with the specified buffer size.
func NewStreamReaderWithBuffer(stream io.Reader, bufferSize int) *StreamReader {
	return &StreamReader{
		reader: bufio.NewReaderSize(stream, bufferSize),
		base:   stream,
	}
}

// BaseStream returns the underlying stream.
func (sr *StreamReader) BaseStream() io.Reader {
	return sr.base
}

// EndOfStream gets a value indicating whether the reader is at the end of the stream.
func (sr *StreamReader) EndOfStream() bool {
	return sr.endOfStream
}

// Read reads the next character or next set of characters from the input stream.
// Equivalent to Stream.Read in .NET.
// Parameters:
//   buffer - the buffer to read data into
//   index - the starting index in buffer
//   count - the number of characters to read
// Returns the number of characters read.
func (sr *StreamReader) Read(buffer []byte, index, count int) (int, error) {
	if sr.closed {
		return 0, fmt.Errorf("stream reader is closed")
	}
	if index < 0 || index > len(buffer) {
		return 0, fmt.Errorf("index %d is out of range", index)
	}
	if count <= 0 {
		return 0, nil
	}
	if index+count > len(buffer) {
		count = len(buffer) - index
	}

	totalRead := 0
	for totalRead < count {
		n, err := sr.reader.Read(buffer[index+totalRead : index+count])
		totalRead += n
		if err == io.EOF {
			sr.endOfStream = true
			break
		}
		if err != nil {
			return totalRead, err
		}
	}
	return totalRead, nil
}

// ReadLine reads a line of characters from the current stream and returns the data as a string.
// Equivalent to StreamReader.ReadLine in .NET.
func (sr *StreamReader) ReadLine() (string, error) {
	if sr.closed {
		return "", fmt.Errorf("stream reader is closed")
	}
	line, err := sr.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	if err == io.EOF {
		sr.endOfStream = true
		if len(line) == 0 {
			return "", err
		}
	}
	// Strip trailing newline characters
	if len(line) > 0 && line[len(line)-1] == '\n' {
		line = line[:len(line)-1]
	}
	if len(line) > 0 && line[len(line)-1] == '\r' {
		line = line[:len(line)-1]
	}
	return line, nil
}

// ReadToEnd reads all characters from the current position to the end of the stream.
// Equivalent to StreamReader.ReadToEnd in .NET.
func (sr *StreamReader) ReadToEnd() (string, error) {
	if sr.closed {
		return "", fmt.Errorf("stream reader is closed")
	}
	data, err := io.ReadAll(sr.reader)
	if err != nil {
		return "", err
	}
	sr.endOfStream = true
	return string(data), nil
}

// ReadBlock reads a specified maximum number of characters from the current stream
// and writes the data to a buffer.
func (sr *StreamReader) ReadBlock(buffer []rune, index, count int) (int, error) {
	if sr.closed {
		return 0, fmt.Errorf("stream reader is closed")
	}
	// Simplified: read bytes and convert to runes
	bytesBuf := make([]byte, count)
	n, err := sr.Read(bytesBuf, 0, count)
	if n > 0 {
		runes := []rune(string(bytesBuf[:n]))
		copyLen := len(runes)
		if index+copyLen > len(buffer) {
			copyLen = len(buffer) - index
		}
		copy(buffer[index:], runes[:copyLen])
		return copyLen, err
	}
	return 0, err
}

// Peek returns the next available character but does not consume it.
// Equivalent to StreamReader.Peek in .NET.
func (sr *StreamReader) Peek() (rune, error) {
	if sr.closed {
		return 0, fmt.Errorf("stream reader is closed")
	}
	// Try peeking up to 4 bytes (max UTF-8 rune length)
	bytes, err := sr.reader.Peek(1)
	if err != nil {
		if err == io.EOF {
			sr.endOfStream = true
		}
		return 0, err
	}
	if len(bytes) > 0 {
		// Determine how many more bytes to peek based on first byte
		peekLen := utf8ByteCount(bytes[0])
		if peekLen > 1 {
			moreBytes, _ := sr.reader.Peek(peekLen)
			if len(moreBytes) > 0 {
				bytes = moreBytes
			}
		}
		r, _ := readRuneFromBytes(bytes)
		return r, nil
	}
	return 0, io.EOF
}

// DiscardBufferedData discards the data currently buffered by the StreamReader.
func (sr *StreamReader) DiscardBufferedData() {
	sr.reader.Reset(sr.base)
}

// Close closes the StreamReader and the underlying stream.
func (sr *StreamReader) Close() error {
	if sr.closed {
		return nil
	}
	sr.closed = true
	if closer, ok := sr.base.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// readRuneFromBytes tries to read a single rune from byte slice.
func readRuneFromBytes(b []byte) (rune, int) {
	if len(b) == 0 {
		return 0, 0
	}
	r, size := utf8.DecodeRune(b)
	return r, size
}

// utf8ByteCount returns the expected length of a UTF-8 sequence based on the first byte.
func utf8ByteCount(firstByte byte) int {
	switch {
	case firstByte < 0x80:
		return 1
	case firstByte&0xE0 == 0xC0:
		return 2
	case firstByte&0xF0 == 0xE0:
		return 3
	case firstByte&0xF8 == 0xF0:
		return 4
	default:
		return 1
	}
}
