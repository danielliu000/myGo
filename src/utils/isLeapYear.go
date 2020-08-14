//2. 判断是否为闰年

package utils

func IsLeapYear(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	} else {
		return false
	}
}
