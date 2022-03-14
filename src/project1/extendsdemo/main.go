package main

import "fmt"

func main() {
	pupil := &Pupil{}
	pupil.Name = "tom~"
	pupil.Age = 8
	pupil.testing()
	pupil.setScore(100)
	pupil.showInfo()

	graduate := &Pupil{}
	graduate.Name = "mary~"
	graduate.Age = 18
	graduate.testing()
	graduate.setScore(90)
	graduate.showInfo()
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
