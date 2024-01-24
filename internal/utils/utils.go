package utils

import (
	"fmt"
	"log"
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
	log.Printf("exPath: %v\n", exPath)
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	log.Printf("home path: %v\n", realPath)
	return realPath
}

// print error
func Error(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
