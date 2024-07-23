package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	// s1  s2 := s1
// 	// var s1 []int{1,2,3,4}
// 	// s2 := s1
// 	// 引用类型 值类型
// 	// 深拷贝 和 浅拷贝
// 	s1 := "dk"
// 	s2 := s1
// 	fmt.Println("s1和s2:", s1, s2)
// 	s2 = "dot"
// 	fmt.Println("s1和s2:", s1, s2) // dk dot
// 	slice1 := []int{1, 2, 3, 4, 5}
// 	slice2 := slice1 // 浅拷贝
// 	fmt.Println("slice1:", slice1)
// 	// fmt.Println("slice2:", slice2)
// 	// slice2[1] = 88 // slice2 -- > 1 88 3 4 5
// 	fmt.Println("修改后的slice1:", slice1)
// 	// fmt.Println("修改后的slice2:", slice2)
// 	// var slice3 []int
// 	slice3 := make([]int, len(slice1), cap(slice1))
// 	copy(slice3, slice1)
// 	fmt.Println("slice3:", slice3)
// 	slice3[1] = 88
// 	fmt.Println("修改后的slice1:", slice1)
// 	fmt.Println("修改后的slice3:", slice3)

// 	// 打印某个值的内存地址
// 	fmt.Println("slice1的内存地址:", unsafe.Pointer(&slice1))
// 	fmt.Println("slice3的内存地址:", unsafe.Pointer(&slice3))
// 	fmt.Println("slice1的第一个元素的内存地址:", unsafe.Pointer(&slice1[0]))
// 	fmt.Println("slice2的第一个元素的内存地址:", unsafe.Pointer(&slice2[0]))
// 	fmt.Println("slice3的第一个元素的内存地址:", unsafe.Pointer(&slice3[0]))
// }
