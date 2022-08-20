package orm

import (
	"example.com/m/v2/define"
	"gorm.io/gorm"
)

// 巡检结果基本单位
type Metric struct {
	gorm.Model
	UUID            string // 经过 hash 散列生成 uuid
	InspectResultID uint   //  结果表ID, 作为外键

	Ttype    string // os | cert | lvs | l7 | horus
	Endpoint string
	Metric   string
	Label    string

	Remark string // 备注

	Values []byte
	// `gorm:"serializer:json"` // 巡检结果值
	Step uint // 巡检结果采样率

	Threshold      float64 // 因为策略是可以调整的, 因此巡检结果的阈值也需要保存
	StartTimeStamp int64   // 第一个超过阈值的点的时间戳, 用于前端绘图复现
	EndTimeStamp   int64   // 最后一个超过阈值的点的时间戳

}

// 确保Metric实现了define.Metric接口
var _ define.Metric = &Metric{}

func NewMetric() *Metric {
	return &Metric{}
}

func (Metric *Metric) GetUUID() string {
	return Metric.UUID
}

func (Metric *Metric) TimeRange() (uint, uint, error) {
	return 0, 0, nil
}

func (Metric *Metric) String() string {
	return ""
}

func (Metric *Metric) IsError() bool {
	return true
}

func (Metric *Metric) IsWarn() bool {
	return true
}

func (Metric *Metric) IsNormal() bool {
	return true
}

func (Metric *Metric) Max() float64 {
	return 0
}

func (Metric *Metric) Min() float64 {
	return 0
}

func (Metric *Metric) Avg() float64 {
	return 0
}
