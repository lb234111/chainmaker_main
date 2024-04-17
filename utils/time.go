// coding:utf-8
// 时间相关处理
package utils

import (
	"time"
)

// 获取当前时区
func GetTimeLocation() string {
	return "Asia/Shanghai"
}

// 获取当前时间"年-月-日 时:分:秒"
func GetTime() string {
	//设置时区
	var cstSh, _ = time.LoadLocation(GetTimeLocation())
	return time.Now().In(cstSh).Format("2006-01-02 15:04:05")
}

// 获取时间戳,以秒为单位
func GetTimestamp() int64 {
	//设置时区
	var cstSh, _ = time.LoadLocation(GetTimeLocation())
	return time.Now().In(cstSh).Unix()
}

// 获取时间戳,以毫秒为单位
func GetTimestampMs() int64 {
	//设置时区
	var cstSh, _ = time.LoadLocation(GetTimeLocation())
	return time.Now().In(cstSh).UnixMilli()
}

// 获取时间戳,以微秒为单位
func GetTimestampMicoro() int64 {
	//设置时区
	var cstSh, _ = time.LoadLocation(GetTimeLocation())
	// fmt.Println("=========")
	// // 微秒
	// fmt.Println(time.Now().UnixMicro())
	// // 毫秒
	// fmt.Println(time.Now().UnixMilli())
	// // return time.Now().In(cstSh).UnixMicro()
	// // 秒
	// fmt.Println(time.Now())
	return time.Now().In(cstSh).UnixMicro()
}

// 休眠一定时间
func Sleep(n int) {
	time.Sleep(time.Duration(n) * time.Second)
}
