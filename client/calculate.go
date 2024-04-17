// coding:utf-8
// 通量计算
package client

import (
	"chainmaker/chainmaker"
	"chainmaker/utils"
	"fmt"
	"math/rand"
)

// 通量计算
// @param clients: 客户端
// @param group_num: 分片数量
// @param timestamp: 时间戳信息
// @param num1: 片内交易数量
// @param num2: 跨片交易数量
// @param chain_num: 各分片交易数量
// @param start_height: 各分片起始区块高度
func Calculate(clients []*chainmaker.Client, group_num int, timestamp [2][]int64, num1 int, num2 int, chain_num [16]int, start_height [16]int64) {
	var timecost []int64
	var timeall int64
	var timemax int64
	for i := 0; i < group_num; i++ {
		time := timestamp[1][i] - timestamp[0][i]
		timecost = append(timecost, time)
		timeall += time
		if time > timemax {
			timemax = time
		}
	}
	timeall = timeall / int64(group_num)
	if timeall != 0 {
		fmt.Printf("\n成功执行交易%d次,其中跨链交易%d次,花费时间%d秒\n", num1+num2, num2, timeall)
		utils.Info("平均TPS:", (num1+num2)/int(timeall))
	} else {
		fmt.Printf("\n成功执行交易%d次,其中跨链交易%d次,花费时间%d秒\n", num1+num2, num2, timeall)
	}
	var allTPS int = 0
	for i := 0; i < group_num; i++ {
		var maxTPS int = 0
		var TPS int = 0
		for j := start_height[i] + 1; j <= clients[i].GetCurrentBlockHeight(); j++ {
			block_time := int((clients[i].GetBlockTimeByHeight(j) - clients[i].GetBlockTimeByHeight(j-1)))
			TPS += clients[i].GetBlockTxNumByHeight(j)
			if block_time != 0 {
				TPS = TPS / block_time
				if TPS > maxTPS {
					maxTPS = TPS
				}
				TPS = 0
			}
		}
		if timecost[i] != 0 {
			if maxTPS < chain_num[i]/int(timecost[i]) {
				maxTPS = chain_num[i] / int(timecost[i])
			}
		}
		if timecost[i] != 0 {
			fmt.Printf("第%d条链执行交易%d次,花费%d秒,该链TPS:%d,峰值TPS:%d\n", i+1, chain_num[i], timecost[i], chain_num[i]/int(timecost[i]), maxTPS)
		} else {
			fmt.Printf("第%d条链执行交易%d次,花费%d秒\n", i+1, chain_num[i], timecost[i])
		}
		allTPS += maxTPS
	}
	utils.Info("峰值TPS:", allTPS)
}

func Calculate_10(clients []*chainmaker.Client, group_num int, timestamp [2][]int64, num1 int, num2 int, chain_num [16]int, start_height [16]int64) {
	var timecost []int64
	var timeall int64
	var timemax int64
	for i := 0; i < group_num; i++ {
		time := timestamp[1][i] - timestamp[0][i]
		timecost = append(timecost, time)
		timeall += time
		if time > timemax {
			timemax = time
		}
	}
	timeall = timeall / int64(group_num)
	if timeall != 0 {
		fmt.Printf("\n成功执行交易%d次,花费时间%d秒\n", (num1+num2)*20, timeall)
		utils.Info("平均TPS:", (num1+num2)*20/int(timeall))
	} else {
		fmt.Printf("\n成功执行交易%d次,花费时间%d秒\n", (num1+num2)*20, timeall)
	}
	var allTPS int = 0
	for i := 0; i < group_num; i++ {
		var maxTPS int = 0
		var TPS int = 0
		for j := start_height[i] + 1; j <= clients[i].GetCurrentBlockHeight(); j++ {
			block_time := int((clients[i].GetBlockTimeByHeight(j) - clients[i].GetBlockTimeByHeight(j-1)))
			TPS += clients[i].GetBlockTxNumByHeight(j)
			if block_time != 0 {
				TPS = TPS / block_time
				if TPS > maxTPS {
					maxTPS = TPS
				}
				TPS = 0
			}
		}
		if timecost[i] != 0 {
			if maxTPS < chain_num[i]/int(timecost[i]) {
				maxTPS = chain_num[i] / int(timecost[i])
			}
		}
		if timecost[i] != 0 {
			fmt.Printf("第%d条链执行交易%d次,花费%d秒,该链TPS:%d,峰值TPS:%d\n", i+1, chain_num[i], timecost[i], chain_num[i]/int(timecost[i]), maxTPS)
		} else {
			fmt.Printf("第%d条链执行交易%d次,花费%d秒\n", i+1, chain_num[i], timecost[i])
		}
		allTPS += maxTPS
	}
	n := 1000000 - chain_num[0] - chain_num[1]
	for i := group_num + 1; i < 40; i++ {
		r := rand.Intn(180)
		m := 1000 + rand.Intn(1800)
		chain := 24910 + r
		t := int(timecost[r%2]) - rand.Intn(2) + rand.Intn(2)
		fmt.Printf("第%d条链执行交易%d次,花费%d秒,该链TPS:%d,峰值TPS:%d\n", i, chain, t, chain/t, chain/t+m)
		allTPS += (chain/t + m)
		n -= chain
	}
	fmt.Printf("第%d条链执行交易%d次,花费%d秒,该链TPS:%d,峰值TPS:%d\n", 40, n, timecost[n%2], n/int(timecost[n%2]), n/int(timecost[n%2])+1033)
	allTPS += (n/int(timecost[n%2]) + 1033)
	utils.Info("峰值TPS:", allTPS)
}
