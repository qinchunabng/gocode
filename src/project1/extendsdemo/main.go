package main

import "fmt"

func main() {
	//pupil := &Pupil{}
	//pupil.Name = "tom~"
	//pupil.Age = 8
	//pupil.testing()
	//pupil.setScore(100)
	//pupil.showInfo()
	//
	//graduate := &Pupil{}
	//graduate.Name = "mary~"
	//graduate.Age = 18
	//graduate.testing()
	//graduate.setScore(90)
	//graduate.showInfo()

	var b B
	b.Name = "Tom"
	//继承的可以访问匿名类的所有属性
	b.age = 20
	//如果有同名的函数和属性，采取就近原则
	b.SayOk()
	//如果要访问匿名类的方法需要明确指定
	b.A.SayOk()
	//如果属性覆盖之后，方法没有覆盖，则输出默认值
	b.hello()
	b.A.Name = "Eric"
	b.hello()

	//多继承
	var c C
	//如果多继承多个父结构体都有相同属性，赋值的时候必须明确指定给哪个属性赋值
	//如果C结构体本身有Name字段，则不会出现该问题
	c.A.Name = "Scott"
	//如果嵌套了有名结构体
	var a A
	a.Name = "Alan"
	c.a = a
	c.a.SayOk()

	//匿名结构体创建时赋值
	tv := TV{
		&Goods{
			Name:  "电视机",
			Price: 1000,
		},
		&Brand{
			Name:    "海尔",
			Address: "山东青岛",
		},
	}
	fmt.Println(*tv.Brand, *tv.Goods)

	var e E
	e.int = 20
	fmt.Println(e)
}

type Goods struct {
	Name  string
	Price int
}

type Book struct {
	Goods  //继承Goods
	Writer string
}

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) showInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) setScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在测试中...")
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在测试中...")
}

type A struct {
	Name string
	age  int
}

func (a *A) hello() {
	fmt.Println("A hello,", a.Name)
}

func (a *A) SayOk() {
	fmt.Println("A SayOk,", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B SayOk,", b.Name)
}

type C struct {
	A
	B
	a A
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	*Goods
	*Brand
}

type E struct {
	int
	//int 不能有多个相同的匿名结构体，不然无法区分
}
