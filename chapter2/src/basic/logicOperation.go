package basic

import "fmt"

/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2020-12-18
 * Description: 逻辑运算的学习
 */

func TestLogicOperation() {
	fmt.Print("进行逻辑运算:\n")
	arithmeticOperation()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	fmt.Print("进行关系运算:\n")
	relationOperation()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	fmt.Print("进行逻辑运算:\n")
	logicOperation()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	fmt.Print("进行位运算:\n")
	bitOperation()
	fmt.Print("\n")
	fmt.Print("-----------------------")

}

/*
测试算术运算符
*/
func arithmeticOperation() {
	var a, b int = 1, 2
	var c int
	c = a + b
	fmt.Print("a+b =", c)
	fmt.Print("\n")
	fmt.Printf("a+b =%d", c)
}

/*
关系运算符
*/
func relationOperation() {
	var a, b int = 1, 2
	if (a > b) {
		fmt.Print("a 大于 b")
	}else{
		fmt.Print("a 小于等于 b")
	}
}

/*
逻辑运算符
 */
func logicOperation() {
	var a, b bool = true, false
	if (a && b) {
		fmt.Print("条件为true")
	}else{
		fmt.Print("条件为false")
	}
}

/*
位运算符
 */
func bitOperation() {
	var a int = 1
	var b int
	b = a <<2
	fmt.Print("a <<2  ", b)
}