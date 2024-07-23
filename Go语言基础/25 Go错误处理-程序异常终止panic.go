package main

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// // 实现数据库的链接
// func connectDatabase(address string, port int) (string, error) {
// 	// 如果address和port为空
// 	if address == "" || port == 0 {
// 		return "", errors.New("无法链接数据库")
// 	} else {
// 		return "数据库链接成功", nil
// 	}
// }

// func main() {
// 	// panic：可以在异常的时候让程序终止执行，退出程序。或者是程序所强依赖的基础组件不可用
// 	// 此时程序已经无法继续正常工作，此时可以使用panic抛出异常，并且把程序退出
// 	s, err := connectDatabase("", 0)
// 	for {
// 		time.Sleep(5 * time.Second)
// 		// 模式启动程序
// 		if err != nil {
// 			// 说明无法链接数据库
// 			fmt.Println(err)
// 			panic(err) // 就会退出程序
// 		} else {
// 			// 链接成功
// 			fmt.Println(s)
// 			// 正常启动程序
// 		}
// 	}
// 	// 切记：panic不能滥用，不能到处使用panic
// }
