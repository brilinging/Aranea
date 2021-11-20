package main

import "fmt"

type Address struct{		//type关键字造的都是类型
	Province 	string			//省份
	City  	 	string			//城市
	UpdateTime  string			//上传时间
}

type Email struct{
	Addr		string			//邮件地址
	UpdateTime  string			//上传时间
}

type Person struct{
	Name 		string			//姓名
	Age  		int				//年龄
	Gender 		string			//性别
	Address 				//嵌套上面的Address结构体
	Email					//嵌套邮件结构体
}

func main(){				//主函数
	p1 := Person{
		Name: 			"键仙",
		Gender: 		"男",
		Age: 			18,
		Address: Address{
			Province:   "互凉网",
			City: 	    "评论区",
			UpdateTime: "2021,11,13",
		},
		Email:Email{
			Addr:		"3298225433@qq.com",
			UpdateTime: "2022,02,22",
		},
	}
	fmt.Printf("%#v\n",p1)
	//%v:只输出所有的值；
	//%+v：先输出字段名字，再输出该字段的值；
	//%#v：先输出结构体名字值，再输出结构体（字段名字+字段的值）
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.Gender)
	fmt.Println(p1.Address.Province)//通过嵌套的匿名结构体字段访问其内部的字段
	// fmt.Println(p1.Province)——->//此处若直接访问匿名结构体中的字段会发生字段冲突(嵌套结构体中包含多个同名UpdateTime)
	fmt.Println(p1.Address.UpdateTime)
	fmt.Println(p1.Email.Addr)
	fmt.Println(p1.Email.UpdateTime)
}
