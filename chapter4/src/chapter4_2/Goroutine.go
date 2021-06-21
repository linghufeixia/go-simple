/**
 * Created with IntelliJ IDEA.
 * User: 令狐飞侠
 * Date: 2021-6-18
 * Description: 线程goroutine的使用
 */
package chapter4_2

import (
	"fmt"
	"sync"
	"time"
)

//goroutine的使用
func TestGoroutine() {
	var x = 0
	for i := 0; i < 5; i++ {
		// go 函数或函数名 启动 goroutine
		go func() {
			x++
			fmt.Println(x)
		}()
	}
	//time.Sleep(time.Second * time.Duration(1))
}

//channel的使用
func TestChannel() {
	//make创建channel
	channel1 := make(chan string)
	//channel写入数据
	go channelWrite(channel1)
	//channel读取数据
	channelRead(channel1)

	//关闭channel
	close(channel1)
}

//channel 写入数据
func channelWrite(channel1 chan string) {
	//channel写入
	channel1 <- "I "
	fmt.Println("channel中写入第1个字符串：I")
	channel1 <- "love"
	fmt.Println("channel中写入第2个字符串：love")
	channel1 <- "you"
	fmt.Println("channel中写入第3个字符串：you")
}

//channel 读取数据
func channelRead(channel1 chan string)  {
	// channel输出
	a := <-channel1
	fmt.Println("从channel中取第1个字符串：", a)
	b := <-channel1
	fmt.Println("从channel中取第2个字符串：", b)
	c := <-channel1
	fmt.Println("从channel中取第3个字符串：", c)
}

// 测试缓存区channel
func TestBufferChannel() {
	// 设置channel缓存区的大小
	channel2 := make(chan int, 3)
	// channel缓存区写入
	go bufferChannelWrite(channel2)
	//channel缓存区读取
	bufferChannelRead(channel2)
}

//channel 缓存区写入数据
func bufferChannelWrite(channel2 chan int)  {
	channel2 <- 3
	fmt.Println("channel缓存区中写入第1个数：3")
	channel2 <- 8
	fmt.Println("channel缓存区中写入第2个数：8")
	channel2 <- 9
	fmt.Println("channel缓存区中写入第3个数：9")
	//关闭channel
	close(channel2)
}

//channel 缓存区读取数据
func bufferChannelRead(channel2 chan int)  {
	for a := range channel2 {
		fmt.Println("channel缓存区中读取元素:" , a)
	}
}


// goroutine阻塞，select的使用 channel协程阻塞
func TestSelect() {
	channel3 := make(chan int)
	go func() {
		// 线程阻塞
		time.Sleep(time.Duration(10) * time.Second)
		channel3 <- 1
	}()

	select {
	case i := <-channel3:
		fmt.Println("从channel接收到%d",i)
	default:
		fmt.Println("其他协程都阻塞了。。。")
	}
}

var(
	x = 0
	wg = sync.WaitGroup{}
	m sync.Mutex
)

// goroutine数据同步，WaitGroup、互斥锁的使用 多协程数据同步
func TestMutex()  {
	var x = 0
	for i:=0;i<5;i++ {
		go func() {
			wg.Add(1)
			m.Lock()
			defer wg.Done()
			defer m.Unlock()
			x++
			fmt.Println(x)
		}()
	}
	wg.Wait()
}
