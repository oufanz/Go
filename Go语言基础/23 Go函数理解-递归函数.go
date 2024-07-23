package main

// import "fmt"

// // 实现数字阶乘
// //
// func factorial(n int) (result int) {
// 	// 5 5 * 4 * 3
// 	// 5 * factorial(4)
// 	// 4 * factorial(3)
// 	// 3 * factorial(2)
// 	// 2 * factorial(1)
// 	// 1 * factorial(0)
// 	// 递归函数的原理：
// 	// 在递归时，相当于创建了一个管道，把每次的调用都扔进到管理内
// 	// 先进后出，先进来的后出去，后进来的先出去。
// 	if n > 0 {
// 		result = n * factorial(n-1)
// 		fmt.Println("当前的数值是:", n)
// 		fmt.Println("当前的计算结果是:", result)
// 		return
// 	} else {
// 		return 1
// 	}
// }

// func main() {
// 	// 递归函数：函数自己调用自己，直到达到条件后才能结束返回结果
// 	//	n! = n*(n-1)*(n-2)***1
// 	// 	5! = 5 * 4 * 3 * 2 * 1
// 	i := 5
// 	res := factorial(i)
// 	fmt.Printf("%d的阶乘结果是:%d", i, res)
// }
