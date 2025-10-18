package console

import "fmt"

func WriteLine(a ...any) {
	fmt.Println(a...)
}
func Write(a ...any) {
	fmt.Print(a...)
}
func ReadLine() string {
	var str string
	_, err := fmt.Scanln(&str)
	if err != nil {
		return str
	}
	return ""
}
