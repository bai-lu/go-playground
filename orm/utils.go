package orm

import (
	"math/rand"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		panic(err)
	}
	return db
}

func initData() [1000]float64 {
	var data [1000]float64
	for i := 0; i < len(data); i++ {
		data[i] = rand.Float64()
	}
	return data

}
