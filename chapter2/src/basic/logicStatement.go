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
	fmt.Print("switch语句:\n")
	var day int8 = 5
	fmt.Println("星期",day)
	switchPrint(day)
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	fmt.Print("for语句:\n")
	forPrint()
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
switch语句 查询天气预报
 */
func switchPrint(day int8)  {
	switch day {
	case 1:
		fmt.Println("周一阴转多云！")
	case 2:
		fmt.Println("周二天气晴朗！")
	case 3:
		fmt.Println("周三小雪！")
	case 4:
		fmt.Println("周四晴转多云！")
	case 5:
		fmt.Println("周五多云！")
	case 6:
		fmt.Println("周六晴！")
	case 7:
		fmt.Println("周日晴！")
	default:
		fmt.Println("请确认一周之内！")
	}
}

/*
for  求1-10连加的和
*/
func forPrint() {
	var i, sum int
	for i = 1;i <= 10;i++ {
        sum = sum + i
	}
	fmt.Printf("1-10连加的和%d" , sum)
}
