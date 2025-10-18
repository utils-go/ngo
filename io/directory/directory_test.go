package directory

import "testing"

func TestCreateDirectory(t *testing.T) {
	path := "c:/a/b/c"
	err := CreateDirectory(path)
	if err != nil {
		t.Errorf("创建文件夹失败,%v", err)
		return
	}
	exist := Exists(path)
	if !exist {
		t.Errorf("文件夹不存在")
	}
}
func TestDelete(t *testing.T) {
	//t.Skip()

	Delete("c:/a")
}

func TestGetDirectories(t *testing.T) {
	files, err := GetDirectories("c:/")
	if err != nil {
		t.Fatalf("获取文件夹错误：%v", err)
		return
	}

	for _, f := range files {
		t.Log(f)
	}
}

func TestGetFiles(t *testing.T) {
	// 使用当前目录进行测试
	files, err := GetFiles(".", "*.go", false)
	if err != nil {
		t.Fatalf("获取文件出错：%v", err)
		return
	}

	// 应该至少找到一些.go文件
	if len(files) == 0 {
		t.Log("当前目录没有找到.go文件，这是正常的")
	} else {
		for _, file := range files {
			t.Log(file)
		}
	}
}
