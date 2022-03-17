package main

import (
	"fmt"
)

func main() {
	//computer := Computer{}
	//var u = computer
	//phone := Phone{}
	//camera := Camera{}
	//
	//computer.work(phone)
	//computer.work(camera)
	//
	//var i integer = 10
	//var usb Usb = i
	//usb.Start()
	//
	//var a test.AInterface = i
	//a.A()
	//
	////任何类型都可以赋值给空接口
	//var t T = computer
	//fmt.Println(t)
	//
	//var hs test.HeroSlice
	//for i := 0; i < 10; i++ {
	//	hero := test.Hero{
	//		Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
	//		Age:  rand.Intn(100),
	//	}
	//	hs = append(hs, hero)
	//}
	////排序前的顺序
	//fmt.Println("排序前的顺序:")
	//for _, h := range hs {
	//	fmt.Println(h)
	//}
	//sort.Sort(hs)
	////排序后的顺序
	//fmt.Println("排序后的顺序:")
	//for _, h := range hs {
	//	fmt.Println(h)
	//}

	//类型断言
	var t float32
	var x interface{}
	x = t
	y, ok := x.(int)
	if ok {
		fmt.Println("转换成功", y)
	} else {
		fmt.Println("转换失败")
	}

}

func (i integer) A() {
	fmt.Println("A()", i)
}

type T interface {
}

//只要是自定义数据类型都可以实现接口
type integer int

func (i integer) Start() {
	fmt.Println("Start", i)
}

func (i integer) Stop() {
	fmt.Println("Stop", i)
}

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

func (p Phone) Call() {
	fmt.Println("拨打电话")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

func (c Computer) work(usb Usb) {
	usb.Start()
	usb.Stop()
}
