package main

// import "fmt"

// func printSliceData(s []string) {
// 	// 使用recover进行异常捕获
// 	defer func() { // 匿名函数
// 		fmt.Println("程序执行失败, 捕获异常")
// 		if err := recover(); err != nil {
// 			// recover是用来捕获panic的报错的
// 			// 尝试恢复，防止程序异常退出
// 			fmt.Println("捕获到了一个错误:", err)
// 			// 发出一个告警
// 			// 记录一条日志
// 			// 返回给前端：说传入的值不对
// 		}
// 	}()
// 	fmt.Println("切片的内容:", s)
// 	// 打印一下切片的第三个值
// 	fmt.Println("切片的第三个值是:", s[2])
// }

// func main() {
// 	// recover 异常的捕获和处理
// 	//
// 	s := []string{"a", "b"}
// 	printSliceData(s)

// }
