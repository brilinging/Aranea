package main

import "fmt"

type person struct{					//人的结构体
	Name string						//姓名
	Age int							//年龄
	Gender string					//性别
}

type dove interface{
	Gugugu()						//鸽
}

type repeater interface{				//复读机
	repeat(string)						//复读
}

type lemon interface{					//柠檬精
	AcheNum() 					   	 	//酸
}

type piggish interface{					//真香怪
	PigNum () 							//真香的次数
}

func(p*person)repeat(word string){		//复读
	fmt.Println(word)
}

func(p*person)AcheNum(){
	fmt.Println(p.Name,"酸了...QAQ")
}

func(p*person)PigNum(){
	fmt.Println(p.Name,"是真香怪@W@")
}

func(p*person)Gugugu(){
	fmt.Println(p.Name,"又鸽了！！！")
}

func main(){							//主函数
	p := &person{
		Name:	"1cm",
		Age:	18,
		Gender:"male",
	}
	p.Gugugu()
	p.AcheNum()
	p.PigNum()

	var r repeater

	r = p
	r.repeat("HelloWorld!")
}