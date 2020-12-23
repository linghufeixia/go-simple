package basic

import "fmt"

/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2020-12-18
 * Description: 变量 常量的使用
 */

func TestVariable() {
	fmt.Print("变量的使用:\n")
	variablePrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	fmt.Print("常量的使用:\n")
	constantsPrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
}

//全局变量
var status int = 8

/*
变量可以在三个地方声明：
1 函数内定义的变量称为局部变量
2 函数外定义的变量称为全局变量
3 函数定义中的变量称为形式参数
*/
func variablePrint() {
	// 局部变量
	var a, b, c int
	a = 5
	b = 6
	fmt.Printf("局部变量:%d,%d", a, b)
	fmt.Print("\n")
	fmt.Printf("全局变量:%d", status)
	fmt.Print("\n")
	var status int = 6
	fmt.Printf("和全局变量同名的局部变量:%d", status)
	fmt.Print("\n")
	c = area(a, b)
	fmt.Printf("面积:%d", c)
}

func area(x, y int) int {
	return x * y;
}

/*
常量的使用
*/
func constantsPrint() {
	const Finish int = 100
	fmt.Printf("常量:%d", Finish)
	fmt.Print("\n")
	const(
		Male = 1
		Female = 2
	)
	var sex int32
	sex = Male
	fmt.Println("枚举常量:男性 ", sex)
	const(
		a = iota
		b
		c = 100
		d
		e= "hello"
		f
		g= iota
		h
	)
	fmt.Print("iota可修改的常量:\n")
	fmt.Printf("a:%d\n",a)
	fmt.Printf("b:%d\n",b)
	fmt.Printf("c:%d\n",c)
	fmt.Printf("d:%d\n",d)
	fmt.Printf("e:%s\n",e)
	fmt.Printf("f:%s\n",f)
	fmt.Printf("g:%d\n",g)
	fmt.Printf("h:%d",h)
}
