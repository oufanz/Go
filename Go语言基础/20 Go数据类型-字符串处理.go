package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// 字符串的定义
// 	// ``  ""
// 	// "\t \n"  ``
// 	s := "\t\txxx\n"
// 	fmt.Println("双引号字符串:", s)
// 	s2 := `\t\txxx\n`
// 	fmt.Println("反引号字符串:", s2)
// 	s3 := `
// 	我是杜宽
// 	我主要教授的课程是: 云原生系列，k8s，go，python
// 	`
// 	fmt.Println("多行字符串:", s3)
// 	// 字符长度的计算
// 	s4 := "dukuan"
// 	s5 := "杜宽" //中文占用3个字符
// 	s4Length := len(s4)
// 	s5Length := len(s5)
// 	fmt.Println(s4Length, s5Length)
// 	// 字符串的截取  , 一般
// 	// s6 := s4[2]
// 	fmt.Println("前两位:", s4[:2])
// 	s7 := "dukuan"
// 	// 大小写转换, 一般
// 	fmt.Println("转成大写：", strings.ToUpper(s7))
// 	fmt.Println("首字母大写", strings.Title(s7))
// 	s8 := "dUKUan"
// 	fmt.Println("转成小写：", strings.ToLower(s8))
// 	// 字符串是否包含某个元素 ，还行
// 	fmt.Println("查看字符串是否包含uk这个元素:", strings.Contains(s7, "uk"))
// 	fmt.Println("查看字符串是否包含任意一个字符:", strings.ContainsAny(s7, "uw"))
// 	// 忽略大小写进行比较
// 	fmt.Println("忽略大小写比较:", strings.EqualFold(s7, s8))
// 	// 判断字符串中某个元素有多个个
// 	s9 := "dukuan and dot is me, my age is 18"
// 	fmt.Println("u在字符串中出现了:", strings.Count(s9, "u"))
// 	// 字符串拆分 //很多
// 	s9Split := strings.Split(s9, ",")
// 	fmt.Println("使用逗号拆分字符串:", s9Split)
// 	fmt.Println("拆分后的切片的第一个数据:", s9Split[0])
// 	s9SplitAfter := strings.SplitAfter(s9, ",")
// 	fmt.Println("使用逗号拆分字符串,并且保留逗号:", s9SplitAfter)
// 	// 字符串拼接 //很多
// 	slice1 := []string{"dot", "balo", "du", "kuan"}
// 	fmt.Println("拼接字符串:", strings.Join(slice1, " "))
// 	// 是否是某个元素开头的 ,还行
// 	s10 := "我是一个中国人，我非常热爱中国"
// 	fmt.Println("字符串是以 ‘我’ 开头的:", strings.HasPrefix(s10, "我"))
// 	// 此处的视频出现错误，应该是HasSuffix
// 	fmt.Println("字符串是以 爱 结尾的:", strings.HasSuffix(s10, "爱"))
// 	// 重复字符串
// 	fmt.Println("打印五个我:", strings.Repeat("我", 5))
// 	// 字符串替换 ,还行
// 	s11 := "dsad3afd3rq3adawdwarq3a"
// 	fmt.Println("把3替换为dukuan:", strings.ReplaceAll(s11, "3", "dukuan"))
// 	fmt.Println("把3替换为杜宽:", strings.Replace(s11, "3", "杜宽", 1)) // 0替换所有
// 	// trim 字符串修剪 ,很多
// 	s12 := "  dukuan  ,"
// 	fmt.Println("去掉字符串的前后空格:", strings.Trim(s12, ","))
// }
