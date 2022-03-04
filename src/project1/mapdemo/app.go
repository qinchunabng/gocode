package main

import "fmt"

func main() {
	var m map[string]string
	//m = make(map[string]string, 10)
	m = make(map[string]string)
	m["hello"] = "world"
	fmt.Println(m)

	var cities map[string]string = map[string]string{
		"no2": "武汉",
	}
	cities["no1"] = "北京"
	fmt.Println(cities)

	//删除
	delete(cities, "no1")
	fmt.Println(cities)

	//如果需要清空map
	//1.遍历delete
	//2.直接make一个新的

	//map的查找
	val, ok := cities["no1"]
	if ok {
		fmt.Printf("有no1 key值为%v\n", val)
	} else {
		fmt.Println("没有no1 key\n")
	}

	//map遍历
	for k, v := range cities {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}
	fmt.Println("len=", len(cities))

	//map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 1)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	//这里会数组越界
	//if monsters[1] == nil {
	//	monsters[1] = make(map[string]string)
	//	monsters[1]["name"] = "玉兔精"
	//	monsters[1]["age"] = "200"
	//}
	//动态增加需要用append
	newMonster := map[string]string{
		"name": "狐狸精",
		"age":  "300",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}
