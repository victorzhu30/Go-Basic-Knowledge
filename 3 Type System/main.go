package main

import (
	"3_Type_System/mypackage"
	"fmt"
)

type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

// 定义admin结构体
type admin struct {
	name string
}

func (a admin) notify() {
	fmt.Printf("admin %s is notified\n", a.name)
}

// 定义user结构体
type user struct {
	name  string
	email string
}

// 值接收者
func (u user) notify() {
	fmt.Printf("user %s is notified\n", u.name)
}

func (u user) changeName(newName string) {
	u.name = newName
}

func (u user) printName() {
	fmt.Println("name:" + u.name)
}

// 指针接收者
func (u *user) changeName2(newName string) {
	u.name = newName
}

type person struct {
	level int
	u     user
}

func main() {
	u := user{
		name:  "heli",
		email: "heli@qq.com",
	}
	u.changeName("chong")
	u.printName()
	//u.changeName2("chong")
	//u.printName()

	a := admin{
		name: "chong",
	}

	// 多态
	sendNotification(u)
	sendNotification(a)

	//大小写
	Cancan1 := mypackage.Can1{
		Name: "chong",
	}
	fmt.Println(Cancan1)

	Cancan2 := mypackage.Can2{}
	fmt.Println(Cancan2)

	p := person{
		level: 100,
		u: user{
			name:  "zrp",
			email: "zrp@qq.com",
		},
	}
	fmt.Println(p)
	p.u.printName()
}
