package key_generators

import (
	"gobin/database"
	"gobin/model"
	"math"
	"math/rand"
	"time"
)

const KeySpace = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const Length = 10

func GenerateRandomKey() string {
	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0;i < Length;i++ {
		result += string(KeySpace[rand.Intn(len(KeySpace))])
	}
	return result
}

func GetRandomKeySecurity() float64 {
	var result []model.PasteModel
	database.DB.Find(&result)
	pastes := len(result)
	combinations := math.Pow(float64(len(KeySpace)), Length)
	if pastes == 0 {
		return combinations
	}
	return combinations / float64(pastes)
}

