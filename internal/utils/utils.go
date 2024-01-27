package utils

import (
	"fmt"
	"local-build/internal/lblog"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// generate ID string
// format: yyyyMMddHHmmss + 三位数字
func GenerateIdStr() string {
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

// generate ID
// format: yyyyMMddHHmmss + 三位数字
func GenerateId() int64 {
	result := GenerateIdStr()
	number, _ := strconv.ParseInt(result, 10, 64)

	return number
}

// get home path
func GetHomePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	lblog.Info("exPath: %s", exPath)
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	lblog.Info("home path: %s", realPath)
	return realPath
}
