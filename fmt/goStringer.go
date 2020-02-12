package fmt

import "fmt"

//GoStringer 由具有 GoString 方法的任何值实现，该方法定义该值的 Go 语法。GoString 方法用于将作为操作数传递给 ％＃v 格式的值打印出来
//只要实现了GoString方法，就可以个性化定制打印的格式，可以集成化在打印某些特定的日志上，根据不同的日志内容或者日志等级，打印不同风格的日志内容

//type GoStringer interface {
//	GoString() string
//}

// Address has a City, State and a Country.
type Address struct {
	City    string
	State   string
	Country string
}

// Person has a Name, Age and Address.
type Person struct {
	Name string
	Age  uint
	Addr *Address
}

// GoString makes Person satisfy the GoStringer interface.
// The return value is valid Go code that can be used to reproduce the Person struct.
func (p Person) GoString() string {
	if p.Addr != nil {
		return fmt.Sprintf("Person{Name: %q, Age: %d, Addr: &Address{City: %q, State: %q, Country: %q}}", p.Name, int(p.Age), p.Addr.City, p.Addr.State, p.Addr.Country)
	}
	return fmt.Sprintf("Person{Name: %q, Age: %d}", p.Name, int(p.Age))
}
