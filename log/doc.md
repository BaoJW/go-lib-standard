#log

## 概述
Log包实现了一个简单的日志包。它定义了一个类型，记录器，用于格式化输出的方法。
它还有一个预定义的"standard"记录器，可以通过帮助函数Print[f|ln], Fatal[f|ln], and Panic[f|ln]访问
，比手动创建记录器更易于使用。该记录器写入标准错误并打印每条记录的消息的日期和时间。每条日志消息都在一个单独的行上输出：
如果正在打印的消息不以换行符结尾，则记录器将添加一条。写入日志消息后，Fatal函数调用os.Exit(1);写入日志消息后，Panic函数调用panic。

## 常量
这些标志定义了哪些文本要加入由Logger生成的每个日志条目。
```asciidoc
const (
        // Bit 或'ed 一起控制什么打印。
        // 无法控制它们出现的顺序（列出的顺序）
        // 这里）或他们提出的格式（如注释中所述）。
        // 仅当Llongfile或Lshortfile时，前缀后跟冒号
        // 已指定。
        // 例如，标志Ldate|Ltime（或LstdFlags）产生，
        //	2009/01/23 01:23:23 信息
        // 而标志Ldate | Ltime | Lmicroseconds | Llongfile生成，
        //	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: 信息
        Ldate         = 1 << iota     // 当地时区的日期：2009/01/23
        Ltime                         // 在当地时区的时间：01:23:23
        Lmicroseconds                 // 微秒分辨率：01：23：23.123123。 假设Ltime。
        Llongfile                     // 完整的文件名和行号：/a/b/c/d.go:23
        Lshortfile                    // 最终文件名元素和行号：d.go:23。 覆盖Llongfile
        LUTC                          // 如果设置了日期或时间，请使用UTC而不是本地时区
        LstdFlags     = Ldate | Ltime // 标准记录器的初始值
)

```

# 日志系统 | log/syslog

## 概述
软件包系统日志为系统日志服务提供了一个简单的界面。他可以使用UNIX套接字，UDP或TCP将消息发送到
syslog守护进程。

只需要调用一次Dial即可。在写入失败时，系统日志客户端将尝试重新连接到服务器并重新写入。

syslog软件包被冻结，并且不能接受新的函数。一些外部软件包提供更多功能。参见：
```asciidoc
https://godoc.org/?q=syslog
```

