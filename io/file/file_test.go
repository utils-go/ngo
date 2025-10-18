package file

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

// ok
func TestAppendAllText(t *testing.T) {
	dir, _ := os.Getwd()
	path := "./test.txt"
	path = filepath.Join(dir, path)
	t.Log(path)
	err := AppendAllText(path, "test")
	if err != nil {
		t.Error(err)
	}
	t.Log("test finish")
}
func TestAppendAllLines(t *testing.T) {
	dir, _ := os.Getwd()
	path := "./test.txt"
	path = filepath.Join(dir, path)
	t.Log(path)
	contents := []string{"1", "2", "3", "4", "5"}
	err := AppendAllLines(path, contents)
	if err != nil {
		t.Error(err)
	}
	t.Log("test finish")
}

// ok
func TestWriteAllText(t *testing.T) {
	dir, _ := os.Getwd()
	path := "./test.txt"
	path = filepath.Join(dir, path)
	t.Log(path)
	err := WriteAllText(path, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		t.Error(err)
	}
	t.Log("test finish")
}
func TestWriteAllLines(t *testing.T) {
	dir, _ := os.Getwd()
	path := "./test.txt"
	path = filepath.Join(dir, path)
	t.Log(path)
	contents := []string{"hello", "world", "I", "am", "go"}
	err := WriteAllLines(path, contents)
	if err != nil {
		t.Error(err)
	}
	t.Log("test finish")
}

func TestCopy(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test.txt"
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test1.txt"
	dstPath = filepath.Join(dir, dstPath)
	
	// 清理目标文件
	os.Remove(dstPath)
	
	err := Copy(srcPath, dstPath)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Log("test finish")
}
func TestCopyError(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test111.txt" //not exist file
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test1.txt"
	dstPath = filepath.Join(dir, dstPath)
	err := Copy(srcPath, dstPath)
	if err == nil {
		t.Error("copy a file that not exist success")
	} else {
		t.Logf("test ok with error:%v", err)
	}

}

func TestCopyWithOverwrite(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test.txt"
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test1.txt"
	dstPath = filepath.Join(dir, dstPath)
	err := CopyWithOverwrite(srcPath, dstPath)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Log("test finish")
}
func TestMove(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test.txt"
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test_move.txt"
	dstPath = filepath.Join(dir, dstPath)
	
	// 清理目标文件
	os.Remove(dstPath)
	// 确保源文件存在
	WriteAllText(srcPath, "test content for move")
	
	err := Move(srcPath, dstPath)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Log("test finish")
}
func TestMoveError(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test1.txt"
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test_move.txt"
	dstPath = filepath.Join(dir, dstPath)
	err := Move(srcPath, dstPath)
	if err == nil {
		t.Errorf("%v", err)
		return
	}
	t.Log("test finish")
}
func TestMoveWithOverwrite(t *testing.T) {
	dir, _ := os.Getwd()
	srcPath := "./test.txt"
	srcPath = filepath.Join(dir, srcPath)
	dstPath := "./test_move.txt"
	dstPath = filepath.Join(dir, dstPath)
	
	// 确保源文件存在
	WriteAllText(srcPath, "test content for move with overwrite")
	
	err := MoveWithOverwrite(srcPath, dstPath)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Log("test finish")
}
func TestReadAllBytes(t *testing.T) {
	path := "./xxx.txt"
	data := []byte{1, 2, 3, 4}
	err := WriteAllBytes(path, data)
	t.Log("write finish")
	if err != nil {
		t.Error(err)
		return
	}
	dataNew, err := ReadAllBytes(path)
	t.Log("read finish")
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(dataNew, dataNew) {
		t.Errorf("new bytes :%v not equals to old bytes:%v", dataNew, data)
		return
	}
}
func TestReadAllLines(t *testing.T) {
	path := "./xxx.txt"
	data := []string{"1", "2", "3", "4"}
	err := WriteAllLines(path, data)
	t.Log("write finish")
	if err != nil {
		t.Error(err)
		return
	}
	dataNew, err := ReadAllLines(path)
	t.Log("read finish")
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(data, dataNew) {
		t.Errorf("new bytes :%v not equals to old bytes:%v", dataNew, data)
		return
	}
}
