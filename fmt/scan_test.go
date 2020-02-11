package fmt

import (
	"fmt"
	"testing"
)

func TestScan_Fscan(t *testing.T) {
	s := new(Scan)

	result, err := s.Fscan()
	if err != nil {
		fmt.Println("error is: ", err)
	}

	fmt.Println(result)
}

func TestScan_Fscanf(t *testing.T) {
	s := new(Scan)

	result, err := s.Fscanf()
	if err != nil {
		fmt.Println("error is: ", err)
	}

	fmt.Println(result)
}

func TestScan_Fscanln(t *testing.T) {
	s := new(Scan)

	result, err := s.Fscanln()
	if err != nil {
		fmt.Println("error is: ", err)
	}

	fmt.Println(result)
}

func TestScan_Scan(t *testing.T) {
	//关于Scan的测试，默认是从os.stdin即控制台输入
	//如果想使用单测测试，详情请参考官方包fmt.scan_test.go，这里就不在续写了
}

func TestScan_Scanf(t *testing.T) {
	//关于Scanf的测试，默认是从os.stdin即控制台输入
	//如果想使用单测测试，详情请参考官方包fmt.scan_test.go，这里就不在续写了
}

func TestScan_scanln(t *testing.T) {
	//关于Scanln的测试，默认是从os.stdin即控制台输入
	//如果想使用单测测试，详情请参考官方包fmt.scan_test.go，这里就不在续写了
}

func TestScan_Sscan(t *testing.T) {
	s := new(Scan)

	result, err := s.Sscan()
	if err != nil {
		fmt.Println("error is: ", err)
	}

	fmt.Println(result)
}

func TestScan_Sscanf(t *testing.T) {
	s := new(Scan)

	result, err := s.Sscanf()
	if err != nil {
		fmt.Println("error is: ", err)
	}

	fmt.Println(result)
}

func TestScan_Sscanln(t *testing.T) {
	s := new(Scan)

	result, err := s.Sscanln()
	if err != nil {
		fmt.Println("error is: ", err)
	}

	fmt.Println(result)
}
