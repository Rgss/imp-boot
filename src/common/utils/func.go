package utils

import (
	"encoding/json"
	"strconv"
	"crypto/md5"
	"io"
	"fmt"
)

func JsonNumberToInt(n json.Number) (int) {
	m, _ := strconv.ParseFloat(string(n), 64)
	return int(m)
}

func JsonNumber2Float64(n json.Number) (float64) {
	m, _ := strconv.ParseFloat(string(n), 64)
	return m
}

func Str2Int(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func Int2Str(i int) (string) {
	return strconv.Itoa(i)
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
