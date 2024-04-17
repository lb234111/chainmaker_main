// coding:utf-8
// 负载均衡
package client

import (
	"chainmaker/utils"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// 负载均衡
// @param queue: 交易队列
// @param group_num: 分片数量
// @param client_num: 分片客户端数量
// @param tx_map: 第二阶段跨片跨片
func BalanceTxs(queue [MAXQUEUE][]Tx, group_num int, client_num int, tx_map map[string]Tx) ([MAXQUEUE][]Tx, map[string]Tx) {
	var account_tx [16][16]int
	var account_group [16]int
	type account struct {
		index  int
		tx_num int
	}
	var account_tx_num []account
	for i := 0; i < group_num*client_num; i++ {
		for j := 0; j < len(queue[i]); j++ {
			if queue[i][j].Method == "transfer" {
				From, _ := strconv.Atoi(string(queue[i][j].From[4]))
				account_from := queue[i][j].Num*2 + From - 1
				to, _ := strconv.Atoi(string(queue[i][j].To[4]))
				account_to := queue[i][j].Num*2 + to - 1
				account_tx[account_from][account_to]++
			} else {
				From, _ := strconv.Atoi(string(queue[i][j].From[4]))
				account_from := queue[i][j].Num*2 + From - 1
				To, _ := strconv.Atoi(string(tx_map[queue[i][j].Id].From[4]))
				account_to := tx_map[queue[i][j].Id].Num*2 + To - 1
				account_tx[account_from][account_to]++
			}
		}
	}
	var chain_num [MAXQUEUE][2]int
	for i := 0; i < group_num*client_num; i++ {
		for j := 0; j < len(queue[i]); j++ {
			if queue[i][j].Method == "transfer" {
				chain_num[queue[i][j].Num][0]++
			} else {
				chain_num[queue[i][j].Num][0]++
				chain_num[tx_map[queue[i][j].Id].Num][1]++
			}

		}
	}
	for i := 0; i < group_num; i++ {
		fmt.Println("调整前第", i, "条链有", chain_num[i][0], "片内笔交易,", chain_num[i][1], "笔跨片交易,预计该链需要等待", chain_num[i][0]+chain_num[i][1]*2, "笔交易执行时间")
	}

	for i := 0; i < group_num*2; i++ {
		//fmt.Println(account_tx[i])
		account_tx_num = append(account_tx_num, account{i, AccountTxs(i, account_tx)})
	}
	var max int = 0
	for i := 0; i < group_num; i++ {
		if twoAccountTxs(i*2, i*2+1, account_tx) > max {
			max = twoAccountTxs(i*2, i*2+1, account_tx)
		}
	}
	utils.Info("等待最久的时间的链需要等待", max, "笔交易时间")

	sort.SliceStable(account_tx_num, func(i, j int) bool {
		return account_tx_num[i].tx_num > account_tx_num[j].tx_num
	})

	max = 0
	for i := 0; i < group_num; i++ {
		var account1, account2 int
		var j int
		for j = 0; account_group[account_tx_num[j].index] != 0; j++ {
		}
		account1 = account_tx_num[j].index
		account_group[account1] = i + 1
		min := 100000
		for j = 0; j < 2*group_num; j++ {
			if account_group[j] == 0 && twoAccountTxs(account1, j, account_tx) < min {
				account2 = j
				min = twoAccountTxs(account1, j, account_tx)
			}
		}
		account_group[account2] = i + 1
		if twoAccountTxs(account1, account2, account_tx) > max {
			max = twoAccountTxs(account1, account2, account_tx)
		}
	}
	fmt.Println("调整分片", account_group)
	queue, tx_map = ChangeTxs(account_tx, account_group, group_num, client_num, queue, tx_map)

	for i := 0; i < group_num*client_num; i++ {
		chain_num[i][0] = 0
		chain_num[i][1] = 0
	}
	for i := 0; i < group_num*client_num; i++ {
		for j := 0; j < len(queue[i]); j++ {
			if queue[i][j].Method == "transfer" {
				chain_num[queue[i][j].Num][0]++
			} else {
				chain_num[queue[i][j].Num][0]++
				chain_num[tx_map[queue[i][j].Id].Num][1]++
			}

		}
	}
	for i := 0; i < group_num; i++ {
		fmt.Println("调整后第", i, "条链有", chain_num[i][0], "片内笔交易,", chain_num[i][1], "笔跨片交易,预计该链需要等待", chain_num[i][0]+chain_num[i][1]*2, "笔交易执行时间")
	}
	utils.Info("等待最久的时间的链需要等待", max, "笔交易时间")
	return queue, tx_map
}

// 替换为负载均衡后交易
func ChangeTxs(account_tx [16][16]int, account_group [16]int, group_num int, client_num int, queue [MAXQUEUE][]Tx, tx_map map[string]Tx) ([MAXQUEUE][]Tx, map[string]Tx) {

	txid, err := ioutil.ReadFile("Txid.txt")
	if err != nil {
		fmt.Println(err)
	}
	new_txid, err = strconv.Atoi(strings.Trim(string(txid), "\n"))
	if err != nil {
		fmt.Println(err)
	}

	var queue_o [MAXQUEUE][]Tx
	tx_map_o := make(map[string]Tx, 100000)

	for i := 0; i < 2*group_num; i++ {
		for j := 0; j < 2*group_num; j++ {
			if account_group[i] == account_group[j] {
				for k := 0; k < account_tx[i][j]; k++ {
					var tx Tx
					if i >= j {
						tx = Tx{strconv.Itoa(new_txid), "transfer", account_group[i] - 1, "addr1", "addr2", 0}
					} else {
						tx = Tx{strconv.Itoa(new_txid), "transfer", account_group[i] - 1, "addr2", "addr1", 0}
					}
					new_txid++
					client := rand.Intn(client_num) * group_num
					queue_o[client+account_group[i]-1] = append(queue_o[client+account_group[i]-1], tx)
				}
			} else {
				for k := 0; k < account_tx[i][j]; k++ {
					var tx1, tx2 Tx
					if i >= j {
						tx1 = Tx{strconv.Itoa(new_txid), "withdraw", account_group[i] - 1, "addr1", "", 0}
						tx2 = Tx{strconv.Itoa(new_txid + 1), "store", account_group[j] - 1, "addr2", "", 0}
					} else {
						tx1 = Tx{strconv.Itoa(new_txid), "withdraw", account_group[i] - 1, "addr2", "", 0}
						tx2 = Tx{strconv.Itoa(new_txid + 1), "store", account_group[j] - 1, "addr1", "", 0}
					}
					client := rand.Intn(client_num) * group_num
					queue_o[client+account_group[i]-1] = append(queue_o[client+account_group[i]-1], tx1)
					tx_map_o[strconv.Itoa(new_txid)] = tx2
					new_txid += 2
				}
			}
		}
	}
	data := []byte(strconv.Itoa(new_txid + 1))
	ioutil.WriteFile("Txid.txt", data, 0664)
	return queue_o, tx_map_o
}

func AccountTxs(num int, account_tx [16][16]int) int {
	tx_num := 0
	for i := 0; i < 16; i++ {
		tx_num += account_tx[num][i]
	}
	for i := 0; i < 16; i++ {
		tx_num += account_tx[i][num] * 2
	}
	return tx_num
}

func twoAccountTxs(num1 int, num2 int, account_tx [16][16]int) int {
	return AccountTxs(num1, account_tx) + AccountTxs(num2, account_tx) - 2*(account_tx[num1][num2]+account_tx[num2][num1])
}
