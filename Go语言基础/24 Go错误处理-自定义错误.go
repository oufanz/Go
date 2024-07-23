package main

// import (
// 	"errors"
// 	"fmt"
// )

// // 定义一个函数，用来实现除法
// func division(i1, i2 float64) (res float64, err error) {
// 	fmt.Println("需要计算的数字是:", i1, i2)
// 	if i2 == 0 {
// 		return 0, errors.New("输入的分母不能为0")
// 	} else {
// 		res = i1 / i2
// 		return res, nil
// 	}
// }

// func main() {
// 	// go 错误定义为一种类型，err = ccc
// 	// 像使用其他类型的变量一样，去处理我们的错误。
// 	// go打开一个本地文件。
// 	// 1. 打开这个文件 文件不存在  没有权限
// 	// 2. 写入内容/读取内容
// 	// f, err := ioutil.ReadFile("./text.txt")
// 	// if err != nil {
// 	// 	// 此时readfile报错，出现了问题
// 	// 	fmt.Println("读取文件内容失败:", err.Error())
// 	// } else {
// 	// 	fmt.Println(string(f))
// 	// }
// 	// 自定义err
// 	err := errors.New("这是一个自定义错误")
// 	fmt.Println(err)
// 	err2 := fmt.Errorf("这是一个自定义错误: %s,它是使用fmt生成的", "这是错误内容")
// 	fmt.Println("这是一个使用fmt定义的错误:", err2.Error())
// 	// var
// 	res, err3 := division(2, 0)
// 	if err3 != nil {
// 		fmt.Println("计算错误:", err3.Error())
// 	} else {
// 		fmt.Println("计算结果:", res)
// 	}
// }
