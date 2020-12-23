package basic

import "fmt"

/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2020-12-18
 * Description: 基本数据类型的学习
 * 派生类型:
(a) 指针类型（Pointer）
*/

func TestBasicType() {
	fmt.Print("布尔型的使用:\n")
	boolPrinf()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	fmt.Print("数字类型的使用:\n")
	numberPrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	fmt.Print("字符串类型的使用:\n")
	stringPrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
}

/*
布尔型的值输出 只可以是常量 true 或者 false
*/
func boolPrinf() {
	var b bool
	b = true
	fmt.Print("布尔bool:")
	fmt.Print(b)
}

/*
数字类型输出 整型 int 和浮点型 float32、float64，
*/
func numberPrint() {
	var i int16
	i = 20
	fmt.Print("int:")
	fmt.Print(i)
	fmt.Print("\n")
	var j float32
	j = 50
	fmt.Print("float32 ",j)
}

/*
字符串类型的输出
*/
func stringPrint() {
	var str1 string = "nice to meet you!"
	// := 是一个声明语句
	str2 := "nice to meet you too!"
	fmt.Print(str1)
	fmt.Print("\n")
	fmt.Printf("%v %v", str1, str2)
}




