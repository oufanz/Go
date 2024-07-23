package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	// 在go语言当中定义变量
// 	// 第一种方法: var: 关键字， name: 变量名, string: 变量的类型
// 	// = 表示进行直接赋值
// 	var name string = "杜宽"
// 	fmt.Println("我的名字是:", name)
// 	var name2 string
// 	name2 = "dot"
// 	fmt.Printf("My English name is: %s\n", name2)
// 	// var name3, name4 string = "name3", "name4"
// 	var (
// 		name3 string = "name3"
// 		name4 int
// 	)
// 	println(name3, name4)
// 	// 第二种方法，使用语法糖进行声明变量
// 	b := true
// 	// b = "string"
// 	f := 3.1415926
// 	fmt.Println(b, f)
// 	b2 := false
// 	// 可以进行赋值吗？
// 	b = b2
// 	fmt.Println(b)
// 	// 如何确定一个变量的类型？
// 	// reflect
// 	// name2Type := reflect.TypeOf(name2)
// 	fmt.Println("name2的类型是: ", reflect.TypeOf(name2))
// 	fmt.Println("f的类型是: ", reflect.TypeOf(f))
// 	// 定义一个常量
// 	// 关键字const
// 	const c1, c2, c3 = 1, 2, 3
// 	// const ()
// 	// 可以更改c1的值吗？
// 	// c1 = 5
// 	fmt.Println("常量的值是:", c1, c2, c3)
// }
