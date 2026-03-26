package fileinfo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewFileInfo(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	content := "Hello, FileInfo!"
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	info := newFileInfo(tempFile.Name())

	if !info.Exists {
		t.Error("文件应该存在")
	}

	if info.Name == "" {
		t.Error("文件名不应该为空")
	}

	if info.FullName == "" {
		t.Error("完整路径不应该为空")
	}

	if info.DirectoryName == "" {
		t.Error("目录名不应该为空")
	}

	if info.Extension == "" {
		t.Error("扩展名不应该为空")
	}

	if info.Length != int64(len(content)) {
		t.Errorf("文件长度应该为%d，实际为%d", len(content), info.Length)
	}

	if info.CreationTime.IsZero() {
		t.Error("创建时间不应该为零")
	}

	if info.LastWriteTime.IsZero() {
		t.Error("最后写入时间不应该为零")
	}

	if info.LastAccessTime.IsZero() {
		t.Error("最后访问时间不应该为零")
	}
}

func TestGetFileInfo(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	content := "Test content"
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	info := GetFileInfo(tempFile.Name())

	if !info.Exists {
		t.Error("文件应该存在")
	}

	expectedName := filepath.Base(tempFile.Name())
	if info.Name != expectedName {
		t.Errorf("文件名应该为'%s'，实际为'%s'", expectedName, info.Name)
	}

	if info.Length != int64(len(content)) {
		t.Errorf("文件长度应该为%d，实际为%d", len(content), info.Length)
	}
}

func TestNewFileInfoNonExistent(t *testing.T) {
	info := newFileInfo("non_existent_file.txt")

	if info.Exists {
		t.Error("不存在的文件应该返回Exists为false")
	}

	if info.Name == "" {
		t.Error("文件名不应该为空")
	}

	if info.FullName == "" {
		t.Error("完整路径不应该为空")
	}
}

