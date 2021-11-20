package main

import "fmt"

type Animal struct{
	Name string
}
func(a*Animal)move(){
	fmt.Printf("%s会咬人QAQ\n",a.Name)
}

type Dog struct {
	Feet    int
	*Animal//匿名嵌套结构体指针
}

func (d*Dog) Dong() {
	fmt.Printf("%s会汪汪汪~\n",d.Name)
}
func main() {
	d1 :=&Dog{
		Feet:4,
		Animal:&Animal{
			Name:"蔡英文",
		},
	}
	d1.Dong()
	d1.move()
}