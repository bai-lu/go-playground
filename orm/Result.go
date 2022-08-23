package orm

import (
	"time"

	"gorm.io/gorm"
)

var DB = NewDB()

type Result struct {
	gorm.Model // include created_time, updated_time, deleted_time

	PID           int       `gorm:"column:pid;index;not null"` //  sre tree ID
	RelatedTags   string    // 关联标签, productionID 下可以包含多个 pdl tags, org, iam 节点, 巡检结果可能包含一个或多个, 因此也需要记录
	StartTime     time.Time // 巡检开始时间
	CompletedTime time.Time // 巡检完成时间

	Metrics []Metric // 巡检详细指标及其结果存储
}

// 写表记录
func (res *Result) Save() {
	DB.Create(res)
}

// 读表记录
func (res *Result) Load(PID int) {
	DB.Select(PID)
}

// 根据类型获取对应巡检的数量
func (res *Result) Count(Type string) int {
	return 0
}

// 获取本次巡检结果中错误级别巡检指标的数量
func (res *Result) ErrorCount() int {
	count := 0
	for _, metric := range res.Metrics {
		if metric.IsError() {
			count++
		}
	}
	return count
}

// 获取本次巡检结果中错误级别巡检指标的数量
func (res *Result) NormalCount() int {
	count := 0
	for _, metric := range res.Metrics {
		if metric.IsNormal() {
			count++
		}
	}
	return count
}

// 获取本次巡检结果中错误级别巡检指标的数量
func (res *Result) WarnCount() int {
	count := 0
	for _, metric := range res.Metrics {
		if metric.IsWarn() {
			count++
		}
	}
	return count
}
