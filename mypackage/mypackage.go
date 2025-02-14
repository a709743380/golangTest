package mypackage

import "fmt"

func init() {
	fmt.Println("mypackage init")
}

// SomeFunction 只是一个示例函数，用于调用
func SomeFunction() {
	fmt.Println("SomeFunction in mypackage is called!")
}
