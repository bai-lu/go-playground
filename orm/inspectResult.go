package orm

import (
	"time"

	"example.com/m/v2/define"
	"gorm.io/gorm"
)

var DB = NewDB()

type InspectResult struct {
	gorm.Model // include created_time, updated_time, deleted_time

	PID           int       `gorm:"index;not null"` //  sre tree ID
	RelatedTags   string    //  关联标签, productionID 下可以包含多个 pdl tags, org, iam 节点, 巡检结果可能包含一个或多个, 因此也需要记录
	StartTime     time.Time // 巡检开始时间
	CompletedTime time.Time // 巡检完成时间

	Metrics []Metric // 巡检详细指标及其结果存储
}

var _ define.InspectResult = &InspectResult{}

// 写表记录

func (res *InspectResult) Save() {
	DB.Create(res)
}

// 读表记录
func (res *InspectResult) Load(PID int) {
	DB.Select(PID)
}

// 根据类型获取对应巡检的数量
func (res *InspectResult) Count(ttype string) int {
	return 0
}

// 获取本次巡检结果中错误级别巡检指标的数量
func (res *InspectResult) ErrorCount() int {
	count := 0
	for _, metric := range res.Metrics {
		if metric.IsError() {
			count++
		}
	}
	return count
}

// 获取本次巡检结果中错误级别巡检指标的数量
func (res *InspectResult) NormalCount() int {
	count := 0
	for _, metric := range res.Metrics {
		if metric.IsNormal() {
			count++
		}
	}
	return count
}

// 获取本次巡检结果中错误级别巡检指标的数量
func (res *InspectResult) WarnCount() int {
	count := 0
	for _, metric := range res.Metrics {
		if metric.IsWarn() {
			count++
		}
	}
	return count
}
