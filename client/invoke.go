// coding:utf-8
// 交易执行
package client

import (
	"chainmaker/chainmaker"
	"sync"
)

// 多线程发送所有队列交易
// @param clients: 客户端
// @param group_num: 分片数量
// @param cnt: 交易总数
// @param queue: 交易队列
// @param tx_map: 等待交易区
func InvokeAllTx(clients []*chainmaker.Client, group_num int, client_num int, cnt int, queue [MAXQUEUE][]Tx, tx_map map[string]Tx) {
	var wg sync.WaitGroup
	var timestamp = [2][]int64{}    // 记录开始结束时间
	var tx_num1, tx_num2 int = 0, 0 // 记录非跨链和跨链交易数量
	var chain_num [16]int           // 记录各链执行交易数量
	var start_height [16]int64
	for i := 0; i < group_num; i++ {
		clients[i].CallStore("", "addr1", 10000, true)
		timestamp[0] = append(timestamp[0], clients[i].GetCurrentBlockTime())
		start_height[i] = clients[i].GetCurrentBlockHeight()
	}
	wg.Add(1)
	go Monitor(clients, cnt, group_num, &tx_map, &tx_num1, &tx_num2, &chain_num, &wg)
	for i := 0; i < group_num*client_num; i++ {
		wg.Add(1)
		go InvokefromQueue(clients, &queue[i], &wg)
	}
	wg.Wait()
	for i := 0; i < group_num; i++ {
		timestamp[1] = append(timestamp[1], clients[i].GetCurrentBlockTime())
	}
	Calculate(clients, group_num, timestamp, tx_num1, tx_num2, chain_num, start_height)
}

// 发送队列中所有交易
func InvokefromQueue(clients []*chainmaker.Client, q *[]Tx, wg *sync.WaitGroup) {
	defer (*wg).Done()
	//fmt.Printf("进入invoke\n")
	for {
		if len(*q) < 1 {
			//fmt.Printf("退出invoke\n")
			return
		}
		tx := (*q)[0]
		(*q) = (*q)[1:]
		InvokefromTx(clients, tx, false)
	}
	//fmt.Printf("退出invoke\n")
}

// 发送单笔交易
func InvokefromTx(clients []*chainmaker.Client, tx1 Tx, withSyncResult bool) {
	switch tx1.Method {
	case "transfer":
		clients[tx1.Num].CallStore(tx1.Id, tx1.From, tx1.Value, withSyncResult)
		//clients[tx1.Num].CallTransfer(tx1.Id, tx1.From, tx1.To, tx1.Value, withSyncResult)
		//clients[tx1.Num].CallDonothing(tx1.Id, withSyncResult)
	case "store":
		//clients[tx1.Num].CallStore(tx1.Id, tx1.From, tx1.Value, withSyncResult)
		clients[tx1.Num].CallDonothing(tx1.Id, withSyncResult)
	case "withdraw":
		//clients[tx1.Num].CallWithdraw(tx1.Id, tx1.From, tx1.Value, withSyncResult)
		clients[tx1.Num].CallDonothing(tx1.Id, withSyncResult)
	default:
	}
}
