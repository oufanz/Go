package main

// import (
// 	"fmt"
// 	"math"
// 	"reflect"
// )

// func main() {
// 	// 数值类型：int int8 int16 int32 int64 uint
// 	// int: 正负数  uint：不带符号的数字 // int
// 	defaultIntType := 1
// 	fmt.Println("默认的数值类型是:", reflect.TypeOf(defaultIntType))
// 	// int和操作系统是有关系的
// 	// 64位的，int64  32位的 int32
// 	var int64Num int64 = 1
// 	fmt.Println("int64Num的数值类型是:", reflect.TypeOf(int64Num))
// 	var uintNum uint = 1
// 	fmt.Println("uintNum的数值类型是:", reflect.TypeOf(uintNum))
// 	fmt.Println("int的取值范围:", math.MinInt, math.MaxInt)
// 	fmt.Println("uint的取值范围:", uint(math.MaxUint))
// 	fmt.Println(18446744073709551615 > 9223372036854775807)
// 	// float float32和float64
// 	var floatNum1 float64 = 3.14
// 	var floatNum2 float32 = 3.15
// 	// floatSum := floatNum1 + floatNum2
// 	fmt.Println(floatNum1, floatNum2)
// }
