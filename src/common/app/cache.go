package app

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func GetCache() []string {
	filepath := "./cache.txt"
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	res := strings.Split(string(content), ",")
	return res
}

func GetCacheByKey(key int) int {

	filepath := GetBasePath() + "conf/cache.txt"
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	res := strings.Split(string(content), ",")
	if len(res) <= key {
		return 0
	} else {
		//fmt.Println(res)
		//fmt.Println(key)
		intValue, err := strconv.Atoi(res[key])
		if err != nil {
			intValue = 0
		}
		return intValue
	}
}
func SetCache(key int, value string) {
	filepath := GetBasePath() + "conf/cache.txt"
	var newStrArr []string
	strArr := GetCache()
	newStrArr = strArr
	if key >= len(strArr) {
		newStrArr = make([]string, key+1, key+1)
		copy(newStrArr, strArr)
		// fmt.Println("ddd")
	}
	fmt.Println("len:")
	fmt.Println(len(strArr))
	fmt.Println("bbb")
	fmt.Println(key + 1)
	newStrArr[key] = value
	fmt.Println("ccc")

	resStr := strings.Join(newStrArr, ",")
	err := ioutil.WriteFile(filepath, []byte(resStr), 0644)
	if err != nil {
		panic(err)
	}
}
