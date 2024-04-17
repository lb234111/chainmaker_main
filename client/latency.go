// coding:utf-8
// 延迟测试
package client

import (
	"chainmaker/chainmaker"
	"fmt"
	"strconv"
	"time"
)

// 延迟测试
// @param clients: 客户端
// @param num: 交易总数
// @param group_num: 分片数量
func Latency(clients []*chainmaker.Client, num int, group_num int) {
	var timestamp int64
	var time1, time2, t int64
	var max_time1, max_time2 int64 = 0, 0
	timestamp = 0
	for i := 0; i < num; i++ {
		t = time.Now().UnixNano()
		_, tx1, _ := Txmaker(0, group_num)
		tx1.Id = ""
		InvokefromTx(clients, tx1, true)
		t = time.Now().UnixNano() - t
		if t > max_time1 {
			max_time1 = t
		}
		timestamp += t
	}
	time1 = timestamp / 1000000 / int64(num)
	max_time1 = max_time1 / 1000000
	timestamp = 0
	for i := 0; i < num; i++ {
		t = time.Now().UnixNano()
		_, tx1, tx2 := Txmaker(100, group_num)
		tx1.Id = ""
		tx2.Id = ""
		InvokefromTx(clients, tx1, true)
		InvokefromTx(clients, tx2, true)
		t = time.Now().UnixNano() - t
		if t > max_time2 {
			max_time2 = t
		}
		timestamp += t
	}
	time2 = timestamp / 1000000 / int64(num)
	max_time2 = max_time2 / 1000000
	fmt.Println("非跨片延迟平均" + strconv.Itoa(int(time1)) + "毫秒,跨片延迟平均" + strconv.Itoa(int(time2)) + "毫秒")
	fmt.Println("非跨片延迟最大" + strconv.Itoa(int(max_time1)) + "毫秒,跨片延迟最大" + strconv.Itoa(int(max_time2)) + "毫秒")
}
