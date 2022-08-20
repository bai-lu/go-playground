package main

import (
	"fmt"

	"example.com/m/v2/tmp"
)

const NewName = "run.wu"

type Person interface {
	getName() string
	setName(name string)
}

const (
	_          = iota // 通过赋予空白标识符来忽略第一个值
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type Male struct {
	Name string
}

func (m Male) getName() string {
	return m.Name
}

func (m *Male) setName(name string) {
	m.Name = name
}

// 字节的单位转换 保留两位小数
func formatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2fB", float64(fileSize))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/KB)
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/MB)
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/GB)
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/TB)
	} else {
		return fmt.Sprintf("%.2fPB", float64(fileSize)/PB)
	}
}

func main() {
	var p1 Person = &Male{} // error
	p1.getName()
	p1.setName(NewName)

	var p2 Person = &Male{}
	p2.getName()
	p2.setName(NewName)
	println(p2)
	tmp.F3()
}
