package stream

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// ---- StreamReader tests ----

func TestStreamReader_ReadLine(t *testing.T) {
	input := "line1\nline2\nline3\n"
	sr := NewStreamReader(bytes.NewReader([]byte(input)))

	lines := []string{}
	for {
		line, err := sr.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, line)
	}

	if len(lines) != 3 {
		t.Errorf("expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "line1" || lines[1] != "line2" || lines[2] != "line3" {
		t.Errorf("unexpected lines: %v", lines)
	}
}

func TestStreamReader_ReadToEnd(t *testing.T) {
	input := "hello world"
	sr := NewStreamReader(bytes.NewReader([]byte(input)))
	result, err := sr.ReadToEnd()
	if err != nil {
		t.Fatalf("ReadToEnd error: %v", err)
	}
	if result != input {
		t.Errorf("expected %q, got %q", input, result)
	}
}

func TestStreamReader_Peek(t *testing.T) {
	input := "ABC"
	sr := NewStreamReader(bytes.NewReader([]byte(input)))
	r, err := sr.Peek()
	if err != nil {
		t.Fatalf("Peek error: %v", err)
	}
	if r != 'A' {
		t.Errorf("expected 'A', got %c", r)
	}
	// Read should still return the same character
	var buf [1]byte
	sr.Read(buf[:], 0, 1)
	if buf[0] != 'A' {
		t.Errorf("expected 'A' after peek, got %c", buf[0])
	}
}

func TestStreamReader_Closed(t *testing.T) {
	sr := NewStreamReader(bytes.NewReader([]byte("test")))
	sr.Close()
	_, err := sr.ReadToEnd()
	if err == nil {
		t.Error("expected error reading from closed StreamReader")
	}
}

// ---- StreamWriter tests ----

func TestStreamWriter_Write(t *testing.T) {
	var buf bytes.Buffer
	sw := NewStreamWriter(&buf)
	err := sw.Write("hello")
	if err != nil {
		t.Fatalf("Write error: %v", err)
	}
	sw.Flush()
	if buf.String() != "hello" {
		t.Errorf("expected 'hello', got %q", buf.String())
	}
}

func TestStreamWriter_WriteLine(t *testing.T) {
	var buf bytes.Buffer
	sw := NewStreamWriter(&buf)
	sw.WriteLine("line1")
	sw.WriteLine("line2")
	sw.Flush()
	expected := "line1\nline2\n"
	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}

func TestStreamWriter_AutoFlush(t *testing.T) {
	var buf bytes.Buffer
	sw := NewStreamWriter(&buf)
	sw.SetAutoFlush(true)
	sw.Write("test")
	if buf.String() != "test" {
		t.Errorf("expected 'test', got %q", buf.String())
	}
}

// ---- MemoryStream tests ----

func TestMemoryStream_ReadWrite(t *testing.T) {
	ms := NewMemoryStream()
	data := []byte("hello")
	n, err := ms.Write(data)
	if err != nil || n != len(data) {
		t.Fatalf("Write error: %v", err)
	}

	ms.SetPosition(0)
	buf := make([]byte, len(data))
	n, err = ms.Read(buf)
	if err != nil || n != len(data) {
		t.Fatalf("Read error: %v", err)
	}
	if string(buf) != "hello" {
		t.Errorf("expected 'hello', got %q", string(buf))
	}
}

func TestMemoryStream_ToArray(t *testing.T) {
	ms := NewMemoryStream()
	ms.Write([]byte{1, 2, 3, 4})
	arr := ms.ToArray()
	if len(arr) != 4 || arr[0] != 1 {
		t.Errorf("unexpected ToArray result: %v", arr)
	}
}

func TestMemoryStream_Capacity(t *testing.T) {
	ms := NewMemoryStreamWithCapacity(100)
	if ms.Capacity() < 100 {
		t.Errorf("expected capacity >= 100, got %d", ms.Capacity())
	}
}

func TestMemoryStream_WriteTo(t *testing.T) {
	ms := NewMemoryStream()
	ms.Write([]byte("from-memory"))
	
	var buf bytes.Buffer
	_, err := ms.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo error: %v", err)
	}
	if buf.String() != "from-memory" {
		t.Errorf("expected 'from-memory', got %q", buf.String())
	}
}