func TestCreate(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_create_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	fs, err := Create(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer fs.Close()

	if fs == nil {
		t.Fatal("文件流不应该为nil")
	}
}

func TestCreateText(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_createtext_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	writer, err := CreateText(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if writer == nil {
		t.Fatal("写入器不应该为nil")
	}

	_, err = writer.WriteString("Test content")
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Flush()
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenRead(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_openread_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	content := "Test content for reading"
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	fs, err := OpenRead(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer fs.Close()

	if fs == nil {
		t.Fatal("文件流不应该为nil")
	}

	buffer := make([]byte, len(content))
	n, err := fs.Read(buffer)
	if err != nil {
		t.Fatal(err)
	}

	if string(buffer[:n]) != content {
		t.Errorf("读取的内容应该为'%s'，实际为'%s'", content, string(buffer[:n]))
	}
}

func TestOpenText(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_opentext_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	content := "Test content for text reading"
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	reader, err := OpenText(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if reader == nil {
		t.Fatal("读取器不应该为nil")
	}

	readContent, err := reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		t.Fatal(err)
	}

	if readContent != content {
		t.Errorf("读取的内容应该为'%s'，实际为'%s'", content, readContent)
	}
}

func TestOpenWrite(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_openwrite_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	fs, err := OpenWrite(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer fs.Close()

	if fs == nil {
		t.Fatal("文件流不应该为nil")
	}

	_, err = fs.Write([]byte("New content"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenWithParameters(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_open_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	content := "Original content"
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	fs, err := Open(tempFile.Name(), FileModeOpen, FileAccessRead, FileShareRead)
	if err != nil {
		t.Fatal(err)
	}
	defer fs.Close()

	if fs == nil {
		t.Fatal("文件流不应该为nil")
	}

	buffer := make([]byte, len(content))
	n, err := fs.Read(buffer)
	if err != nil {
		t.Fatal(err)
	}

	if string(buffer[:n]) != content {
		t.Errorf("读取的内容应该为'%s'，实际为'%s'", content, string(buffer[:n]))
	}
}

func TestCopyTo(t *testing.T) {
	sourceFile, err := os.CreateTemp("", "fileinfo_source_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(sourceFile.Name())

	content := "Content to copy"
	_, err = sourceFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if err := sourceFile.Close(); err != nil {
		t.Fatal(err)
	}

	destinationFile, err := os.CreateTemp("", "fileinfo_destination_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(destinationFile.Name())

	err = CopyTo(sourceFile.Name(), destinationFile.Name(), true)
	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(destinationFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != content {
		t.Errorf("复制的内容应该为'%s'，实际为'%s'", content, string(data))
	}
}

func TestCopyToNoOverwrite(t *testing.T) {
	sourceFile, err := os.CreateTemp("", "fileinfo_source_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(sourceFile.Name())

	destinationFile, err := os.CreateTemp("", "fileinfo_destination_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(destinationFile.Name())

	_, err = destinationFile.WriteString("Destination content")
	if err != nil {
		t.Fatal(err)
	}
	if err := destinationFile.Close(); err != nil {
		t.Fatal(err)
	}

	err = CopyTo(sourceFile.Name(), destinationFile.Name(), false)
	if err != os.ErrExist {
		t.Errorf("当overwrite为false且目标文件存在时，应该返回os.ErrExist，实际为%v", err)
	}
}

func TestDelete(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_delete_test")
	if err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	err = Delete(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(tempFile.Name())
	if err == nil {
		t.Error("文件应该已被删除")
	}
}

func TestMoveTo(t *testing.T) {
	sourceFile, err := os.CreateTemp("", "fileinfo_move_source_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(sourceFile.Name())

	content := "Content to move"
	_, err = sourceFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if err := sourceFile.Close(); err != nil {
		t.Fatal(err)
	}

	destinationFile, err := os.CreateTemp("", "fileinfo_move_destination_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(destinationFile.Name())
	if err := destinationFile.Close(); err != nil {
		t.Fatal(err)
	}

	err = MoveTo(sourceFile.Name(), destinationFile.Name(), true)
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(sourceFile.Name())
	if err == nil {
		t.Error("源文件应该已被移动")
	}

	data, err := os.ReadFile(destinationFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != content {
		t.Errorf("移动的内容应该为'%s'，实际为'%s'", content, string(data))
	}
}

func TestMoveToNoOverwrite(t *testing.T) {
	sourceFile, err := os.CreateTemp("", "fileinfo_move_source_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(sourceFile.Name())

	destinationFile, err := os.CreateTemp("", "fileinfo_move_destination_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(destinationFile.Name())

	_, err = destinationFile.WriteString("Destination content")
	if err != nil {
		t.Fatal(err)
	}
	if err := destinationFile.Close(); err != nil {
		t.Fatal(err)
	}

	err = MoveTo(sourceFile.Name(), destinationFile.Name(), false)
	if err != os.ErrExist {
		t.Errorf("当overwrite为false且目标文件存在时，应该返回os.ErrExist，实际为%v", err)
	}
}

func TestFileStreamOperations(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_stream_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	fs, err := OpenWrite(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	content := "Test stream content"
	_, err = fs.Write([]byte(content))
	if err != nil {
		t.Fatal(err)
	}

	err = fs.Close()
	if err != nil {
		t.Fatal(err)
	}

	fs, err = OpenRead(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer fs.Close()

	buffer := make([]byte, len(content))
	n, err := fs.Read(buffer)
	if err != nil {
		t.Fatal(err)
	}

	if string(buffer[:n]) != content {
		t.Errorf("读取的内容应该为'%s'，实际为'%s'", content, string(buffer[:n]))
	}

	_, err = fs.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileStreamWrite(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_stream_write_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	fs, err := Create(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	content := "Write test content"
	_, err = fs.Write([]byte(content))
	if err != nil {
		t.Fatal(err)
	}

	err = fs.Close()
	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != content {
		t.Errorf("写入的内容应该为'%s'，实际为'%s'", content, string(data))
	}
}

func TestFileStreamSeek(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_stream_seek_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	fs, err := Create(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	content := "Hello, World!"
	_, err = fs.Write([]byte(content))
	if err != nil {
		t.Fatal(err)
	}

	err = fs.Close()
	if err != nil {
		t.Fatal(err)
	}

	fs, err = OpenRead(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer fs.Close()

	_, err = fs.Seek(7, 0)
	if err != nil {
		t.Fatal(err)
	}

	buffer := make([]byte, 5)
	n, err := fs.Read(buffer)
	if err != nil {
		t.Fatal(err)
	}

	if string(buffer[:n]) != "World" {
		t.Errorf("seek后读取的内容应该为'World'，实际为'%s'", string(buffer[:n]))
	}
}

func TestCreateTextWriteAndRead(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_createtext_read_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	writer, err := CreateText(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	content := "Hello, World!"
	_, err = writer.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Flush()
	if err != nil {
		t.Fatal(err)
	}

	reader, err := OpenText(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	readContent, err := reader.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		t.Fatal(err)
	}

	if readContent != content {
		t.Errorf("读取的内容应该为'%s'，实际为'%s'", content, readContent)
	}
}

func TestOpenWithDifferentModes(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_open_modes_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	content := "Original content"
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	fs, err := Open(tempFile.Name(), FileModeOpen, FileAccessReadWrite, FileShareReadWrite)
	if err != nil {
		t.Fatal(err)
	}
	defer fs.Close()

	if fs == nil {
		t.Fatal("文件流不应该为nil")
	}
}

func TestOpenWithAppendMode(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_open_append_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	content := "Original content"
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	fs, err := Open(tempFile.Name(), FileModeAppend, FileAccessWrite, FileShareReadWrite)
	if err != nil {
		t.Fatal(err)
	}

	appendContent := " Appended content"
	_, err = fs.Write([]byte(appendContent))
	if err != nil {
		t.Fatal(err)
	}

	err = fs.Close()
	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := content + appendContent
	if string(data) != expectedContent {
		t.Errorf("追加后的内容应该为'%s'，实际为'%s'", expectedContent, string(data))
	}
}

func TestFileInfoStructure(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_structure_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	content := "Structure test"
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	info := newFileInfo(tempFile.Name())

	if info.Path != tempFile.Name() {
		t.Errorf("Path应该为'%s'，实际为'%s'", tempFile.Name(), info.Path)
	}

	if info.Name == "" {
		t.Error("Name不应该为空")
	}

	if info.FullName == "" {
		t.Error("FullName不应该为空")
	}

	if info.DirectoryName == "" {
		t.Error("DirectoryName不应该为空")
	}

	if info.Extension == "" {
		t.Error("Extension不应该为空")
	}

	if info.Length != int64(len(content)) {
		t.Errorf("Length应该为%d，实际为%d", len(content), info.Length)
	}

	if !info.Exists {
		t.Error("Exists应该为true")
	}
}

func TestCreateTextWithMultipleWrites(t *testing.T) {
	tempFile, err := os.CreateTemp("", "fileinfo_createtext_multi_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	writer, err := CreateText(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	lines := []string{"Line 1", "Line 2", "Line 3"}
	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			t.Fatal(err)
		}
	}

	err = writer.Flush()
	if err != nil {
		t.Fatal(err)
	}

	reader, err := OpenText(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	for i, expectedLine := range lines {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			t.Fatal(err)
		}

		expected := expectedLine + "\n"
		if line != expected {
			t.Errorf("第%d行应该为'%s'，实际为'%s'", i+1, expected, line)
		}
	}
}
