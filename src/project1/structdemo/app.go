package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//方式一
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println("cat1=", cat1)
	if cat1.Slice == nil {
		fmt.Println("cat1.Slice == nil")
	}
	if cat1.Map1 == nil {
		fmt.Println("cat1.Map1 == nil")
	}
	cat1.Slice = make([]int, 10)
	cat1.Map1 = make(map[string]string)
	fmt.Println("cat1=", cat1)

	//结构体默认是值拷贝
	var cat2 = cat1
	cat2.Name = "小黑"
	cat2.Color = "黑色"
	fmt.Println("cat1=", cat1)
	fmt.Println("cat2=", cat2)

	//方式二
	var cat3 = Cat{Name: "小灰", Color: "灰色"}
	fmt.Println("cat3=", cat3)

	//方式三
	var cat4 *Cat = new(Cat)
	(*cat4).Name = "小花"
	//go设计者为了程序员使用方便，底层对cat4.Name=""进行了处理
	//会给cat4加上取值运算符(*cat4).Name=""
	cat4.Name = "花花"
	fmt.Println("cat4=", cat4)

	//方式四
	var cat5 *Cat = &Cat{}
	cat5.Name = "白白"
	fmt.Println("cat5=", cat5)

	r := Rec{LeftUp: Point{
		X: 10, Y: 10,
	}, RightDown: Point{
		X: 20, Y: 20,
	}}
	fmt.Printf("r的LeftUp.X地址:%p,r的LeftUp.Y地址:%p,r的RightDown.X的地址:%p,r的RightDown.Y的地址:%p\n",
		&r.LeftUp.X, &r.LeftUp.Y, &r.RightDown.X, &r.RightDown.Y)

	var p1 = Point{5, 5}
	var p2 = Point{10, 10}
	var r1 = Rec1{LeftUp: &p1, RightDown: &p2}
	fmt.Printf("r1的LeftUp.X地址:%p,r1的LeftUp.Y地址:%p,r1的RightDown.X的地址:%p,r1的RightDown.Y的地址:%p\n",
		&r1.LeftUp.X, &r1.LeftUp.Y, &r1.RightDown.X, &r1.RightDown.Y)

	monster := Monster{"牛魔王", 500, "芭蕉扇"}
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json序列化错误", err)
	}
	fmt.Println("jsonStr=", string(jsonStr))
	monster.test()
	fmt.Println("Monster Name:", monster.Name)

	//返回结构体指针
	var p3 = &Point{X: 10, Y: 20}
	fmt.Printf("p1 type %T", p3)
}

//定义一个cat结构体
type Cat struct {
	Name  string
	Age   int
	Color string
	Slice []int
	Map1  map[string]string
}

type Point struct {
	X int
	Y int
}

type Rec struct {
	LeftUp, RightDown Point
}

type Rec1 struct {
	LeftUp, RightDown *Point
}

type Monster struct {
	Name  string `json:"name"` //结构体标签
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

//表明是Monster结构体的方法
func (m *Monster) test() {
	m.Name = "铁扇公主"
	fmt.Println("test", m.Name)
}
