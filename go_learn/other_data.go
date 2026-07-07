package main

import "fmt"

type Code int

const (
	SuccessCode    Code = 0
	ValidCode      Code = 7 // 校验失败的错误
	ServiceErrCode Code = 8 // 服务错误
)

func (c Code) GetMsg() string {
	// 可能会有更加响应码返回不同消息内容的要求，我们在这个函数里面去实现即可
	// 可能还会有国际化操作
	return "成功"
}

// Animal 定义一个animal的接口，它有唱，跳，rap的方法
type Animal interface {
	sing()
	jump()
	rap()
}

// Chicken 需要全部实现这些接口
type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Println("chicken 唱")
}

func (c Chicken) jump() {
	fmt.Println("chicken 跳")
}
func (c Chicken) rap() {
	fmt.Println("chicken rap")
}

// 全部实现完之后，chicken就不再是一只普通的鸡了
