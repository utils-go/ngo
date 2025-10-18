package directory

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
func Delete(path string) error {
	return os.RemoveAll(path)
}
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
