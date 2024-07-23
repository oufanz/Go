package main

// import "fmt"

// func main() {
// 	// 	var teacherNameArray = [3]string{"杜宽", "Dot", "嘚嘚"}
// 	// 	teacherAgeArray := [3]int{18, 19, 20}
// 	// object
// 	// dukuan: 18
// 	// dot: 20
// 	// xiaoming: 21
// 	// var :=
// 	// teacherAge := make(map[keyType]valueType)
// 	teacherAge := make(map[string]int)
// 	fmt.Println("对象的初始化值：", teacherAge)
// 	teacherAge["dukuan"] = 18
// 	teacherAge["dot"] = 20
// 	teacherAge["xiaoming"] = 21
// 	teacherAge["dukuan"] = 21 // 后面的值会覆盖前面的值
// 	fmt.Println("赋值后的值", teacherAge)

// 	// 在声明变量的时候直接进行赋值操作
// 	teacherAge2 := map[string]int{
// 		"d1": 2,
// 		"d2": 3,
// 	}
// 	fmt.Println("teacherAge2:", teacherAge2)
// 	// 使用var赋值
// 	var teacherAddress map[string]string
// 	teacherAddress = make(map[string]string) // 非常非常非常容易犯错的地方
// 	teacherAddress["dukuan"] = "北京"
// 	teacherAddress["dot"] = "加利福尼亚"
// 	fmt.Println("teacherAddress:", teacherAddress)
// 	// s1[0] , teacherAddress["dukuan"]
// 	fmt.Println("查询dukuan老师的地址:", teacherAddress["dukuan"])
// 	searchName := "dot"
// 	fmt.Printf("查询老师:%s的地址:%s\n", searchName, teacherAddress[searchName])
// 	// 查询值 for
// 	for k, v := range teacherAge {
// 		fmt.Printf("老师: %s, 年龄: %d\n", k, v)
// 	}
// 	fmt.Println("取一个不存在的值:", teacherAddress["eeeeee"])
// 	address, ok := teacherAddress["dukuan"]
// 	if ok {
// 		fmt.Println("能查看到dukuan:", address)
// 	} else {
// 		fmt.Println("不能查看到dukuan")
// 	}
// 	// 修改值
// 	teacherAddress["dukuan"] = "上海"
// 	fmt.Println("修改后的值:", teacherAddress["dukuan"])
// 	// 删除值
// 	delete(teacherAddress, "dukuan")
// 	fmt.Println("删除后的map:", teacherAddress)
// }
