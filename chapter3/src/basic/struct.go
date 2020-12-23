package basic

import "fmt"

/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2020-12-21
 * Description: 结构体类型的使用 go中没有类，没有继承
 */

func TestStruct() {
	structPrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	modelStructPrint();
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	nestStructPrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
}

type User struct {
	id       float32 // id
	name     string  // 姓名
	password string  // 密码
	sex      int     // 性别
}

//嵌套struct 实现继承
type Student struct {
	user   User
	school string
	grade  string
}

/*
结构体的使用输出
*/
func structPrint() {
	fmt.Print("结构体的使用:\n")
	// 创建匿名结构体
	fmt.Print(User{1, "令狐冲", "123456", 1})
	fmt.Print("\n")
	// 声明结构体
	var user User
	// 修改结构体成员
	user.id = 2
	user.name = "任盈盈"
	userPrint(user)
}

/*
仿造结构体构造函数
 */
func modelStructPrint() {
	fmt.Print("仿造结构体的构造函数:\n")
	var user *User
	var id float32 = 12
	user = NewUserById(id, "令狐冲")
	fmt.Print(user)
}

/*
嵌套结构体的实现
 */
func nestStructPrint()  {
	fmt.Print("嵌套结构体实现继承的特性:\n")
	var user User
	user.id = 2
	user.name = "任盈盈"
	var student Student
	student.user = user
	student.school="中华女子武术学校"
	fmt.Println("姓名：",student.user.name)
	fmt.Print("就读学校：",student.school)
}

/*
模拟结构体的构造函数
*/
func NewUserById(id float32, name string) *User {
	return &User{
		id:       id,
		name:     name,
		password: "",
		sex:      0,
	}
}

func userPrint(user User) {
	fmt.Printf("ID : %f\n", user.id)
	fmt.Printf("姓名 : %s\n", user.name)
	fmt.Printf("密码 : %s\n", user.password)
	fmt.Printf("性别 : %d\n", user.sex)
}
