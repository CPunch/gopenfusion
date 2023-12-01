package util

import "time"

func GetTime() uint64 {
	return uint64(time.Now().UnixMilli())
}
