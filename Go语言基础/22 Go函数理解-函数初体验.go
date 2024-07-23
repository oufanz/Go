package main

// import (
// 	"fmt"
// 	"strings"
// )

// // func 函数名称(接收的参数 参数的类型, 参数2 类型2)
// // func 函数名称(参数1，参数2，参数3 类型) (返回值1， 返回值2 类型) {代码块/函数体}
// // func 函数名称(参数1，参数2，参数3 类型) {代码块/函数体}
// // func 函数名称(参数1，参数2，参数3 类型)  类型 {代码块/函数体}

// // 定义一个函数，用来判断两个数据的大小，并且返回最大值
// func max(i1, i2 int) int {
// 	if i1 > i2 {
// 		return i1
// 	} else {
// 		return i2
// 	}
// }

// // 求和的函数
// func qiuhe(i1, i2 int) (sum int) {
// 	sum = i1 + i2
// 	return
// }

// func paixu(i1, i2 int) (min, max int) {
// 	if i1 > i2 {
// 		max = i1
// 		min = i2
// 	} else {
// 		max = i2
// 		min = i1
// 	}
// 	return
// }

// // 接收不定长度参数的函数
// func hasName(s ...string) string {
// 	// 接收到的参数是个切片
// 	fmt.Println("接收不定长度的函数:", s)
// 	m := strings.Join(s, "-")
// 	return m
// }

// func main() {
// 	// 可以将重复的代码或者相同的逻辑进行模块化，把大问题转换为小问题
// 	// 逻辑上属于能够完成特定功能的代码块，函数可以接收参数并处理参数，最后返回结果
// 	// 函数不仅可以降低代码的重复率，还可以提高代码的可读性和开发效率
// 	res := max(727585, 266)
// 	fmt.Println("最大值:", res)
// 	res = qiuhe(7275, 85266)
// 	fmt.Println("求和:", res)
// 	min, max := paixu(43, 34)
// 	fmt.Println("从小到大排序:", min, max)

// 	m := hasName("du", "kuan", "is")
// 	fmt.Println("拼接后的字符串:", m)
// }
