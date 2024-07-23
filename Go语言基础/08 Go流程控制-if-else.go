package main

// import "fmt"

// // 定义价格的函数
// func printPrice(weather string) {
// 	defaultPrice := 10
// 	// if 表达式 {代码块} else {代码块}
// 	// if 表达式 {代码块} else if  {代码块} else {}
// 	if weather == "sunny" {
// 		// 成立的话就执行下面的代码
// 		fmt.Println("今天是晴天，雨伞的价格是：", defaultPrice)
// 	} else {
// 		// 如果不成立执行else里面的代码，当然else可以没有
// 		fmt.Println("今天不是晴天，雨伞的价格是：", defaultPrice+10)
// 	}
// }

// func printPriceWithWeather(weather string) {
// 	defaultPrice := 10
// 	// 小雨涨5块
// 	if weather == "lightRain" {
// 		// 成立的话就执行下面的代码
// 		fmt.Println("下小雨了，雨伞的价格是：", defaultPrice+5)
// 	} else if weather == "heavyRain" {
// 		// 如果不成立执行else里面的代码，当然else可以没有
// 		fmt.Println("下大雨了，雨伞的价格是：", defaultPrice+10)
// 	} else if weather == "rainStorm" {
// 		fmt.Println("下暴雨了，雨伞的价格是：", defaultPrice+20)
// 	} else {
// 		fmt.Println("雨伞的价格是：", defaultPrice)
// 	}
// }

// func main() {
// 	// if else
// 	// 雨伞：10  下雨：20
// 	weather := "rain"
// 	printPrice(weather)

// 	// 晴天了
// 	weather = "sunny"
// 	printPrice(weather)

// 	// 小雨
// 	printPriceWithWeather("lightRain")
// 	printPriceWithWeather("rainStorm")
// 	// 晴天了
// 	printPriceWithWeather("")
// }