// ---- BinaryReader/BinaryWriter tests ----

func TestBinaryReaderWriter_RoundTrip(t *testing.T) {
	var buf bytes.Buffer
	bw := NewBinaryWriter(&buf)

	// Write various types
	bw.WriteInt32(42)
	bw.WriteDouble(3.14)
	bw.WriteString("hello")
	bw.WriteBoolean(true)
	bw.WriteInt16(-8)

	br := NewBinaryReader(bytes.NewReader(buf.Bytes()))

	v32, err := br.ReadInt32()
	if err != nil || v32 != 42 {
		t.Errorf("ReadInt32: expected 42, got %d (err=%v)", v32, err)
	}

	v64, err := br.ReadDouble()
	if err != nil || v64 != 3.14 {
		t.Errorf("ReadDouble: expected 3.14, got %f (err=%v)", v64, err)
	}

	s, err := br.ReadString()
	if err != nil || s != "hello" {
		t.Errorf("ReadString: expected 'hello', got %q (err=%v)", s, err)
	}

	b, err := br.ReadBoolean()
	if err != nil || !b {
		t.Errorf("ReadBoolean: expected true, got %v (err=%v)", b, err)
	}

	v16, err := br.ReadInt16()
	if err != nil || v16 != -8 {
		t.Errorf("ReadInt16: expected -8, got %d (err=%v)", v16, err)
	}
}

func TestBinaryReader_ReadBytes(t *testing.T) {
	data := []byte{10, 20, 30, 40}
	br := NewBinaryReader(bytes.NewReader(data))
	result, err := br.ReadBytes(3)
	if err != nil {
		t.Fatalf("ReadBytes error: %v", err)
	}
	if len(result) != 3 || result[0] != 10 || result[1] != 20 || result[2] != 30 {
		t.Errorf("unexpected result: %v", result)
	}
}

// ---- FileStream tests ----

func TestFileStream_ReadWrite(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.txt")

	fs, err := NewFileStream(path, FileModeCreate, FileAccessReadWrite)
	if err != nil {
		t.Fatalf("NewFileStream error: %v", err)
	}

	data := []byte("filestream test")
	_, err = fs.Write(data)
	if err != nil {
		t.Fatalf("Write error: %v", err)
	}
	fs.Close()

	// Read back
	fs2, err := NewFileStream(path, FileModeOpen, FileAccessRead)
	if err != nil {
		t.Fatalf("NewFileStream (read) error: %v", err)
	}
	defer fs2.Close()

	buf := make([]byte, len(data))
	_, err = fs2.Read(buf)
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if string(buf) != string(data) {
		t.Errorf("expected %q, got %q", string(data), string(buf))
	}
}

func TestFileStream_Seek(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "seek_test.txt")
	os.WriteFile(path, []byte("0123456789"), 0644)

	fs, err := NewFileStream(path, FileModeOpen, FileAccessRead)
	if err != nil {
		t.Fatalf("NewFileStream error: %v", err)
	}
	defer fs.Close()

	pos, err := fs.Seek(5, 0)
	if err != nil || pos != 5 {
		t.Errorf("Seek: expected pos=5, got %d (err=%v)", pos, err)
	}

	buf := make([]byte, 3)
	fs.Read(buf)
	if string(buf) != "567" {
		t.Errorf("expected '567', got %q", string(buf))
	}
}

func TestFileStream_Length(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "len_test.txt")
	os.WriteFile(path, []byte("01234"), 0644)

	fs, _ := NewFileStream(path, FileModeOpen, FileAccessRead)
	defer fs.Close()

	if fs.Length() != 5 {
		t.Errorf("expected length 5, got %d", fs.Length())
	}
}

func TestFileStream_ClosedError(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "closed_test.txt")
	os.WriteFile(path, []byte("test"), 0644)

	fs, _ := NewFileStream(path, FileModeOpen, FileAccessRead)
	fs.Close()

	_, err := fs.Read(make([]byte, 4))
	if err == nil {
		t.Error("expected error reading from closed FileStream")
	}
}
