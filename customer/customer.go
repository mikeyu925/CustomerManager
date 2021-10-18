package customer

import(
	"fmt"
)


type Customer struct{
	ID int64
	Age int64
	Name string
	Gender string
	Phone string
	Email string
}

//简单工厂模式返回一个 Customer对象
func NewCustomer(id ,age int64,name ,gender, phone, email string) Customer{
	return Customer{
		ID:id,
		Age:age,
		Name:name,
		Gender:gender,
		Phone:phone,
		Email:email,
	}
}

func NewCustomer2(age int64,name ,gender, phone, email string)Customer{
	return Customer{
		Age:age,
		Name:name,
		Gender:gender,
		Phone:phone,
		Email:email,
	}
}

func (this Customer)GetInfo() string{
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",
		this.ID,this.Name,this.Gender,this.Age,this.Phone,this.Email)
	return info
}