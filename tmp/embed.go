package tmp

type Reader interface { //定义读取的接口类型
	Read(p []byte) (n int, err error) //定义函数   传入[]byte类型  返回一个整型和err
}

type Writer interface { //定义写入的接口类型
	Write(p []byte) (n int, err error) //定义函数   传入[]byte类型  返回一个整型和err
}

// ReadWriter 接口结合了 Reader 接口 和 Writer 接口
type ReadWriter interface {
	Reader
	Writer
}

// // ReadWriter 存储了指向 Reader 和 Writer 的指针。
// // 它实现了 io.ReadWriter。
type ReadWriter1 struct {
	Reader // *bufio.Reader
	Writer // *bufio.Writer
}

func (rw *ReadWriter1) Read(p []byte) (n int, err error) {
	return rw.Reader.Read(p)
}
