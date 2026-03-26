package fileinfo

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/utils-go/ngo/io/file"
)

// FileInfo 提供文件的属性和操作方法，对应.NET中的System.IO.FileInfo
// 参考: https://learn.microsoft.com/en-us/dotnet/api/system.io.fileinfo?view=netframework-4.7.2

// FileInfo 结构体表示文件的信息
type FileInfo struct {
	Path           string    // 文件路径
	Name           string    // 文件名
	FullName       string    // 完整路径
	DirectoryName  string    // 目录名
	Extension      string    // 扩展名
	Length         int64     // 文件大小（字节）
	CreationTime   time.Time // 创建时间
	LastWriteTime  time.Time // 最后写入时间
	LastAccessTime time.Time // 最后访问时间
	Exists         bool      // 文件是否存在
}

// FileAccess 指定文件的读写访问权限
type FileAccess int

const (
	// FileAccessRead 只读访问
	FileAccessRead FileAccess = iota
	// FileAccessWrite 只写访问
	FileAccessWrite
	// FileAccessReadWrite 读写访问
	FileAccessReadWrite
)

// FileShare 指定文件共享级别
type FileShare int

const (
	// FileShareNone 不共享
	FileShareNone FileShare = iota
	// FileShareRead 允许其他进程读取
	FileShareRead
	// FileShareWrite 允许其他进程写入
	FileShareWrite
	// FileShareReadWrite 允许其他进程读写
	FileShareReadWrite
)

// FileMode 指定文件打开模式
type FileMode int

const (
	// FileModeCreateNew 创建新文件，如果文件已存在则失败
	FileModeCreateNew FileMode = iota
	// FileModeCreate 创建文件，如果文件已存在则覆盖
	FileModeCreate
	// FileModeOpen 打开现有文件
	FileModeOpen
	// FileModeOpenOrCreate 打开文件，如果文件不存在则创建
	FileModeOpenOrCreate
	// FileModeTruncate 打开文件并截断为0字节
	FileModeTruncate
	// FileModeAppend 打开文件并在末尾追加
	FileModeAppend
)

// FileStream 表示文件流
type FileStream struct {
	file *os.File
}

// Close 关闭文件流
func (fs *FileStream) Close() error {
	return fs.file.Close()
}

// Read 从文件流读取数据
func (fs *FileStream) Read(buffer []byte) (int, error) {
	return fs.file.Read(buffer)
}

// Write 向文件流写入数据
func (fs *FileStream) Write(buffer []byte) (int, error) {
	return fs.file.Write(buffer)
}

// Seek 设置文件指针位置
func (fs *FileStream) Seek(offset int64, whence int) (int64, error) {
	return fs.file.Seek(offset, whence)
}

// NewFileInfo 创建新的FileInfo实例
// 参数:
//
//	path: 文件路径
//
// 返回值:
//
//	*FileInfo: 文件信息对象
func newFileInfo(path string) *FileInfo {
	info := &FileInfo{
		Path: path,
	}

	// 获取文件信息
	fileInfo, err := os.Stat(path)
	if err == nil {
		info.Exists = true
		info.Name = fileInfo.Name()
		info.FullName, _ = filepath.Abs(path)
		info.DirectoryName = filepath.Dir(path)
		info.Extension = filepath.Ext(path)
		info.Length = fileInfo.Size()
		info.CreationTime = fileInfo.ModTime()
		info.LastWriteTime = fileInfo.ModTime()
		info.LastAccessTime = fileInfo.ModTime()
	} else {
		info.Exists = false
		info.Name = filepath.Base(path)
		info.FullName, _ = filepath.Abs(path)
		info.DirectoryName = filepath.Dir(path)
		info.Extension = filepath.Ext(path)
	}

	return info
}

// GetFileInfo 获取文件的完整信息
// 参数:
//
//	path: 文件路径
//
// 返回值:
//
//	*FileInfo: 文件信息对象
func GetFileInfo(path string) *FileInfo {
	return newFileInfo(path)
}

