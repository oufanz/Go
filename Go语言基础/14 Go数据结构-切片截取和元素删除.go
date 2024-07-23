package main

// import "fmt"

// func main() {
// 	var s = []int{12, 2312, 33, 314, 2345, 346, 7342, 438, 3439}
// 	fmt.Println("最初的数据是:", s)
// 	// s[0] [1] [2] [3] --> [:5]   [x:x]
// 	fmt.Println("切片s前四位数据是:", s[:4])
// 	fmt.Println("切片s前四位数据是:", s[0:4])
// 	fmt.Println("切片3-4位的数据:", s[2:4])
// 	fmt.Println("切片从第5个开始的数据:", s[4:])

// 	// 删除一个元素
// 	// s.remove(1)  s.delete(4)
// 	// fmt.Println("切片从第2个开始的数据:", s[1:])
// 	fmt.Println("现在的数据是:", s)
// 	s = s[1:] // 删除一个元素
// 	fmt.Println("删除第一个元素:", s)
// 	// length := len(s)  0-7 0-6
// 	s = s[:len(s)-1] // [:7]
// 	fmt.Println("删除最后一个元素:", s)
// 	// 删除314  --> 2  0-1 [0:2]  [3:]
// 	s = append(s[0:2], s[3:]...)
// 	fmt.Println("删除314后的数据:", s)
// }
