package console

import "fmt"

// WriteLine 将指定的数据写入标准输出，后跟当前行终止符
// 参数:
//   a: 要写入的任意类型的数据
func WriteLine(a ...any) {
	fmt.Println(a...)
}

// Write 将指定的数据写入标准输出
// 参数:
//   a: 要写入的任意类型的数据
func Write(a ...any) {
	fmt.Print(a...)
}

// ReadLine 从标准输入读取一行
// 返回值:
//   string: 读取的字符串
func ReadLine() string {
	var str string
	_, err := fmt.Scanln(&str)
	if err != nil {
		return str
	}
	return ""
}
