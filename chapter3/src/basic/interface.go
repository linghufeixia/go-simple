package basic

import "fmt"

/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2020-12-21
 * Description: 接口类型的使用
隐式接口
1 接口的方法和实现接口的方法一致；
2 接口中所有的方法均被实现；
3 接口和接口的实现；一个接口可以被多个struct实现，一个struct可以实现多个接口
 */


func TestInterface() {
	fmt.Print("接口的使用:\n")
	var animal Animal
	animal = new(Dog)
	animal.name()
	fmt.Print("\n")
	fmt.Print("-----------------------")
}

// 第1步：定义接口
type Animal interface {
	name()
}

// 第2步：定义struct 用于给接口赋值
type Dog struct {

}

// 第3步：实现接口
func  (dog Dog) name()  {
	fmt.Print("我是一条可爱的小狗!")
}






