package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func f2str(f float64) string {
	str := strconv.FormatFloat(f, 'f', -1, 64)
	return str
}
func Between(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}
func log_e(e interface{}) {
	f, err := os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// 完成后延迟关闭，而不是习惯!
	defer f.Close()
	//设置日志输出到 f
	log.SetOutput(f)
	//测试用例
	log.Println(e)

}
func log_keep(e interface{}) {
	f, err := os.OpenFile("keep.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// 完成后延迟关闭，而不是习惯!
	defer f.Close()
	//设置日志输出到 f
	log.SetOutput(f)
	//测试用例
	log.Println(e)

}

func Decimal8(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.8f", value), 64)
	return value
}
func Decimal2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
func Decimal5(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", value), 64)
	return value
}
func PathExists(path string) bool {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func md5str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {

}
