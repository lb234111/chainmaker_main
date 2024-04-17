// coding:utf-8
// 获取多个区块信息
package client

import (
	"chainmaker/chainmaker"
	"fmt"
)

// 获取多个区块信息
// @param clients: 客户端
// @param group_num: 分片数量
// @param num: 交易总数
// @param height: 起始区块高度
func GetTxCount(clients []*chainmaker.Client, group_num int, num int, height int) {
	time := clients[num-1].GetBlockTimeByHeight(int64(height - 1))
	for i := int64(height); i <= clients[num-1].GetCurrentBlockHeight(); i++ {
		list := clients[num-1].GetBlockTxIdByHeight(i)
		time_next := clients[num-1].GetBlockTimeByHeight(i)
		fmt.Println(len(list), time_next-time)
		time = time_next
	}
}
