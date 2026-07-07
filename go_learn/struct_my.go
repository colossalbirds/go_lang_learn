package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"` // 结构体标签
	Age  int    `json:"age"`
}

type Student struct {
	People     // 继承People结构体
	Grade  int `json:"grade"`
}

// PrintInfo 给机构体绑定一个方法
func (s Student) PrintInfo() {
	fmt.Printf("name:%s age:%d grade:%d\n", s.Name, s.Age, s.Grade)
	byteData, err := json.Marshal(s) // 将结构体转换为JSON格式
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}
	fmt.Println("JSON:", string(byteData))
}
