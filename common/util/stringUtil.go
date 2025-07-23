package util

import (
	"strconv"
)

// StringToUint 将字符串转换为uint
func StringToUint(s string) uint {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint(i)
}