// Create 创建新文件
// 参数:
//
//	path: 文件路径
//
// 返回值:
//
//	*FileStream: 文件流
//	error: 错误信息
func Create(path string) (*FileStream, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &FileStream{file: file}, nil
}

// CreateText 创建文本文件并返回写入器
// 参数:
//
//	path: 文件路径
//
// 返回值:
//
//	*bufio.Writer: 文本写入器
//	error: 错误信息
func CreateText(path string) (*bufio.Writer, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return bufio.NewWriter(file), nil
}

// Open 打开文件
// 参数:
//
//	path: 文件路径
//	mode: 打开模式
//	access: 访问权限
//	share: 共享级别
//
// 返回值:
//
//	*FileStream: 文件流
//	error: 错误信息
func Open(path string, mode FileMode, access FileAccess, share FileShare) (*FileStream, error) {
	var osFlag int
	var osPerm os.FileMode

	switch mode {
	case FileModeCreateNew:
		osFlag = os.O_CREATE | os.O_EXCL | os.O_RDWR
	case FileModeCreate:
		osFlag = os.O_CREATE | os.O_TRUNC | os.O_RDWR
	case FileModeOpen:
		osFlag = os.O_RDWR
	case FileModeOpenOrCreate:
		osFlag = os.O_CREATE | os.O_RDWR
	case FileModeTruncate:
		osFlag = os.O_TRUNC | os.O_RDWR
	case FileModeAppend:
		osFlag = os.O_APPEND | os.O_WRONLY
	}

	switch access {
	case FileAccessRead:
		osFlag = osFlag &^ (os.O_WRONLY | os.O_RDWR)
		osFlag = osFlag | os.O_RDONLY
	case FileAccessWrite:
		osFlag = osFlag &^ (os.O_RDONLY | os.O_RDWR)
		osFlag = osFlag | os.O_WRONLY
	case FileAccessReadWrite:
		osFlag = osFlag &^ (os.O_RDONLY | os.O_WRONLY)
		osFlag = osFlag | os.O_RDWR
	}

	osPerm = 0666

	file, err := os.OpenFile(path, osFlag, osPerm)
	if err != nil {
		return nil, err
	}

	return &FileStream{file: file}, nil
}

// OpenRead 打开只读文件
// 参数:
//
//	path: 文件路径
//
// 返回值:
//
//	*FileStream: 文件流
//	error: 错误信息
func OpenRead(path string) (*FileStream, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &FileStream{file: file}, nil
}

// OpenText 打开文本文件并返回读取器
// 参数:
//
//	path: 文件路径
//
// 返回值:
//
//	*bufio.Reader: 文本读取器
//	error: 错误信息
func OpenText(path string) (*bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(file), nil
}

// OpenWrite 打开只写文件
// 参数:
//
//	path: 文件路径
//
// 返回值:
//
//	*FileStream: 文件流
//	error: 错误信息
func OpenWrite(path string) (*FileStream, error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	return &FileStream{file: file}, nil
}

// CopyTo 将文件复制到目标路径
// 参数:
//
//	sourcePath: 源文件路径
//	destinationPath: 目标文件路径
//	overwrite: 是否覆盖已存在的文件
//
// 返回值:
//
//	error: 错误信息
func CopyTo(sourcePath, destinationPath string, overwrite bool) error {
	if file.Exists(destinationPath) && !overwrite {
		return os.ErrExist
	}

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

// Delete 删除文件
// 参数:
//
//	path: 文件路径
//
// 返回值:
//
//	error: 错误信息
func Delete(path string) error {
	return os.Remove(path)
}

// MoveTo 将文件移动到目标路径
// 参数:
//
//	sourcePath: 源文件路径
//	destinationPath: 目标文件路径
//	overwrite: 是否覆盖已存在的文件
//
// 返回值:
//
//	error: 错误信息
func MoveTo(sourcePath, destinationPath string, overwrite bool) error {
	if file.Exists(destinationPath) && !overwrite {
		return os.ErrExist
	}

	return os.Rename(sourcePath, destinationPath)
}
