package main

// import "fmt"

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

// func printPriceWithSwitch(weather string) {
// 	defaultPrice := 10
// 	switch weather {
// 	case "lightRain":
// 		fmt.Println("下小雨了，雨伞的价格是：", defaultPrice+5)
// 	case "heavyRain":
// 		fmt.Println("下大雨了，雨伞的价格是：", defaultPrice+10)
// 	case "rainStorm":
// 		fmt.Println("下暴雨了，雨伞的价格是：", defaultPrice+20)
// 	case "snowing", "sunny":
// 		// if else
// 		fmt.Println("雨伞的价格是：", defaultPrice)
// 	default:
// 		fmt.Println("我不知道现在是什么天气，所以我不卖了~")
// 	}
// }

// func main() {
// 	// // 小雨
// 	// printPriceWithWeather("lightRain")
// 	// printPriceWithWeather("rainStorm")
// 	// // 晴天了
// 	// printPriceWithWeather("")
// 	printPriceWithSwitch("")
// 	printPriceWithSwitch("snowing")
// }
