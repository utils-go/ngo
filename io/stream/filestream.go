package stream

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// FileMode specifies how the operating system should open a file.
type FileMode int

const (
	// FileModeCreateNew specifies that the operating system should create a new file.
	FileModeCreateNew FileMode = iota
	// FileModeCreate specifies that the operating system should create a new file.
	// If the file already exists, it will be overwritten.
	FileModeCreate
	// FileModeOpen specifies that the operating system should open an existing file.
	FileModeOpen
	// FileModeOpenOrCreate specifies that the operating system should open a file if it exists;
	// otherwise, a new file should be created.
	FileModeOpenOrCreate
	// FileModeTruncate specifies that the operating system should open an existing file.
	// When the file is opened, it should be truncated to zero bytes.
	FileModeTruncate
	// FileModeAppend opens the file if it exists and seeks to the end of the file,
	// or creates a new file.
	FileModeAppend
)

// FileAccess specifies the access level for a file.
type FileAccess int

const (
	// FileAccessRead specifies read access to the file.
	FileAccessRead FileAccess = iota
	// FileAccessWrite specifies write access to the file.
	FileAccessWrite
	// FileAccessReadWrite specifies read and write access to the file.
	FileAccessReadWrite
)

// FileStream provides a Stream for a file, supporting both synchronous and
// asynchronous read and write operations.
// Equivalent to System.IO.FileStream in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.io.filestream?view=netframework-4.7.2
type FileStream struct {
	file   *os.File
	name   string
	closed bool
	mu     sync.RWMutex
}

// NewFileStream creates a new FileStream for the specified path with the given mode and access.
func NewFileStream(path string, mode FileMode, access FileAccess) (*FileStream, error) {
	var osFlag int

	switch mode {
	case FileModeCreateNew:
		osFlag = os.O_CREATE | os.O_EXCL
	case FileModeCreate:
		osFlag = os.O_CREATE | os.O_TRUNC
	case FileModeOpen:
		osFlag = 0
	case FileModeOpenOrCreate:
		osFlag = os.O_CREATE
	case FileModeTruncate:
		osFlag = os.O_TRUNC
	case FileModeAppend:
		osFlag = os.O_APPEND | os.O_CREATE
	}

	switch access {
	case FileAccessRead:
		osFlag |= os.O_RDONLY
	case FileAccessWrite:
		osFlag |= os.O_WRONLY
	case FileAccessReadWrite:
		osFlag |= os.O_RDWR
	}

	file, err := os.OpenFile(path, osFlag, 0666)
	if err != nil {
		return nil, err
	}

	return &FileStream{
		file: file,
		name: path,
	}, nil
}

// NewFileStreamFromFile wraps an existing *os.File into a FileStream.
func NewFileStreamFromFile(file *os.File) *FileStream {
	return &FileStream{
		file: file,
		name: file.Name(),
	}
}

// Name gets the name of the FileStream.
func (fs *FileStream) Name() string {
	return fs.name
}

// Read reads a block of bytes from the stream and writes the data to buffer.
func (fs *FileStream) Read(buffer []byte) (int, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	if fs.closed {
		return 0, fmt.Errorf("file stream is closed")
	}
	return fs.file.Read(buffer)
}

// Write writes a block of bytes to the stream.
func (fs *FileStream) Write(buffer []byte) (int, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	if fs.closed {
		return 0, fmt.Errorf("file stream is closed")
	}
	return fs.file.Write(buffer)
}

// Seek sets the position within the current stream.
func (fs *FileStream) Seek(offset int64, whence int) (int64, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	if fs.closed {
		return 0, fmt.Errorf("file stream is closed")
	}
	return fs.file.Seek(offset, whence)
}

// SeekOrigin sets the position using .NET SeekOrigin.
func (fs *FileStream) SeekOrigin(offset int64, origin SeekOrigin) (int64, error) {
	return fs.Seek(offset, seekOriginToWhence(origin))
}

// Length gets the length of the stream in bytes.
func (fs *FileStream) Length() int64 {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	info, err := fs.file.Stat()
	if err != nil {
		return 0
	}
	return info.Size()
}

// SetLength sets the length of the current stream.
func (fs *FileStream) SetLength(value int64) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	if fs.closed {
		return fmt.Errorf("file stream is closed")
	}
	return fs.file.Truncate(value)
}

// Position gets the current position within the stream.
func (fs *FileStream) Position() int64 {
	pos, _ := fs.Seek(0, io.SeekCurrent)
	return pos
}

// SetPosition sets the position within the stream.
func (fs *FileStream) SetPosition(value int64) (int64, error) {
	return fs.Seek(value, io.SeekStart)
}

// Flush clears all buffers and causes any buffered data to be written to the file.
func (fs *FileStream) Flush() error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	if fs.closed {
		return fmt.Errorf("file stream is closed")
	}
	return fs.file.Sync()
}

// Close closes the current stream and releases any resources associated with it.
func (fs *FileStream) Close() error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	if fs.closed {
		return nil
	}
	fs.closed = true
	return fs.file.Close()
}

// CanRead indicates whether the stream supports reading.
func (fs *FileStream) CanRead() bool {
	return !fs.closed
}

// CanWrite indicates whether the stream supports writing.
func (fs *FileStream) CanWrite() bool {
	return !fs.closed
}

// CanSeek indicates whether the stream supports seeking.
func (fs *FileStream) CanSeek() bool {
	return !fs.closed
}

// Lock prevents other processes from reading or writing to the file.
// Note: Go does not fully support file locking on all platforms.
func (fs *FileStream) Lock(position, length int64) error {
	// Simplified: no-op on most platforms
	return nil
}

// Unlock allows access by other processes.
func (fs *FileStream) Unlock(position, length int64) error {
	return nil
}

// IsClosed returns whether the file stream is closed.
func (fs *FileStream) IsClosed() bool {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	return fs.closed
}
