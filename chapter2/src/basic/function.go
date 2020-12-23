package basic

import "fmt"

/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2020-12-18
 * Description: 函数的学习
 */

func TestFunction() {
	callFuctionPrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	functionReturnValuePrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	functionVariablePrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
}

/*
调用函数的输出
 */
func callFuctionPrint()  {
	fmt.Print("调用函数的使用:\n")
	var x, y float32
	x = 5.6
	y = 7
	fmt.Println("长方形的长：", y)
	fmt.Println("长方形的宽：", x)
	fmt.Println("面积：",getArea(x,y))
}

/*
函数返回值的输出
 */
func functionReturnValuePrint()  {
	fmt.Print("函数多个返回值的使用:\n")
	var msg string
	var errorCode int
	msg, errorCode = getResponseMessage()
	fmt.Println("错误信息:" , msg)
	fmt.Print("错误code:" , errorCode)
}


/*
函数变量的使用
 */
func functionVariablePrint()  {
	fmt.Print("函数变量的使用:\n")
	var function func()
	function =  sayHello
	function()
}

/*
求长方形的面积
 */
func getArea(a, b float32)  float32  {
	var area float32
	area = a * b
	return area
}

/*
获取返回的消息
*/
func getResponseMessage() (message string, code int) {
	message = "用户名或密码错误!"
	code = 1
	return message, code
}

func sayHello()  {
	fmt.Print("hello! 我是通过函数变量调用的")
}
