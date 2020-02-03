package fmt

//个性化定制扫描仪

//fmt包接口说明

/*
ScanState 表示传递给定制扫描仪的扫描仪状态。扫描仪可以做一次一个符号扫描，或要求 ScanState 发现下一个由空格分隔的令牌。

type ScanState interface {
        // ReadRune从输入中读取下一个符文（Unicode代码点）。
        // 如果在Scanln，Fscanln或Sscanln期间调用，ReadRune()将会
        // 返回第一个'\n'或读取超出后返回EOF
        // 指定的宽度。
        ReadRune() (r rune, size int, err error)
        // UnreadRune导致下一次调用ReadRune返回相同的符文。
        UnreadRune() error
        // SkipSpace会跳过输入中的空格。 新线得到适当的处理
        // 正在进行的操作; 请参阅包文档
        // 了解更多信息。
        SkipSpace()
        // 如果skipSpace为true，Token会跳过输入中的空格，然后返回
        // 运行Unicode代码点c满足f(c)。 如果f为零，
        // !unicode.IsSpace(c)用于; 也就是说，令牌将保持非空格
        // 字符。 对于正在进行的操作，新线被适当地处理
        // 执行; 有关更多信息，请参阅包文档。
        // 返回的切片指向可能被覆盖的共享数据
        // 通过下一次调用Token，使用ScanState调用Scan功能
        // 作为输入，或当调用Scan方法返回时。
        Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
        // Width返回width选项的值以及是否已设置。
        // 该单位是Unicode代码点。
        Width() (wid int, ok bool)
        // 因为ReadRune是由接口实现的，所以Read永远不应该
        // 由扫描例程和有效实现调用
        // ScanState可以选择始终从Read返回错误。
        Read(buf []byte) (n int, err error)
}
*/

/*
Scanner 是由任何具有 Scan 方法的值实现的，Scan 方法扫描输入以表示值并将结果存储在接收器中，该接收器必须是有用的指针。扫描方法被称为 Scan，Scanf 或 Scanln 的任何参数来实现它。

type Scanner interface {
        Scan(state ScanState, verb rune) error
}
*/

/*
State 表示传递给自定义格式化程序的打印机状态。它提供对io.Writer 接口的访问，以及有关操作数格式说明符的标志和选项的信息。

type State interface {
        // Write是调用以发出要打印的格式化输出的函数。
        Write(b []byte) (n int, err error)
        // Width返回width选项的值以及是否已设置。
        Width() (wid int, ok bool)
        // Precision返回精度选项的值以及是否已设置。
        Precision() (prec int, ok bool)

        // 标志报告是否已设置标志c（字符）。
        Flag(c int) bool
}
*/
