package orm

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCreateTable(t *testing.T) {
	DB.AutoMigrate(&Metric{})
}

func TestCreateInspectResult(t *testing.T) {
	DB.AutoMigrate(&InspectResult{})
}

func TestCreateUser(t *testing.T) {
	DB.AutoMigrate(&User{})
}

func TestInsertMetric(t *testing.T) {
	values, _ := json.Marshal(initData())
	// t.Log(len(values))
	// t.Log(len(md5.Sum(values)))
	DB.Create(&Metric{
		UUID:           fmt.Sprintf("%x", md5.Sum(values)),
		Ttype:          "os",
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

func TestInsertInspectResult(t *testing.T) {
	values, _ := json.Marshal(initData())
	metric := Metric{
		UUID:           fmt.Sprintf("%x", md5.Sum(values)),
		Ttype:          "os",
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
	DB.Create(&InspectResult{
		PID:           0,
		RelatedTags:   "test",
		StartTime:     time.Now(),
		CompletedTime: time.Now(),
		Metrics:       []Metric{metric},
	})
}

func TestSelectInspectResult(t *testing.T) {
	var results []InspectResult
	// DB.Find(&results)
	DB.Preload("Metrics").Find(&results)
	for _, result := range results {
		t.Log(result)
	}
}
