package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func numOperations(a, b int) {
// 	fmt.Printf("%d + %d = %d\n", a, b, a+b)
// 	fmt.Printf("%d - %d = %d\n", a, b, a-b)
// 	fmt.Printf("%d * %d = %d\n", a, b, a*b)
// 	fmt.Printf("%d / %d = %f\n", a, b, float64(a)/float64(b))
// 	// 1. 转换类型
// 	// 2. 除数不能为0 if-else
// 	fmt.Printf("%d 取余 %d = %d\n", a, b, a%b)
// }

// // 字符串运算
// func stringOperation(a, b string) {
// 	fmt.Printf("a和b拼接后: %s\n", a+b) // 12 字符串
// 	ab := a + b
// 	fmt.Printf("ab: %s, 类型是:%s", ab, reflect.TypeOf(ab))
// }

// // fmt.Sprintf
// func stringSprintf(firstName, secondName string) {
// 	// fullName := secondName + " " + firstName
// 	fullName := fmt.Sprintf("%s %s", secondName, firstName)
// 	fmt.Println("你的全名是:", fullName)
// }

// func main() {
// 	// 数值的运算
// 	// i1 := 1
// 	// i2 := 2
// 	// fmt.Println("i1 + i2 = ", i1+i2)
// 	numOperations(1, 2)
// 	// numOperations(11111215, 22878454)
// 	name1 := "杜宽"
// 	name2 := "Dot"
// 	stringOperation(name1, name2)
// 	stringOperation("1", "2") // =3?  12?
// 	// numOperations(name1, name2)
// 	stringSprintf("宽", "杜") // 杜宽宽
// 	// strings.Join
// 	// 自增和自减
// 	p1 := 8
// 	p1++
// 	fmt.Println("自增后的值:", p1)
// 	p1--
// 	p1--
// 	fmt.Println("自减后的值:", p1)
// }
