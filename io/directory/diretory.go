package directory

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// CreateDirectory 创建一个目录，包括任何必要的父目录
// 参数:
//   path: 要创建的目录路径
// 返回值:
//   error: 创建过程中的错误
func CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Delete 删除指定的目录及其所有内容
// 参数:
//   path: 要删除的目录路径
// 返回值:
//   error: 删除过程中的错误
func Delete(path string) error {
	return os.RemoveAll(path)
}

// GetDirectories 获取指定路径下的所有子目录名称
// 参数:
//   path: 要获取子目录的路径
// 返回值:
//   []string: 子目录名称列表
//   error: 操作过程中的错误
func GetDirectories(path string) ([]string, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			result = append(result, dir.Name())
		}
	}
	return result, nil
}

// GetFiles 获取指定路径下的文件列表
// 参数:
//   path: 要搜索的路径
//   searchPattern: 文件搜索模式，如 "*.txt"
//   recursion: 是否递归搜索子目录
// 返回值:
//   []string: 符合条件的文件名列表
//   error: 操作过程中的错误
func GetFiles(path string, searchPattern string, recursion bool) ([]string, error) {
	result := make([]string, 0)
	extension := strings.TrimPrefix(searchPattern, "*")
	if !recursion {
		dirs, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		for _, dir := range dirs {
			if dir.IsDir() {
				continue
			}
			if extension == "" {
				result = append(result, dir.Name())
			} else {
				if filepath.Ext(dir.Name()) == extension {
					result = append(result, dir.Name())
				}
			}
		}
	} else {
		err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d == nil || d.IsDir() {
				return nil
			}

			if extension == "" {
				result = append(result, d.Name())
			} else {
				if extension == filepath.Ext(d.Name()) {
					result = append(result, d.Name())
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// Exists 检查指定的路径是否存在且是一个目录
// 参数:
//   path: 要检查的路径
// 返回值:
//   bool: 如果路径存在且是目录，返回true；否则返回false
func Exists(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
