package main

import (
	"fmt"
	"time"
)

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Asia/Shanghai") //定义一个时区Location 指针
}
func basic() {

	//获取当前时间
	now := time.Now()
	//获取精确时间
	fmt.Println(
		now.Unix(),
		now.UnixMilli(),
		now.UnixMicro(),
		now.UnixNano(),
	)

	//计算其他时间
	fmt.Println(now.Day())        //  当月中的第几天
	fmt.Println(now.YearDay())    // 当年中的第几天
	fmt.Println(now.Month())      // 哪月
	fmt.Println(int(now.Month())) // 哪月 数字
	fmt.Println(now.Weekday())    //星期几
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//计算时间差
	time.Sleep(time.Second * 2)
	endTime := time.Since(now)
	fmt.Println(endTime)

}

func time_fmt() {
	const FmtTime = "2006-01-02 15:04:05.999"
	//构建时间差
	now := time.Now()
	dur := time.Duration(time.Hour * 8)
	end := now.Add(dur).Format(FmtTime) // 格式化时间 200612345
	fmt.Println(end)

	//从字符串转变为时间, 使用Parse(不带时区,不推荐) or ParseInLocation, 需要location指针
	if tLoc, err := time.ParseInLocation(FmtTime, "1999-02-03 3:43:43.324", loc); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tLoc)
	}
}

// 定时器
func ticker() {

}

func main() {
	basic()
}
