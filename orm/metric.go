package orm

import (
	"time"
)

// 巡检结果基本单位
type Metric struct {
	UUID      string `gorm:"primarykey"` // 经过 hash 散列生成 uuid
	CreatedAt time.Time
	ResultID  uint //  结果表ID, 作为外键

	Type     string // os | cert | lvs | l7 | horus
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

func NewMetric() *Metric {
	return &Metric{}
}

func (metric *Metric) GetUUID() string {
	return metric.UUID
}

func (metric *Metric) TimeRange() (int64, int64, error) {
	return metric.StartTimeStamp, metric.EndTimeStamp, nil
}

func (metric *Metric) String() string {
	return ""
}

func (metric *Metric) IsError() bool {
	return true
}

func (metric *Metric) IsWarn() bool {
	return true
}

func (metric *Metric) IsNormal() bool {
	return true
}

func (metric *Metric) Max() float64 {
	return 0
}

func (metric *Metric) Min() float64 {
	return 0
}

func (metric *Metric) Avg() float64 {
	return 0
}
