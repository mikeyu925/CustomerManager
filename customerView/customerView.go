package main
//379 / 656
import(
	"fmt"
	"gocode/customerManager/customer"
	"gocode/customerManager/customerService"
)

type CustomerView struct{
	key string
	loop bool

	customerService * service.CustomerService
}

func (this * CustomerView) list(){
	customers := this.customerService.List()

	fmt.Println("      客户列表")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:=0;i < len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------------------------------")
}

func (this * CustomerView)MainMenu(){
	for{// 等价于while(1)
		fmt.Println("-----------客户信息管理系统-----------")
		fmt.Println("           1.添加客户")
		fmt.Println("           2.修改客户")
		fmt.Println("           3.删除客户")
		fmt.Println("           4.显示客户")
		fmt.Println("           5.退出系统")
		fmt.Println("输入命令(1-5):")
		//读取命令
		fmt.Scanln(&this.key)
		//判断命令
		switch this.key{
		case"1":
			this.addCustomer()
		case"2":
			this.alter()
		case"3":
			this.delete()
		case"4":
			this.list()
		case"5":
			{
				fmt.Println("确认退出系统(Y/N)")
				choice := ""
				fmt.Scanln(&choice)
				if choice == "y" || choice == "Y"{
					fmt.Println("正在退出系统...")
					this.loop = false
				}
			}
		default:
			fmt.Println("输入有误！请重新输入!")
		}
		if !this.loop {
			fmt.Println("退出系统成功!")
			break
		}
	}

}

func (this * CustomerView) addCustomer(){
	fmt.Println("-------添加用户-------")
	fmt.Printf("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年龄:")
	var age int64 = 0
	fmt.Scanln(&age)
	fmt.Printf("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮箱:")
	email := ""
	fmt.Scanln(&email)
	customer := customer.NewCustomer2(age,name ,gender, phone, email)
	if this.customerService.Add(customer){
		fmt.Println("---------用户添加完成--------")
	}else{
		fmt.Println("---------用户添加失败--------")
	}
}

func (this * CustomerView)delete(){
	var id int64 = -1
	fmt.Println("-------删除用户-------")
	fmt.Printf("请输入待删除用户id(输入-1取消):")
	fmt.Scanln(&id)
	if id == -1{
		return 
	}
	fmt.Println("确认是否删除?(Y/N)")

	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y"{
		if this.customerService.DeleteByID(id){
			fmt.Println("成功删除!")
		}else{
			fmt.Println("删除失败，输入id错误!")
		}
	}
}

func (this * CustomerView)alter(){
	var id int64 = -1
	fmt.Println("-------修改用户-------")
	fmt.Printf("请输入待修改用户id(输入-1取消):")
	fmt.Scanln(&id)
	if id == -1{
		return 
	}
	fmt.Printf("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年龄:")
	var age int64 = 0
	fmt.Scanln(&age)
	fmt.Printf("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮箱:")
	email := ""
	fmt.Scanln(&email)
	customer := customer.NewCustomer2(age,name ,gender, phone, email)
	fmt.Println("确认是否修改?(Y/N)")

	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y"{
		if this.customerService.Alter(id,customer){
			fmt.Println("成功修改!")
		}else{
			fmt.Println("修改失败，输入id错误!")
		}
	}
}

func main(){
	View := CustomerView{
		key:"",
		loop:true,
	}
	View.customerService = service.NewCustomerService()
	//进入循环，显示主菜单
	View.MainMenu()
}