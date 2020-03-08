package utils

import (
	"time"
)

func CurrentTime() (int) {
	nowTimeStamp := time.Now().Unix()
	return int(nowTimeStamp)
}