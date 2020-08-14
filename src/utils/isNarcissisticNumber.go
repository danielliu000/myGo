package utils

import "math"

//5. 判断一个数是否为水仙花数
func IsNarcissisticNumber(num int) bool {
	if num >= 1000 || num <= 100 {
		return false
	} else {
		h := num / 100
		t := (num - h * 100) / 10
		d := num - t * 10 - h * 100
		rs :=
			math.Pow(float64(h),3) +
				math.Pow(float64(t),3) +
				math.Pow(float64(d),3)

		if num == int(rs) {
			return true
		} else {
			return false
		}

	}

}