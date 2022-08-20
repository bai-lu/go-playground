package tmp

import (
	"fmt"
)

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

// 字节的单位转换 保留两位小数
func formatFileSize(fileSize int64) (size string) {
	fileSizeFloat64 := float64(fileSize)
	switch {
	case fileSize < int64(KB):
		return fmt.Sprintf("%.2fB", fileSizeFloat64)
	case fileSize < int64(MB):
		return fmt.Sprintf("%.2fKB", fileSizeFloat64/KB)
	case fileSize < int64(GB):
		return fmt.Sprintf("%.2fMB", fileSizeFloat64/MB)
	case fileSize < int64(TB):
		return fmt.Sprintf("%.2fGB", fileSizeFloat64/GB)
	case fileSize < int64(PB):
		return fmt.Sprintf("%.2fTB", fileSizeFloat64/TB)
	case fileSize < int64(EB):
		return fmt.Sprintf("%.2fPB", fileSizeFloat64/PB)
	default:
		return fmt.Sprintf("%.2fEB", fileSizeFloat64)
	}
}

func F(a int) {
	switch {
	case a > 10:
		fmt.Print("case1")
	case a > 20:
	default:
		fmt.Print("default")
	}
}
