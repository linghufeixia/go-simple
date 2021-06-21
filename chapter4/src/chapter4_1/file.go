/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2020-12-24
 * Description: 文件的使用
 */
package chapter4_1

import (
	"fmt"
	"os"
	"syscall"
)

//创建文件
func TestCreateFile() {
	fmt.Print("创建文件:\n")
	_,err:=os.Create("user.txt")
	if err!=nil{
		fmt.Println(err)
	}
}

//写入文件
func TestWriteFile()  {
	file,err:= os.OpenFile("user.txt", os.O_WRONLY,2)
	defer file.Close()
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("写入文件:")
	file.WriteString("{id:1;")
	file.WriteString("userName:令狐冲;")
	file.WriteString("userPwd:123456;")
	file.WriteString("UserPhone:19383928933;")
	file.WriteString("RegTime:2021-6-17}")
}

//读取文件
func TestReadFile()  {
	file,err:= os.OpenFile("user.txt", syscall.O_RDONLY,4)
	defer file.Close()
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("读取文件的内容:")
	var b []byte = make([]byte, 1024)
	for{
		n, _ := file.Read(b) //读文件
		if n == 0 {          //0表示已经到文件结尾
			break
		}
		fmt.Println(string(b))
	}

}

//删除文件
func TestDeleteFile()  {
	fmt.Print("删除文件:\n")
	err:=os.Remove("user.txt")
	if err!=nil{
		fmt.Println(err)
	}
}





