//因为鄙人学习进度慢，目前只能做到这种程度了,学长海涵...@A@
package main

import (
	"fmt"
)

var chan1 = make(chan int, 20)		//创建用于传结果的通道
var chan2 = make(chan int, 20)		//创建用于传n值的通道

func factorial(n int) {				//factorial函数
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	chan1 <- res					//导入n的阶乘结果
	chan2 <- n  					//导入n
}

func main() {
	for i := 1; i <= 20;i++ {		//循环
		go factorial(i)				//启动一个goroutine去执行factorial函数
	}
	var m = map[int]int{
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
		<-chan2:<-chan1,
	}					//用导管给map传入20个键和值。。。
	for i, v := range m {						//遍历map中的键和其对应的值
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
}
