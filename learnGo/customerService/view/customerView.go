package main

import (
	"fmt"
	"learnGo/customerService/model"
	"learnGo/customerService/service"
)

type customerView struct {
	//定义必要字段
	key  string //接受用户输入...
	loop bool   //表示是否循环的显示主菜单
	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示所有的客户信息
func (cv *customerView) list() {

	//首先获取到当前所有的客户信息
	customers := cv.customerService.List()
	//显示
	fmt.Println("-------------------------客户列表--------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-------------------------客户列表完成--------------------------")
}

//得到用户的输入，信息构建新的客户，并完成添加
func (cv *customerView) add() {
	fmt.Println("-------------------------添加客户--------------------------")
	fmt.Println("姓名：")
	name := ""
	_, _ = fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	_, _ = fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	_, _ = fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	_, _ = fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	_, _ = fmt.Scanln(&email)
	//构建一个新的Customer实例
	//注意: id号，没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("-------------------------添加完成--------------------------")
	} else {
		fmt.Println("-------------------------添加失败--------------------------")
	}
}

//得到用户的输入id，删除该id对应的客户
func (cv *customerView) delete() {
	fmt.Println("-------------------------删除客户--------------------------")
	fmt.Print("请选择待删除的客户编号（-1退出）：")
	id := -1
	_, _ = fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N)：")

	choice := ""
	_, _ = fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		//调用customerService 的 Delete方法
		if cv.customerService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("---------------------删除失败，输入的id号不存在----")
		}
	}
}

//退出软件
func (cv *customerView) exit() {
	fmt.Print("确认是否退出(Y/N):")
	for {
		_, _ = fmt.Scanln(&cv.key)
		if cv.key == "Y" || cv.key == "y" || cv.key == "N" || cv.key == "n" {
			break
		}
		fmt.Print("你的输入有误，确认是否退出（Y/N）：")
	}

	if cv.key == "y" || cv.key == "Y" {
		cv.loop = false
	}
}

//显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")
		_, _ = fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("已退出了客户关系管理系统...")
}

func main() {
	//在main函数中，创建一个customerView并运行显示主菜单
	customerView := customerView{
		key:  "",
		loop: true,
	}

	//这里完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
