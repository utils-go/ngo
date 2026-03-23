package file

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 参考: https://learn.microsoft.com/en-us/dotnet/api/system.io?view=netframework-4.7.2

const nextLine string = "\r\n"

// readFile 读取文件内容
// 参数:
//   path: 文件路径
// 返回值:
//   []byte: 文件内容
//   error: 读取过程中的错误
func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// copyFileWithFlag 复制文件，使用指定的标志
// 参数:
//   srcFile: 源文件路径
//   dstFile: 目标文件路径
//   flag: 文件打开标志
// 返回值:
//   error: 复制过程中的错误
func copyFileWithFlag(srcFile, dstFile string, flag int) error {
	dir := filepath.Dir(dstFile)
	if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
		// 文件夹不存在，创建
		return err
	}

	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	dst, err := os.OpenFile(dstFile, flag, os.ModePerm) // file must not exist
	if err != nil {
		return err
	}
	defer src.Close()
	defer dst.Close()

	srcReader := bufio.NewReader(src)
	dstWriter := bufio.NewWriter(dst)
	for {
		tempBuf := make([]byte, 1024)
		n, err := srcReader.Read(tempBuf)
		if err != nil {
			if err == io.EOF {
				dstWriter.Write(tempBuf[:n])
				break
			} else {
				return err
			}
		}
		dstWriter.Write(tempBuf[:n])
	}
	if err := dstWriter.Flush(); err != nil {
		return err
	}
	return nil
}

// writeFile 写入文件内容
// 参数:
//   path: 文件路径
//   contents: 要写入的内容
// 返回值:
//   error: 写入过程中的错误
func writeFile(path string, contents []byte) error {
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	if _, err := buf.Write(contents); err != nil {
		return err
	}
	if err := buf.Flush(); err != nil { // 保存
		return err
	}
	return nil
}

// appendFile 追加文件内容
// 参数:
//   path: 文件路径
//   contents: 要追加的内容
// 返回值:
//   error: 追加过程中的错误
func appendFile(path string, contents []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	if _, err := buf.Write(contents); err != nil {
		return err
	}
	if err := buf.Flush(); err != nil {
		return err
	}
	return nil
}

// Copy 复制文件，如果目标文件已存在则失败
// 参数:
//   srcFile: 源文件路径
//   dstFile: 目标文件路径
// 返回值:
//   error: 复制过程中的错误
func Copy(srcFile string, dstFile string) error {
	return copyFileWithFlag(srcFile, dstFile, os.O_EXCL|os.O_CREATE)
}

// CopyWithOverwrite 复制文件，覆盖已存在的目标文件
// 参数:
//   srcFile: 源文件路径
//   dstFile: 目标文件路径
// 返回值:
//   error: 复制过程中的错误
func CopyWithOverwrite(srcFile string, dstFile string) error {
	return copyFileWithFlag(srcFile, dstFile, os.O_CREATE)
}

// Move 移动文件，如果目标文件已存在则失败
// 参数:
//   srcFile: 源文件路径
//   dstFile: 目标文件路径
// 返回值:
//   error: 移动过程中的错误
func Move(srcFile, dstFile string) error {
	if err := Copy(srcFile, dstFile); err != nil {
		return err
	}
	if err := os.Remove(srcFile); err != nil {
		return err
	}
	return nil
}

// MoveWithOverwrite 移动文件，覆盖已存在的目标文件
// 参数:
//   srcFile: 源文件路径
//   dstFile: 目标文件路径
// 返回值:
//   error: 移动过程中的错误
func MoveWithOverwrite(srcFile, dstFile string) error {
	if err := CopyWithOverwrite(srcFile, dstFile); err != nil {
		return err
	}
	if err := os.Remove(srcFile); err != nil {
		return err
	}
	return nil
}

// ReadAllBytes 读取文件的所有字节
// 参数:
//   path: 文件路径
// 返回值:
//   []byte: 文件内容
//   error: 读取过程中的错误
func ReadAllBytes(path string) ([]byte, error) {
	return readFile(path)
}

// ReadAllLines 读取文件的所有行
// 参数:
//   path: 文件路径
// 返回值:
//   []string: 文件的所有行
//   error: 读取过程中的错误
func ReadAllLines(path string) ([]string, error) {
	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}
	newContents := strings.TrimSuffix(string(contents), nextLine)
	return strings.Split(newContents, nextLine), nil
}

// ReadAllText 读取文件的所有文本
// 参数:
//   path: 文件路径
// 返回值:
//   string: 文件内容
//   error: 读取过程中的错误
func ReadAllText(path string) (string, error) {
	contents, err := readFile(path)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

// AppendAllLines 追加多行文本到文件
// 参数:
//   path: 文件路径
//   contents: 要追加的行
// 返回值:
//   error: 追加过程中的错误
func AppendAllLines(path string, contents []string) error {
	str := strings.Join(contents, nextLine)
	str = str + nextLine
	return appendFile(path, []byte(str))
}

// AppendAllText 追加文本到文件
// 参数:
//   path: 文件路径
//   str: 要追加的文本
// 返回值:
//   error: 追加过程中的错误
func AppendAllText(path string, str string) error {
	return appendFile(path, []byte(str))
}

// WriteAllBytes 写入字节到文件
// 参数:
//   path: 文件路径
//   bytes: 要写入的字节
// 返回值:
//   error: 写入过程中的错误
func WriteAllBytes(path string, bytes []byte) error {
	return writeFile(path, bytes)
}

// WriteAllLines 写入多行文本到文件
// 参数:
//   path: 文件路径
//   contents: 要写入的行
// 返回值:
//   error: 写入过程中的错误
func WriteAllLines(path string, contents []string) error {
	str := strings.Join(contents, nextLine)
	str = str + nextLine
	return writeFile(path, []byte(str))
}

// WriteAllText 写入文本到文件
// 参数:
//   path: 文件路径
//   content: 要写入的文本
// 返回值:
//   error: 写入过程中的错误
func WriteAllText(path string, content string) error {
	return writeFile(path, []byte(content))
}

// Exists 检查文件是否存在
// 参数:
//   path: 文件路径
// 返回值:
//   bool: 如果文件存在且不是目录，返回true；否则返回false
func Exists(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		if !info.IsDir() {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// ChangeModTime 修改文件的最后修改时间
// 参数:
//   path: 文件路径
//   t: 新的修改时间
// 返回值:
//   error: 修改过程中的错误
func ChangeModTime(path string, t time.Time) error {
	return os.Chtimes(path, time.Now(), t)
}
