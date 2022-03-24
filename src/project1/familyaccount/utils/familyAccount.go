package utils

import (
	"fmt"
	"strconv"
)

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	note    string
	flag    bool
	details string
}

// MainMenu 显示主菜单
func (this *FamilyAccount) MainMenu() {
	this.loop = true
	this.flag = false
	this.details = "收支\t账户金额\t收支金额\t说    明"
	this.balance = 1000
	for {
		fmt.Println("\n---------------------家庭收支记账软件--------------------")
		fmt.Println("                      1 收支明细")
		fmt.Println("                      2 登记收入")
		fmt.Println("                      3 登记支出")
		fmt.Println("                      4 退出软件")
		fmt.Println("请选择(1-4)：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输出正确的选项..")
		}

		if !this.loop {
			break
		}
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("---------------------当前收支明细--------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明显...请来一笔！")
	}

}

//登记收入
func (this *FamilyAccount) income() {
	this.flag = true
	input := ""
	var money float64 = 0.0
	var err error
	fmt.Print("本次收入金额:")
	fmt.Scanln(&input)
	money, err = strconv.ParseFloat(input, 64)
	for {
		if money > 0 && err == nil {
			break
		} else {
			fmt.Print("请输入正确的收入金额:")
			//time.Sleep(time.Millisecond * 50)
			fmt.Scanln(&input)
			money, err = strconv.ParseFloat(input, 64)
		}
	}
	this.balance += money
	fmt.Print("本次收入说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, money, this.note)
	fmt.Println(this.details)
}

//支出
func (this *FamilyAccount) pay() {
	this.flag = true
	money := 0.0
	fmt.Print("请输入本次支出金额：")
	fmt.Scanf("%f", &money)
	for {
		if money > 0 {
			break
		} else {
			fmt.Print("请输入正确的支出金额:")
			fmt.Scanf("%f", &money)
		}
	}
	if money > this.balance {
		fmt.Println("余额不足")
	} else {
		this.balance -= money
		fmt.Print("请输入本次支出说明：")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, money, this.note)
		fmt.Println(this.details)
	}
}

//退出
func (this *FamilyAccount) exit() {
	fmt.Print("你确定要退出吗?(y/n)")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			this.loop = !(choice == "y")
			break
		} else {
			fmt.Print("你输入有误，请重新输入(y/n)")
		}
	}
}
