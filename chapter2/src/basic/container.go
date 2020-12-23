package basic

import "fmt"

/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2020-12-21
 * Description: 容器的学习
(1) 数组类型
(2) 切片类型
(3) Map 类型
*/
func TestContainer() {
	fmt.Print("\n")
	fmt.Print("数组的使用:\n")
	arrayPrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	fmt.Print("切片slice的使用:\n")
	slicePrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
	fmt.Print("\n")
	fmt.Print("map的使用:\n")
	mapPrint()
	fmt.Print("\n")
	fmt.Print("-----------------------")
}

/*
数组的输出
*/
func arrayPrint() {
	// 定义数组
	 peopleArray := [...] string{"令狐冲", "任盈盈", "东方不败"}

	// 通过长度遍历数组
	fmt.Println("通过直接访问数组遍历：")
	var i int
	for i = 0; i < 3; i++ {
		fmt.Println(i, peopleArray[i])
	}

	// 通过range遍历数组
	fmt.Println("通过range关键字遍历：")
	i = 0
	for j, str := range peopleArray {
		fmt.Println(j, str)
	}
	fmt.Println("通过range只输出元素：")
	for _, str := range peopleArray {
		fmt.Println(str)
	}
}

/*
切片slice的使用
*/
func slicePrint() {
	// 定义切片 长度和容量都为初始切片的大小
	 pepoleSlice := [] string {"令狐冲"}

	// make创建切片
	//var pepoleSlice = make([]string,1,10)

	var len int = len(pepoleSlice)
	fmt.Println("切片slice的长度：", len)
	var capacity int = cap(pepoleSlice)
	fmt.Println("切片slice的容量：", capacity)

	pepoleSlice = append(pepoleSlice, "任盈盈")
	fmt.Print("切片slice增加1个元素：")
	fmt.Println(pepoleSlice)

	pepoleSlice = append(pepoleSlice, "岳不群", "风清扬")
	fmt.Print("切片slice增加多个元素：")
	fmt.Println(pepoleSlice)

	pepoleSlice2 := make([]string, len, capacity)
	copy(pepoleSlice2, pepoleSlice)
	fmt.Print("切片slice拷贝后新的切片：")
	fmt.Println(pepoleSlice2)

	pepoleSlice3:= pepoleSlice[0:1]
	fmt.Print("切片slice截取后新的切片：")
	fmt.Println(pepoleSlice3)
}

/*
map的输出
*/
func mapPrint() {
	// 定义map
	peopleMap := make(map[string]string)

	// map 增加数据
	peopleMap["1"] = "令狐冲"
	peopleMap["2"] = "任盈盈"
	peopleMap["3"] = "东方不败"

	// 输出map
	fmt.Println("输出map：")
	for id, people := range peopleMap {
		fmt.Println(id, people)
	}

	// 判断元素是否存在
	fmt.Println("判断map中元素是否存在：")
	_, ok := peopleMap ["令狐冲"]
	if (ok) {
		fmt.Println("人物存在")
	} else {
		fmt.Println("人物不存在")
	}

	// 删除元素
	delete(peopleMap, "1")

	// 输出map
	fmt.Println("删除map中元素后：")
	for id := range peopleMap {
		fmt.Println(id, peopleMap[id])
	}
}
