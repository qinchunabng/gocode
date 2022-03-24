package main

import (
	"fmt"
)

func main() {
	var key string
	var loop = true
	//定义账户余额
	var balance = 10000.0
	//每次收支的金额
	var money = 0.0
	//收支说明
	var note = ""
	//收支详情
	var detail = "收支\t账户金额\t收支金额\t说    明"
	//显示主菜单
	for {
		fmt.Println("\n---------------------家庭收支记账软件--------------------")
		fmt.Println("                      1 收支明细")
		fmt.Println("                      2 登记收入")
		fmt.Println("                      3 登记支出")
		fmt.Println("                      4 退出软件")
		fmt.Println("请选择(1-4)：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("---------------------当前收支明细--------------------")
			fmt.Println(detail)
		case "2":
			fmt.Print("本次收入金额:")
			fmt.Scanln(&money)
			balance += money
			fmt.Print("本次收入说明:")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			fmt.Println(detail)
		case "3":
			fmt.Println("登记支出")
		case "4":
			loop = false
		default:
			fmt.Println("请输出正确的选项..")
		}

		if !loop {
			break
		}
	}

	fmt.Println("退出家庭记账的软件的使用..")
}
