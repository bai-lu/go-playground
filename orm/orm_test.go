package orm

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCreateTable(t *testing.T) {
	DB.AutoMigrate(&Metric{}, &Result{})
}

func TestCreateResult(t *testing.T) {
	DB.AutoMigrate(&Result{})
}

func TestInsertMetric(t *testing.T) {
	values, _ := json.Marshal(initData())
	// t.Log(len(values))
	// t.Log(len(md5.Sum(values)))
	DB.Create(&Metric{
		UUID:           fmt.Sprintf("%x", md5.Sum(values)),
		CreatedAt:      time.Now(),
		Type:           "os",
		Endpoint:       "c3-mc-sre00.bj",
		Metric:         "cpu.busy",
		Label:          "",
		Remark:         "CPU负载",
		Values:         values,
		Step:           60,
		Threshold:      90,
		StartTimeStamp: time.Now().Unix() - 86400*3,
		EndTimeStamp:   time.Now().Unix(),
	})

}

func TestInsertResult(t *testing.T) {
	values, _ := json.Marshal(initData())
	metric := Metric{
		UUID:           fmt.Sprintf("%x", md5.Sum(values)),
		CreatedAt:      time.Now(),
		Type:           "os",
		Endpoint:       "c3-mc-sre00.bj",
		Metric:         "cpu.busy",
		Label:          "",
		Remark:         "CPU负载",
		Values:         values,
		Step:           60,
		Threshold:      90,
		StartTimeStamp: time.Now().Unix() - 86400*3,
		EndTimeStamp:   time.Now().Unix(),
	}
	DB.Create(&Result{
		PID:           0,
		RelatedTags:   "test",
		StartTime:     time.Now(),
		CompletedTime: time.Now(),
		Metrics:       []Metric{metric},
	})
}

func TestSelectInspectResult(t *testing.T) {
	var results []Result
	// DB.Find(&results)
	DB.Preload("Metrics").Find(&results)
	for _, result := range results {
		t.Log(result)
	}
}
