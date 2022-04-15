package models

type (
	User struct {
		FirstName      string // 姓
		LastName       string //名
		Age            uint8
		Address        *Address
		privateAddress Address
	}
)

// SetFirstName 设置姓
func (user *User) SetFirstName(firstName string) {
	user.FirstName = firstName
}

func PackageFunction() string {
	return "Hello World"
}
