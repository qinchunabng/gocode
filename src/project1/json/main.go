package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := new(Monster)
	m.Name = "牛魔王"
	m.Age = 500
	m.Skill = "牛魔拳"
	m.Address = "火焰山"
	jsonStr := testStruct(m)

	//json反序列化
	var test Monster
	unmarshal(jsonStr, &test)
	fmt.Println(test)

	//反序列化map不需要make，json.Unmarshal里面会做这个操作
	var testMap map[string]interface{}
	unmarshal(jsonStr, &testMap)
	fmt.Printf("testMap=%v\n", testMap)
}

type Monster struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Skill   string
}

func testStruct(s interface{}) string {
	//转json字符串
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("转json错误:%v\n", err)
		return ""
	}
	json := string(data)
	fmt.Printf("转json结果:%v\n", json)
	return json
}

//json字符串反序列化
func unmarshal(jsonStr string, s interface{}) error {
	return json.Unmarshal([]byte(jsonStr), s)
}
