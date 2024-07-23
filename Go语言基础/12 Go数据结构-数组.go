package main

// import "fmt"

// func main() {
// 	// 数组：一组具有相同类型并且长度固定的一个数据集合
// 	// 使用场景：班级，三位老师，
// 	// var 数组名 = [数组长度]数组的类型{数组的数据}
// 	var teacherNameArray = [3]string{"杜宽", "Dot", "嘚嘚"}
// 	fmt.Println(teacherNameArray)
// 	teacherAgeArray := [3]int{18, 19, 20}
// 	fmt.Println(teacherAgeArray)
// 	fmt.Println("第一位老师的名字是:", teacherNameArray[0])
// 	fmt.Println("第一位老师的年龄是:", teacherAgeArray[0])
// 	// 获取单个元素的格式：数组名称[下标]，下标是从0开始的
// 	// 修改数据：数组名称[下标] = xxx
// 	teacherNameArray[2] = "dotbalo"
// 	fmt.Println("修改后的数据：", teacherNameArray)
// 	// teacherNameArray[3] = "数据4" // 数组的长度是不能修改的
// 	fmt.Println("数组的长度是:", len(teacherNameArray))
// 	for i := 0; i < len(teacherNameArray); i++ {
// 		fmt.Printf("第%d个数据为: %s\n", i+1, teacherNameArray[i])
// 	}
// 	// range --> 推荐使用的
// 	for k, v := range teacherAgeArray {
// 		fmt.Printf("第%d位老师的年龄是: %d\n", k+1, v)
// 	}
// 	// 自动推断长度
// 	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println("array3的长度是：", len(array3))
// }
