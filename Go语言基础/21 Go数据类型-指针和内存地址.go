package main

// import "fmt"

// func updateString(s string) {
// 	s = "这是一个新值"
// }
// func updateStringWithPointer(s *string) {
// 	*s = "这是一个新值"
// }

// func main() {
// 	var s string
// 	s = "这是一个字符串"
// 	fmt.Println("变量s的内存地址是:", &s)
// 	// sp: 指针
// 	// 内存地址: 通过&符号进行取值
// 	sp := &s
// 	fmt.Println("指针sp:", sp)
// 	var sp2 *string
// 	fmt.Println("指针sp2:", sp2) // 指针未进行赋值的时候为nil
// 	sp2 = &s
// 	fmt.Println("指针sp2:", sp2)
// 	// 通过指针获取内存地址的值
// 	fmt.Println("指针对应内存地址的值：", *sp2)

// 	updateString(s)
// 	fmt.Println("修改后的s:", s)

// 	updateStringWithPointer(&s)
// 	fmt.Println("真正修改后的s:", s)
// }
