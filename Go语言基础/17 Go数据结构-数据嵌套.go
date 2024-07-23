package main

// import "fmt"

// func main() {
// 	// teacher
// 	// 年龄、姓名、电话
// 	// 数据嵌套
// 	// map，map[string]int map[string]map[string]string
// 	// [1,2,3,4,5]
// 	//
// 	// 1. 宫保鸡丁: 45 糖醋鱼: 80
// 	// 2. 糖醋里脊: 55 回锅肉: 50
// 	// 3. 奶茶: 20  啤酒: 10
// 	// []map[string]int
// 	order1 := map[string]int{
// 		"宫保鸡丁": 45,
// 		"糖醋鱼":  80,
// 	}
// 	order2 := map[string]int{
// 		"糖醋里脊": 55,
// 		"回锅肉":  50,
// 	}
// 	order3 := map[string]int{
// 		"奶茶": 20,
// 		"啤酒": 10,
// 	}
// 	var menu []map[string]int
// 	menu = append(menu, order1, order2, order3)
// 	fmt.Println(menu)
// 	for k, v := range menu {
// 		fmt.Printf("第%d天的菜单是:\n", k+1)
// 		for name, price := range v {
// 			fmt.Printf("\t菜:%s, 价格:%d\n", name, price)
// 		}
// 	}
// 	// 对象嵌套对象？
// 	// map[string]map[string]string
// 	// info = map[string]string
// 	// info["address"] = "地址"
// 	// info["name"] = "姓名"
// 	// teacher["dukuan"] = info // teacher[dukuan][name] = xxx
// 	// teacher["dukuan"]["name"]
// }
