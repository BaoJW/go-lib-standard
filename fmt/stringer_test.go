package fmt

import (
	"fmt"
	"testing"
)

func TestAnimal_String(t *testing.T) {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
}
