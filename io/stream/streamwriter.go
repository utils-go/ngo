package stream

import (
	"bufio"
	"fmt"
	"io"
)

// StreamWriter implements a TextWriter for writing characters to a stream
// in a particular encoding. Equivalent to System.IO.StreamWriter in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.io.streamwriter?view=netframework-4.7.2
type StreamWriter struct {
	writer    *bufio.Writer
	base      io.Writer
	closed    bool
	autoFlush bool
	newLine   string
}

// NewStreamWriter creates a new StreamWriter for the specified stream.
func NewStreamWriter(stream io.Writer) *StreamWriter {
	return &StreamWriter{
		writer:    bufio.NewWriter(stream),
		base:      stream,
		autoFlush: false,
		newLine:   "\n",
	}
}

// NewStreamWriterWithBuffer creates a new StreamWriter with the specified buffer size.
func NewStreamWriterWithBuffer(stream io.Writer, bufferSize int) *StreamWriter {
	return &StreamWriter{
		writer:    bufio.NewWriterSize(stream, bufferSize),
		base:      stream,
		autoFlush: false,
		newLine:   "\n",
	}
}

// BaseStream returns the underlying stream.
func (sw *StreamWriter) BaseStream() io.Writer {
	return sw.base
}

// AutoFlush gets or sets a value indicating whether the StreamWriter will flush
// its buffer after every Write call.
func (sw *StreamWriter) AutoFlush() bool {
	return sw.autoFlush
}

// SetAutoFlush sets whether the StreamWriter should flush after every Write.
func (sw *StreamWriter) SetAutoFlush(value bool) {
	sw.autoFlush = value
}

// NewLine gets or sets the line terminator string used by the current StreamWriter.
func (sw *StreamWriter) NewLine() string {
	return sw.newLine
}

// SetNewLine sets the line terminator string.
func (sw *StreamWriter) SetNewLine(value string) {
	sw.newLine = value
}

// Write writes a string to the stream.
// Equivalent to StreamWriter.Write in .NET.
func (sw *StreamWriter) Write(value string) error {
	if sw.closed {
		return fmt.Errorf("stream writer is closed")
	}
	_, err := sw.writer.WriteString(value)
	if err != nil {
		return err
	}
	if sw.autoFlush {
		return sw.Flush()
	}
	return nil
}

// WriteBytes writes bytes to the stream.
func (sw *StreamWriter) WriteBytes(data []byte) error {
	if sw.closed {
		return fmt.Errorf("stream writer is closed")
	}
	_, err := sw.writer.Write(data)
	if err != nil {
		return err
	}
	if sw.autoFlush {
		return sw.Flush()
	}
	return nil
}

// WriteLine writes a string followed by a line terminator to the text stream.
// Equivalent to StreamWriter.WriteLine in .NET.
func (sw *StreamWriter) WriteLine(value string) error {
	return sw.Write(value + sw.newLine)
}

// WriteLineBytes writes bytes followed by a line terminator.
func (sw *StreamWriter) WriteLineBytes(data []byte) error {
	if err := sw.WriteBytes(data); err != nil {
		return err
	}
	return sw.Write(sw.newLine)
}

// Flush clears all buffers for the current writer and causes any buffered data
// to be written to the underlying stream.
// Equivalent to StreamWriter.Flush in .NET.
func (sw *StreamWriter) Flush() error {
	if sw.closed {
		return fmt.Errorf("stream writer is closed")
	}
	return sw.writer.Flush()
}

// Close closes the current StreamWriter and the underlying stream.
// Equivalent to StreamWriter.Close in .NET.
func (sw *StreamWriter) Close() error {
	if sw.closed {
		return nil
	}
	sw.closed = true
	if err := sw.writer.Flush(); err != nil {
		return err
	}
	if closer, ok := sw.base.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
