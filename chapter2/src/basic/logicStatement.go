package basic

import "fmt"

/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2020-12-18
 * Description: 逻辑控制的学习
 */

func TestLogicStament() {
	fmt.Print("if语句:\n")
	var age int8 = 35
	fmt.Print(age)
	ifPrint(age)
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	fmt.Print("while语句:\n")
	whilePrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
}

/*
if语句 判断年龄结构
*/
func ifPrint(age int8) {
	if (age > 0 && age <= 6) {
		fmt.Print("婴幼儿！")
	} else if (age >= 7 && age <= 12) {
		fmt.Print("少儿！")
	} else if (age >= 13 && age <= 17) {
		fmt.Print("青少年！")
	} else if (age >= 18 && age <= 45) {
		fmt.Print("青年！")
	} else if (age >= 46 && age <= 65) {
		fmt.Print("中年！")
	} else if (age > 65) {
		fmt.Print("老年！")
	}
}

/*
while  求1-10连加的和
*/
func whilePrint() {
	var i, sum int
	for i = 1;i <= 10;i++ {
        sum = sum + i
	}
	fmt.Printf("1-10连加的和%d" , sum)
}
