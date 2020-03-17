# 概述
Flag包实现了命令行标志解析

### 用法
使用flag.String(), Bool(), Int()等定义标志。

demo1: 这声明了一个整型flag，名称为-flagname，都存储在指针ip中，类型为*int。
```asciidoc
import "flag"
var ip = flag.Int("flagname", 1234, "help message for flagname")
```

demo2: 如果你喜欢，你可以使用Var()函数将标志绑定到一个变量
```asciidoc
var flagvar int
func init() {
    flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

demo3: 您可以创建满足Value接口（使用指针接收器）的自定义标志，并将它们耦合以标记解析
```asciidoc
flag.Var(&flagVal, "name", "help message for flagname")
```
对于这样的标志，默认值只是变量的初始值

所有标志定义后，调用如下方法来将命令行解析为定义的标志
```asciidoc
flag.Parse()
```

标志可以直接使用。如果你自己使用标志，它们都是指针；如果你绑定到变量，它们就是值。
```asciidoc
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```

解析后，标志后面的参数可用作切片flag.Args()或单独作为flag.Arg(i)。参数从0到flag.NArg()-1索引。

命令行标志语法: 可以使用一个或两个减号，它们是等价的。由于命令的含义，最后一种形式不允许用于布尔标志，如果有一个名为0或false等类似的文件，它将会改变。你必须使用 -flag = false来关闭布尔标志。
```asciidoc
-flag
-flag=x
-flag x  //仅限非布尔标志
```

标志解析在第一个非标志参数（"-"是非标志参数）之前停止，或者在终止符"-"之后停止。
```
cmd -x *
```

整数标志接受1234，0644，0x1234并且可能是负数。布尔标志可能是：
```asciidoc
1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
```

持续时间标志接受对time.ParseDuration有效的任何输入。

默认的一组命令行标志由顶层函数控制。FlagSet类型允许您定义独立的标志集，例如在命令行界面中实现子命令。FlagSet的方法类似于命令行标志集的顶层函数。

### 示例
```asciidoc
// 这些示例演示了对标志包的更复杂的使用。
package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

// 示例1：名为“species”的单个字符串标志，默认值为“gopher”。
var species = flag.String("species", "gopher", "the species we are studying")

// 示例2：共享变量的两个标志，因此我们可以使用速记。
// 初始化的顺序是未定义的，因此请确保两者都使用
// 相同的默认值。 必须使用init函数设置它们。
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

// 示例3：用户定义的标志类型，持续时间片。
type interval []time.Duration

// String是格式化标志值的方法，是flag.Value接口的一部分。
// String方法的输出将用于诊断。
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set是设置标志值的方法，flag.Value接口的一部分。
// Set的参数是要解析以设置标志的字符串。
// 这是一个以逗号分隔的列表，因此我们将其拆分。
func (i *interval) Set(value string) error {
	// 如果我们想允许多次设置标志，
	// 累积值，我们将删除此if语句。
	// 这将允许诸如此类的用法
	//	-deltaT 10s -deltaT 15s
	// 和其他组合。
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

// 定义一个标志来累积持续时间。 因为它有特殊的类型，
// 我们需要使用Var函数，因此在期间创建标志
// init。

var intervalFlag interval

func init() {
	// 将命令行标志绑定到intervalFlag变量和
	// 设置用法消息。
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
	// 所有有趣的部分都是上面声明的变量，但是
	// 要使标志包能够看到那里定义的标志，就必须这样做
	// 执行，通常在main（不是init！）的开头执行：
	//	flag.Parse()
	// 我们不在这里运行它，因为这不是主要功能
	//测试套件已经解析了标志。
}
```

### 变量
CommandLine是从os.Args解析的默认命令行标志集。顶级函数（如BoolVar, Arg等）是CommandLine方法的包装器。
```asciidoc
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
```

如果调用了-help或-h标志但没有定义这样的标志，ErrHelp是返回的错误
```asciidoc
var ErrHelp = erros.New("flag: help requested")
```

用法向标准错误输出记录所有定义的命令行标志的用法消息。它在解析标志时发生错误时被调用。该函数是一个可以更改为指向自定义函数的变量。默认情况下，它打印一个简单的标题并调用PrintDefaults；有关输出格式以及如何控制它的详细信息，请参阅PrintDefaultes的文档。
```
var Usage = func() {
    fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
    PrintDefaults()
}
```







