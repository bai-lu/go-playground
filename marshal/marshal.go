package marshal

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

func initData() [1000]float64 {
	var data [1000]float64
	for i := 0; i < len(data); i++ {
		data[i] = rand.Float64()
	}
	return data
}

// float64 to string 精度丢失
func sprintfMarshal(data [1000]float64) string {
	var s string
	for i := 0; i < len(data); i++ {
		s += fmt.Sprintf("%f,", data[i])
	}
	return s
}

func stringBuilderfMarshal(data [1000]float64) string {
	var s []string = make([]string, len(data))
	for i := 0; i < len(data); i++ {
		s[i] = fmt.Sprintf("%f,", data[i])
	}
	res := strings.Join(s, ",")
	return res
}

func formatMarshal(data [1000]float64) string {
	var s string
	for i := 0; i < len(data); i++ {
		s += strconv.FormatFloat(data[i], 'f', 0, 64) + ","
	}
	return s
}

func ItoaMarshal(data [1000]float64) string {
	var s string
	for i := 0; i < len(data); i++ {
		s += strconv.Itoa(int(data[i])) + ","
	}
	return s
}

func jsonMarshal(data [1000]float64) string {
	res, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func gjsonUnMarshal(s string) [1000]float64 {
	var data [1000]float64
	result := gjson.Parse(s).Array()

	for i := 0; i < len(result); i++ {
		data[i] = result[i].Float()
	}

	return data
}

func stringUnMarshal(s string) [1000]float64 {
	var data [1000]float64
	// s = strings.TrimSuffix(s, "]")
	// s = strings.TrimPrefix(s, "[")
	s = strings.Trim(s, "[]")
	dataList := strings.Split(s, ",")
	for i := 0; i < len(dataList); i++ {
		data[i], _ = strconv.ParseFloat(dataList[i], 64)
	}
	return data
}

func jsonUnMarshal(s string) [1000]float64 {
	var data [1000]float64
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		panic(err)
	}
	return data
}
