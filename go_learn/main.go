package main

import "fmt"

func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}
func init() {
	fmt.Println("init3")
}

func main() {
	// a, b := 5, 10
	// sum := calculateSum(a, b)
	// fmt.Println("Sum:", sum)
	// test1()
	// list_my()
	// judge2()
	// circle()
	// var a string = "Hello, World!"
	// var b string = " Hello, Go!"
	// var c string
	// c = functions(a, b)
	// fmt.Println(c)
	// s := Student{
	// 	People: People{
	// 		Name: "枫枫",
	// 		Age:  21,
	// 	},
	// 	Grade: 85,
	// }
	// s.Name = "枫枫知道" // 修改值
	// s.PrintInfo()
	// fmt.Println(SuccessCode.GetMsg())
	// var i int
	// fmt.Println(int(SuccessCode) == i)
	var animal Animal

	animal = Chicken{"ik"}

	animal.sing()
	animal.jump()
	animal.rap()
}
