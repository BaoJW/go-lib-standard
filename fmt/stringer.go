package fmt

import "fmt"

//Stringer 由具有 String 方法的任何值实现，该方法为该值定义“native”格式。String 方法用于将作为操作数传递的值打印为接受字符串或格式未打印的打印机（如Print）的任何格式。

//type Stringer interface {
//	String() string
//}

// Animal has a Name and an Age to represent an animal.
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func StringerModule() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
}
