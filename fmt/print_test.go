package fmt

import (
	"fmt"
	"testing"
)

func TestPrint_Errorf(t *testing.T) {
	p := new(Print)

	fmt.Println(p.Errorf())
}

func TestPrint_Fprint(t *testing.T) {
	p := new(Print)

	resp, err := p.Fprint()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("resp is: ", resp)
}

func TestPrint_Fprintf(t *testing.T) {
	p := new(Print)

	resp, err := p.Fprintf()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("resp is: ", resp)
}

func TestPrint_Fprintln(t *testing.T) {
	p := new(Print)

	resp, err := p.Fprintln()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("resp is: ", resp)

}

func TestPrint_Print(t *testing.T) {
	p := new(Print)

	resp, err := p.Print()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("\n")
	fmt.Println("resp is: ", resp)
}

func TestPrint_Printf(t *testing.T) {
	p := new(Print)

	resp, err := p.Printf()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("\n")
	fmt.Println("resp is: ", resp)
}

func TestPrint_Println(t *testing.T) {
	p := new(Print)

	resp, err := p.Println()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("\n")
	fmt.Println("resp is: ", resp)
}

func TestPrint_Sprint(t *testing.T) {
	p := new(Print)

	fmt.Println(p.Sprint())
}

func TestPrint_Sprintf(t *testing.T) {
	p := new(Print)

	fmt.Println(p.Sprintf())
}

func TestPrint_Sprintln(t *testing.T) {
	p := new(Print)

	fmt.Println(p.Sprintln())
}

func TestPrint_Format(t *testing.T) {
	p := new(Print)

	p.Format(p, 's')

}
