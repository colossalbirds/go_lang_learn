package main

import (
	"fmt"
	"sync"
)

func list_my() {

	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array a:", a)
	b := []int{6, 7, 8, 9, 10}
	fmt.Println("Slice b:", b)
	fmt.Println("a[0]:%d", a[0])
	var c []string = []string{"apple", "banana", "cherry"}
	fmt.Println("Slice c:", c)
	fmt.Println("c[1]:", c[1])
	fmt.Println("数组增加")
	c = append(c, "date")
	fmt.Println("Slice c after append:", c)
	fmt.Println("c[3]:", c[3])
	fmt.Println("数组删除")
	c = append(c[:1], c[2:]...)
	fmt.Println("Slice c after deletion:", c)
	fmt.Println("c[1]:", c[1])
	fmt.Println("数组长度函数")
	fmt.Println("len(c):", len(c))
	fmt.Println("cap(c):", cap(c))
	fmt.Println("判断数组是否为空")
	fmt.Println("c == nil:", c == nil)
	fmt.Println("make函数的使用")
	d := make([]string, 3)
	fmt.Println("Slice d:", d)
	fmt.Println("map函数的使用")
	e := make(map[string]int)
	e["apple"] = 5
	e["banana"] = 10
	fmt.Println("Map e:", e)
	fmt.Println("e[\"apple\"]:", e["apple"])
	var m1 = map[string]int{
		"age": 21,
	}
	age1 := m1["age1"] // 取一个不存在的
	fmt.Println(age1)
	age2, ok := m1["age1"]
	fmt.Println(age2, ok)
}
func judge() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	if age <= 0 {
		fmt.Println("未出生")
		return
	}
	if age <= 18 {
		fmt.Println("未成年")
		return
	}
	if age <= 35 {
		fmt.Println("青年")
		return
	}
	fmt.Println("中年")
	// aa := age.(int) // 断言
	// fmt.Println("断言后的值：", aa)

}
func judge2() {
	fmt.Println("请输入星期数字：")
	var week int
	fmt.Scan(&week)

	switch week {
	case 1:
		fmt.Println("周一")
	case 2:
		fmt.Println("周二")
	case 3:
		fmt.Println("周三")
	case 4:
		fmt.Println("周四")
	case 5:
		fmt.Println("周五")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("错误")
	}
}

func circle() {
	s := []string{"枫枫", "知道"}
	for index, s2 := range s {
		fmt.Println(index, s2)
	}

}
func functions(a string, b string) (res string) {
	for i := 0; i < 5; i++ {
		res += a + b
	}
	return
}

var num int
var wait sync.WaitGroup
var lock sync.Mutex

func add() {
	// 谁先抢到了这把锁，谁就把它锁上，一旦锁上，其他的线程就只能等着
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num++
	}
	lock.Unlock()
	wait.Done()
}
func reduce() {
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num--
	}
	lock.Unlock()
	wait.Done()
}
