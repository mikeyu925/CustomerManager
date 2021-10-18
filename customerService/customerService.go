package service
//实现对客户管理的增删改查等功能
import(
	"gocode/customerManager/customer"
)

type CustomerService struct{
	//一个切片，来存储每个顾客的信息
	customers [] customer.Customer
	customerNum int64
}

func NewCustomerService() *CustomerService{
	customerservice  := &CustomerService{}
	customerservice.customerNum = 1
	//id ,age int64,name ,gender, phone, email
	customer := customer.NewCustomer(1,20,"张三","男","151414953","fish@qq.com")
	customerservice.customers = append(customerservice.customers,customer)
	return customerservice
}

//返回用户切片
func (this * CustomerService)List() [] customer.Customer{
	return this.customers
}

//添加用户
func (this * CustomerService)Add(customer customer.Customer)bool{
	this.customerNum ++
	customer.ID = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}

//根据id删除用户
func (this * CustomerService)DeleteByID(id int64)bool{
	index := this.FindByID(id)
	if index == -1{
		return false
	}
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
}

//根据id查找用户
func (this * CustomerService)FindByID(id int64) int{
	index := -1
	for i := 0;i < len(this.customers);i++{
		if this.customers[i].ID == id{
			index = i
			break
		}
	}
	return index
}

func (this *CustomerService)Alter(id int64,customer customer.Customer)bool{
	index := this.FindByID(id)
	if index == -1{
		return false
	}
	this.customers[index].Age = customer.Age
	this.customers[index].Name = customer.Name
	this.customers[index].Gender = customer.Gender
	this.customers[index].Email = customer.Email
	this.customers[index].Phone = customer.Phone
	return true
}