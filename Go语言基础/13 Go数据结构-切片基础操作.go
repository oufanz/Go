package main

// import "fmt"

// func main() {
// 	// 切片的长度是不固定的，并且可以扩容
// 	// var 切片名称 []切片类型  slice
// 	var s1 []int
// 	fmt.Println("最初的切片数据：", s1)
// 	// 默认的两个属性，一个是切片的长度，表示这个切片中有多少个元素
// 	// 容量：表示这个切片可以放入多少个元素
// 	fmt.Println("切片的默认长度是：", len(s1))
// 	fmt.Println("切片的默认容量是：", cap(s1))
// 	s1 = append(s1, 7275, 85266)
// 	fmt.Println("长度是：", len(s1))
// 	fmt.Println("容量是：", cap(s1))
// 	fmt.Println("数据：", s1)
// 	// 第二种声明方式，指定长度
// 	s2 := make([]int, 5, 10) // 如果是string类型切片，默认的数据是什么呢？
// 	fmt.Println("切片的默认长度是：", len(s2))
// 	fmt.Println("切片的默认容量是：", cap(s2))
// 	fmt.Println("最初的切片数据：", s2)
// 	s2 = append(s2, 1, 2, 3, 4, 5, 6)
// 	fmt.Println("最初的切片数据：", s2)
// 	// 现在的切片的容量和长度是多少？
// 	// 11？ 如果 容量， 容量*2
// 	fmt.Println("长度是：", len(s2))
// 	fmt.Println("容量是：", cap(s2))
// 	s2[0] = 88
// 	fmt.Println("切片的数据：", s2)
// 	for k, v := range s2 {
// 		fmt.Printf("第%d个数据是: %d\n", k+1, v)
// 	}
// }
