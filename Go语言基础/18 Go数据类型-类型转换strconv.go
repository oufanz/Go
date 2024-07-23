package main

// import (
// 	"fmt"
// 	"math"
// 	"reflect"
// 	"strconv"
// )

// func area(r float64) float64 {
// 	s := math.Pi * r * r
// 	return s
// }

// func main() {
// 	r := 5
// 	s := area(float64(r))
// 	fmt.Println("面积是:", s)

// 	i := 7275
// 	// 3242342333423432 => "3242342333423432"
// 	// strconv
// 	sn := strconv.Itoa(i)
// 	fmt.Println(sn)
// 	fmt.Println("类型:", reflect.TypeOf(sn))
// 	s1 := "50"
// 	// "20a"
// 	i1, _ := strconv.Atoi(s1)
// 	fmt.Println("转换后的数值:", i1)
// 	s2 := "200"
// 	i2, err := strconv.Atoi(s2)
// 	if err != nil {
// 		// 产生了错误
// 		fmt.Println("转换失败，当前值不能转换为数值：", s2)
// 	} else {
// 		// 转换成功
// 		fmt.Println("转换成功：", i2)
// 	}
// }
