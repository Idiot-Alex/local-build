package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// generate ID
// format: yyyyMMddHHmmss + 三位数字
func GenerateId() string {
	// 生成时间戳，格式为 yyyyMMddHHmmss
	t := time.Now()
	timeStamp := t.Format("20060102150405")

	// 生成 3 位随机数
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := rand.Intn(1000)

	// 拼接时间戳和随机数
	result := fmt.Sprintf("%s%03d", timeStamp, randomNum)
	
	return result
}
