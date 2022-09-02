package api

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type DataPoints [][2]float64

type Metric struct {
	Target     string
	DataPoints DataPoints
}

type MachineMetric struct {
	Endpoint  string  `json:"endpoint"`
	Metric    string  `json:"metric"`
	StartTime int64   `json:"from"`
	EndTime   int64   `json:"to"`
	Avg       float64 `json:"avg"`
}

// 1. statistic
// 2.

func NewMachineData() []MachineMetric {
	var MachineData []MachineMetric
	for i := 0; i < 100; i++ {
		// tmp := MachineMetric{
		// 	Endpoint:  fmt.Sprintf("c3-mc-sre%02d.bj", i),
		// 	Metric:    "cpu.busy",
		// 	StartTime: time.Now().Format("2006-01-02T15:04:05"),
		// 	EndTime:   time.Now().Format("2006-01-02T15:04:05"),
		// 	Avg:       i,
		// }
		tmp := MachineMetric{
			Endpoint:  fmt.Sprintf("c3-mc-sre%02d.bj", i),
			Metric:    "cpu.busy",
			StartTime: time.Now().UnixMilli(),
			EndTime:   time.Now().UnixMilli(),
			Avg:       float64(i),
		}
		MachineData = append(MachineData, tmp)

	}
	return MachineData
}

func NewMachineData1() []map[string]interface{} {
	var MachineData []map[string]interface{}
	for i := 0; i < 100; i++ {
		// tmp := MachineMetric{
		// 	Endpoint: "c3-mc-sre00.bj",
		// 	Metric:   "cpu.busy",
		// 	Time:     time.Now().Unix(),
		// 	Avg:      i,
		// }

		tmp := map[string]interface{}{
			fmt.Sprintf("c3-mc-sre%d.bj", i): i,
			"Time":                           time.Now().Unix(),
		}
		MachineData = append(MachineData, tmp)

	}
	return MachineData

}

func NewMetric1() Metric {
	return Metric{
		Target: "pps in",
		DataPoints: [][2]float64{
			{622, 1450754160000},
			{365, 1450754220000},
			{365, 1450754320000},
			{365, 1450754420000},
			{365, 1450754520000},
			{365, 1450754620000},
			{365, 1450754720000},
		},
	}
}

func LoadFeedJson() interface{} {
	var res interface{}
	fd, err := os.Open("api/machine.json")
	if err != nil {
		panic(err)
	}
	fileContext, err := io.ReadAll(fd)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(fileContext, &res)
	return res
}
