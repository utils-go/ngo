package path

import (
	"os"
	"path/filepath"
	"strings"
)

// 参考: https://learn.microsoft.com/en-us/dotnet/api/system.io.path?view=netframework-4.7.2

// ChangeExtension 更改路径字符串的扩展名
// 参数:
//
//	path: 路径字符串
//	extension: 新的扩展名
//
// 返回值:
//
//	string: 带有新扩展名的路径字符串
func ChangeExtension(path, extension string) string {
	ext := filepath.Ext(path)
	fullFileNameWithoutExtension := strings.TrimSuffix(path, ext)

	extensionNew := extension
	if !strings.HasPrefix(extensionNew, ".") {
		extensionNew = "." + extensionNew
	}
	return filepath.Join(fullFileNameWithoutExtension, extensionNew)
}

// Combine 合并多个路径字符串
// 参数:
//
//	path: 第一个路径
//	paths: 要合并的其他路径
//
// 返回值:
//
//	string: 合并后的路径
func Combine(path string, paths ...string) string {
	pathList := []string{path}
	pathList = append(pathList, paths...)
	return filepath.Join(pathList...)
}

// GetDirectoryName 获取指定路径字符串的目录信息
// 参数:
//
//	path: 路径字符串
//
// 返回值:
//
//	string: 目录路径
func GetDirectoryName(path string) string {
	return filepath.Dir(path)
}

// GetExtension 获取指定路径字符串的扩展名
// 参数:
//
//	path: 路径字符串
//
// 返回值:
//
//	string: 扩展名（包含点号）
func GetExtension(path string) string {
	return filepath.Ext(path)
}

// GetFileName 获取指定路径字符串的文件名和扩展名
// 参数:
//
//	path: 路径字符串
//
// 返回值:
//
//	string: 文件名和扩展名
func GetFileName(path string) string {
	return filepath.Base(path)
}

// GetFileNameWithoutExtension 获取指定路径字符串的文件名（不含扩展名）
// 参数:
//
//	path: 路径字符串
//
// 返回值:
//
//	string: 不含扩展名的文件名
func GetFileNameWithoutExtension(path string) string {
	fileName := GetFileName(path)
	extension := GetExtension(path)
	return strings.TrimSuffix(fileName, extension)
}

// GetFullPath 获取指定路径的绝对路径
// 参数:
//
//	path: 路径字符串
//
// 返回值:
//
//	string: 绝对路径
//	error: 获取过程中的错误
func GetFullPath(path string) (string, error) {
	return filepath.Abs(path)
}

// HasExtension 确定路径是否包含文件扩展名
// 参数:
//
//	path: 路径字符串
//
// 返回值:
//
//	bool: 如果路径包含扩展名，返回true；否则返回false
func HasExtension(path string) bool {
	return filepath.Ext(path) != ""
}

// Equals 比较两个路径是否相同
// 参数:
//
//	path1: 第一个路径
//	path2: 第二个路径
//
// 返回值:
//
//	bool: 如果路径相同，返回true；否则返回false
func Equals(path1, path2 string) bool {
	newPath1 := strings.ReplaceAll(path1, "\\", "/")
	newPath2 := strings.ReplaceAll(path2, "\\", "/")
	return newPath1 == newPath2
}

// Exists 检查路径是否存在
// 参数:
//
//	path: 路径字符串
//
// 返回值:
//
//	bool: 如果路径存在，返回true；否则返回false
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		return false
	}
}

// GetRelativePath 获取相对于基本路径的相对路径
// 参数:
//
//	basePath: 基本路径
//	path: 目标路径
//
// 返回值:
//
//	string: 相对路径
//	error: 获取过程中的错误
func GetRelativePath(basePath, path string) (string, error) {
	return filepath.Rel(basePath, path)
}

// IsSubPath 检查目标路径是否是基本路径的子路径
// 参数:
//
//	basePath: 基本路径
//	targetPath: 目标路径
//
// 返回值:
//
//	bool: 如果目标路径是基本路径的子路径，返回true；否则返回false
func IsSubPath(basePath, targetPath string) bool {
	rel, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		return false
	}
	// 如果相对路径不以 "../" 开头，说明 targetPath 在 basePath 内
	return !strings.HasPrefix(rel, "..") && rel != ".."
}
