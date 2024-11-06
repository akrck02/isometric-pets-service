package utils

import "strconv"

// Convert int to string
//
// [param] num | int: number to convert
//
// [return] string: converted number
func Int2String(num int) string {
	return strconv.Itoa(num)
}
