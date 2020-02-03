/*
 *	Author: lingfeng
 *	Created: 2020-02-03 20:30
 *	Description: Explanation and demo of the scan class method of the official fmt package
 */

package fmt

import (
	"fmt"
	"os"
)

type Scan struct {
}

//Fscan 扫描从r读取的文本，将连续的空格分隔值存储为连续的参数。换行占据空间。它返回成功扫描的项目数量。如果这小于参数的数量，err 会报告原因。
func (s *Scan) Fscan() (string, error) {
	var (
		name    string
		age     int
		married bool
	)

	_, err := fmt.Fscan(os.Stdin, &name, &age, &married)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Fscan 扫描结果 name:%s age:%d married:%t\t", name, age, married), nil
}

//Fscanf 扫描从 r 读取的文本，将连续的空格分隔值存储为由格式确定的连续参数。它返回成功解析的项目数。输入中的换行符必须与格式中的换行符匹配。
func (s *Scan) Fscanf() (string, error) {
	var (
		name    string
		age     int
		married bool
	)

	_, err := fmt.Fscanf(os.Stdin, "%s%d%t", &name, &age, &married)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Fscanf 扫描结果 name:%s age:%d married:%t\t", name, age, married), nil
}

//Fscanln 与 Fscan 类似，但停止在换行符处进行扫描，并且在最终项目后必须有换行符或 EOF。
func (s *Scan) Fscanln() (string, error) {
	var (
		name    string
		age     int
		married bool
	)

	_, err := fmt.Fscanln(os.Stdin, &name, &age, &married)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Fscanln 扫描结果 name:%s age:%d married:%t\t", name, age, married), nil

}

//扫描扫描从标准输入读取的文本，将连续的空格分隔值存储为连续的参数。换行占据空间。它返回成功扫描的项目数量。如果这小于参数的数量，err 会报告原因。
func (s *Scan) Scan() (string, error) {
	var (
		name    string
		age     int
		married bool
	)

	_, err := fmt.Scan(&name, &age, &married)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Scan 扫描结果 name:%s age:%d married:%t\t", name, age, married), nil

}

//Scanf 扫描从标准输入中读取的文本，将连续的空格分隔值存储为由格式确定的连续参数。它返回成功扫描的项目数量。如果这小于参数的数量，err 会报告原因。输入中的换行符必须与格式中的换行符匹配。一个例外：动词％c总是扫描输入中的下一个符文，即使它是空格（或制表符等）或换行符。
func (s *Scan) Scanf() (string, error) {
	var (
		name    string
		age     int
		married bool
	)

	_, err := fmt.Scanf("%s%d%v", &name, &age, &married)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Scanf 扫描结果 name:%s age:%d married:%t\t", name, age, married), nil
}

//Scanln 与 Scan 类似，但停止在换行符处进行扫描，并且在最终项目后必须有换行符或 EOF。
func (s *Scan) Scanln() (string, error) {
	var (
		name    string
		age     int
		married bool
	)

	_, err := fmt.Scanln(&name, &age, &married)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Scanln 扫描结果 name:%s age:%d married:%t\t", name, age, married), nil
}

//Sscan 扫描参数字符串，将连续的空格分隔值存储为连续的参数。换行占据空间。它返回成功扫描的项目数量。如果这小于参数的数量，err 会报告原因。
func (s *Scan) Sscan() error {
	var (
		name string
		age  int
	)

	_, err := fmt.Sscan("Kim 22", &name, &age)
	return err
}

//Sscanf 扫描参数字符串，将连续的空格分隔值存储为由格式确定的连续参数。它返回成功解析的项目数。输入中的换行符必须与格式中的换行符匹配。
func (s *Scan) Sscanf() error {
	var (
		name string
		age  int
	)

	_, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	return err
}

//Sscanln 与 Sscan 类似，但停止在换行符处进行扫描，并且在最终项目后必须有换行符或 EOF。
func (s *Scan) Sscanln() error {
	var (
		name string
		age  int
	)

	_, err := fmt.Sscanln("Kim 22", &name, &age)
	return err
}
