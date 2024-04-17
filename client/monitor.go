// coding:utf-8
// 交易监测
package client

import (
	"chainmaker/chainmaker"
	"sync"
	"time"
)

// 监控线程结束时间
const timeout int64 = 30

// 交易监测
// @param clients: 客户端
// @param cnt: 交易总数
// @param group_num: 分片数量
// @param tx_map: 等待交易区
// @param tx_num1: 片内交易数量
// @param tx_num2: 跨片交易数量
// @param chain_num: 各分片交易数量
// @param wg: 多线程信号量
func Monitor(clients []*chainmaker.Client, cnt int, group_num int, tx_map *map[string]Tx, tx_num1 *int, tx_num2 *int, chain_num *[16]int, wg *sync.WaitGroup) {
	defer (*wg).Done()
	//fmt.Printf("进入monitor\n")
	var block_start, block_end [16]int64
	for i := 0; i < group_num; i++ {
		block_start[i] = clients[i].GetCurrentBlockHeight()
	}
	new_time := time.Now().Unix()
	for {
		if (*tx_num1)+(*tx_num2) >= cnt || (time.Now().Unix()-new_time > timeout) {
			break
		}
		for i := 0; i < group_num; i++ {
			block_end[i] = clients[i].GetCurrentBlockHeight()
			if block_end[i] > block_start[i] {
				new_time = time.Now().Unix()
				for j := block_start[i] + 1; j <= block_end[i]; j++ {
					id_list := clients[i].GetBlockTxIdByHeight(j)
					for k := 0; k < len(id_list); k++ {
						tx, exist := (*tx_map)[id_list[k]]
						if exist {
							delete(*tx_map, id_list[k])
							if tx.Method == "" {
								//
								(*tx_num2)++
								(*chain_num)[i]++
							} else {
								//*queue = append(*queue, tx)
								InvokefromTx(clients, tx, false)
								(*tx_map)[tx.Id] = Tx{id_list[k], "", 0, "", "", 0}
								(*chain_num)[i]++
							}
						} else {
							//
							(*tx_num1)++
							(*chain_num)[i]++
						}
					}
				}
				block_start[i] = block_end[i]
			}
		}
	}
	//fmt.Printf("退出monitor \n")
}
