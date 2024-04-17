// coding:utf-8
// 主客户端程序
package main

import (
	"bufio"
	"chainmaker/chainmaker"
	"chainmaker/client"
	"chainmaker/utils"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var group_num int = 4
var client_num int = 40

func args_parse() {
	flag.IntVar(&group_num, "group", 2, "分片数量")
	flag.IntVar(&client_num, "client", 20, "每分片客户端数量")
	flag.Parse()
}

func main() {
	//解析参数
	args_parse()
	//生成随机数
	rand.Seed(time.Now().Unix())

	var cnt, multi_cnt int                       // 交易总数,跨片交易占比
	var queue [client.MAXQUEUE][]client.Tx       // 片内交易和第一阶段跨片交易
	tx_map := make(map[string]client.Tx, 100000) // 第二阶段跨片交易

	utils.Print("分片原型系统ChainMaker测试开始")

	// 新建长安链链客户端
	var clients = []*chainmaker.Client{}
	for i := 0; i < client_num*group_num; i++ {
		path := "./config_files/sdkconfigs/chain" + strconv.Itoa(i%group_num+1) + "_sdkconfig1.yml" // 证书路径
		new_client, err := chainmaker.NewChainmakerClient(path)
		if err != nil {
			info := utils.InfoError("newclient", err)
			panic(info)
		}
		clients = append(clients, new_client)
		utils.Info("新建长安链客户端", i+1)
	}

	for {
		fmt.Print(">>>")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = strings.Trim(input, "\n")
		cmd := strings.Split(input, " ")

		if cmd[0] == "q" || cmd[0] == "quit" {
			break
		} else if cmd[0] == "getBalance" {
			if len(cmd) != 3 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			clients[num-1].CallGet("", cmd[2], true)
		} else if cmd[0] == "store" {
			if len(cmd) != 4 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			value, _ := strconv.Atoi(cmd[3])
			clients[num-1].CallStore("", cmd[2], value, true)
		} else if cmd[0] == "withdraw" {
			if len(cmd) != 4 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			value, _ := strconv.Atoi(cmd[3])
			clients[num-1].CallWithdraw("", cmd[2], value, true)
		} else if cmd[0] == "transfer" {
			if len(cmd) != 5 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			value, _ := strconv.Atoi(cmd[4])
			clients[num-1].CallTransfer("", cmd[2], cmd[3], value, true)
		} else if cmd[0] == "twowaytransfer" {
			if len(cmd) != 6 {
				fmt.Println("参数数量错误!")
				continue
			}
			num_from, _ := strconv.Atoi(cmd[1])
			num_to, _ := strconv.Atoi(cmd[3])
			value, _ := strconv.Atoi(cmd[5])
			clients[num_from-1].CallWithdraw("", cmd[2], value, true)
			clients[num_to-1].CallStore("", cmd[4], value, true)
		} else if cmd[0] == "getTxById" {
			if len(cmd) != 3 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			clients[num-1].GetTxById(cmd[2])
		} else if cmd[0] == "getTxCountByHeight" {
			if len(cmd) != 3 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			height, _ := strconv.Atoi(cmd[2])
			count := clients[num-1].GetBlockTxNumByHeight(int64(height))
			fmt.Println(count)
		} else if cmd[0] == "getTxCount" {
			if len(cmd) != 2 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			count := clients[num-1].GetTxCount()
			fmt.Println(count)
		} else if cmd[0] == "getBlockInfo" {
			if len(cmd) != 3 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			height, _ := strconv.Atoi(cmd[2])
			client.GetTxCount(clients, group_num, num, height)
		} else if cmd[0] == "getHeight" {
			if len(cmd) != 2 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			height := clients[num-1].GetCurrentBlockHeight()
			fmt.Println(height)
		} else if cmd[0] == "getCurrentTime" {
			if len(cmd) != 2 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			time := clients[num-1].GetCurrentBlockTime()
			fmt.Println(time)
		} else if cmd[0] == "getTxPool" {
			if len(cmd) != 2 {
				fmt.Println("参数数量错误!")
				continue
			}
			num, _ := strconv.Atoi(cmd[1])
			list, _ := clients[num-1].GetTxPool()
			fmt.Println(len(list))
		} else if cmd[0] == "latency" {
			if len(cmd) != 2 {
				fmt.Println("参数数量错误!")
				continue
			}
			cnt, _ = strconv.Atoi(cmd[1])
			client.Latency(clients, cnt, group_num)
		} else if cmd[0] == "addTx" {
			if len(cmd) == 2 {
				multi_cnt = -1
			} else if len(cmd) == 3 {
				multi_cnt, _ = strconv.Atoi(cmd[2])
			} else {
				fmt.Println("参数数量误!")
				continue
			}
			cnt, _ = strconv.Atoi(cmd[1])
			queue = client.AddTxs(cnt, multi_cnt, group_num, client_num, &tx_map)
		} else if cmd[0] == "balance" {
			if len(cmd) != 1 {
				fmt.Println("参数数量错误!")
				continue
			}
			queue, tx_map = client.BalanceTxs(queue, group_num, client_num, tx_map)
		} else if cmd[0] == "invokeTx" {
			if len(cmd) != 1 {
				fmt.Println("参数数量错误!")
				continue
			}
			client.InvokeAllTx(clients, group_num, client_num, cnt, queue, tx_map)
		} else {
			fmt.Println("没有此命令!")
			continue
		}
	}
}
