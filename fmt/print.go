/*
 *	Author: lingfeng
 *	Created: 2020-02-03 20:30
 *	Description: Explanation and demo of the print class method of the official fmt package
 */

package fmt

import (
	"fmt"
	"os"
)

/*
	所有print类方法的底层实现大致都分为三个步骤
	1. 从pool中获取打印指针，定义打印器
	2. 打印buf中的内容到标准输出
	3. 释放内存，将打印指针放回pool中
*/

type Print struct {
	buffer []byte
}

//Errorf 格式根据格式说明符并返回字符串作为满足错误的值。
func (p *Print) Errorf() error {
	return fmt.Errorf("%s", "this is a error")
}

//使用其操作数的缺省格式并写入w的 Fprint 格式。当两个操作数都不是字符串时，空间会被添加。它返回写入的字节数和遇到的任何写入错误。
func (p *Print) Fprint() (int, error) {
	writer, err := os.OpenFile("./print_test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}

	return fmt.Fprint(writer, "Fprint: do you love me, my dear")
}

//Fprintf 根据格式说明符格式化并写入w。它返回写入的字节数和遇到的任何写入错误。
func (p *Print) Fprintf() (int, error) {
	writer, err := os.OpenFile("./print_test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}

	return fmt.Fprintf(writer, "%s", "Fprintf: do you love me, my dear")
}

//Fprintf 根据格式说明符格式化并写入w。它返回写入的字节数和遇到的任何写入错误。
func (p *Print) Fprintln() (int, error) {
	writer, err := os.OpenFile("./print_test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}

	return fmt.Fprintln(writer, "Fprintln: do you love me, my dear")
}

//使用其操作数的默认格式打印格式并写入标准输出。当两个操作数都不是字符串时，空间会被添加。它返回写入的字节数和遇到的任何写入错误。
func (p *Print) Print() (int, error) {
	return fmt.Print("this is print")
}

//Printf 格式根据格式说明符写入标准输出。它返回写入的字节数和遇到的任何写入错误。
func (p *Print) Printf() (int, error) {
	return fmt.Printf("%s", "this is printf")
}

//Println 格式使用其操作数的默认格式并写入标准输出。总是在操作数之间添加空格，并附加一个换行符。它返回写入的字节数和遇到的任何写入错误。
func (p *Print) Println() (int, error) {
	return fmt.Println("this is println")
}

//Sprint 格式使用其操作数的默认格式并返回结果字符串。当两个操作数都不是字符串时，空间会被添加。
func (p *Print) Sprint() string {
	return fmt.Sprint("this is Sprint")
}

//Sprintf 根据格式说明符格式化并返回结果字符串。
func (p *Print) Sprintf() string {
	return fmt.Sprintf("%s", "this is Sprintf")
}

//Sprintln 格式使用其操作数的默认格式并返回结果字符串。总是在操作数之间添加空格，并附加一个换行符。
func (p *Print) Sprintln() string {
	return fmt.Sprintln("this is Sprintf")
}

//Formatter 格式化程序是通过使用自定义格式化程序的值实现的接口。Format 的实现可能会调用 Sprint(f) 或 Fprint(f) 等来生成其输出。

//type Formatter interface {
//	Format(f State, c rune)
//}

func (p *Print) Format(s fmt.State, c rune) {
	_, _ = fmt.Fprintf(s, "%"+string(c), "hello world")
}

func (p *Print) Write(b []byte) (n int, err error) {
	p.buffer = append(p.buffer, b...)
	return len(p.buffer), nil
}

func (p *Print) Width() (wid int, ok bool) {
	return 0, true
}

func (p *Print) Precision() (prec int, ok bool) {
	return 0, true
}

func (p *Print) Flag(c int) bool {
	return true
}
