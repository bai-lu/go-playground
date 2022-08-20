// 巡检结果基本单元, 包含的信息可以复现
package define

type Metric interface {
	TimeRange() (uint, uint, error)
	String() string
	IsError() bool
	IsWarn() bool
	IsNormal() bool
}
