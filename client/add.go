// coding:utf-8
// 交易生成
package client

import (
	"chainmaker/utils"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
)

type Tx struct {
	Id     string
	Method string
	Num    int
	From   string
	To     string
	Value  int
}

// 生成交易的起始id
var new_txid int

// 客户端最大数量
const MAXQUEUE int = 400

// 生成多笔交易
// @param cnt： 生成交易数量
// @param r：生成跨片交易比例
// @param group_num: 分片数量
// @param client_num: 分片客户端数量
// @param tx_map: 第二阶段跨片跨片
func AddTxs(cnt int, r int, group_num int, client_num int, tx_map *map[string]Tx) [MAXQUEUE][]Tx {
	var queue [MAXQUEUE][]Tx
	txid, err := ioutil.ReadFile("Txid.txt")
	if err != nil {
		fmt.Println(err)
	}
	new_txid, err = strconv.Atoi(strings.Trim(string(txid), "\n"))
	if err != nil {
		fmt.Println(err)
	}
	var chain_num [16][3]int
	num1, num2 := 0, 0
	for i := 0; i < cnt; i++ {
		var twochain bool
		var tx1, tx2 Tx
		if r < 0 {
			twochain, tx1, tx2 = Txmakerrand(r, group_num)
		} else {
			twochain, tx1, tx2 = Txmaker(r, group_num)
		}
		client := rand.Intn(client_num) * group_num
		if twochain {
			num2++
			chain_num[tx1.Num][1]++
			chain_num[tx2.Num][2]++
			queue[client+tx1.Num] = append(queue[client+tx1.Num], tx1)
			(*tx_map)[tx1.Id] = tx2
		} else {
			num1++
			chain_num[tx1.Num][0]++
			queue[client+tx1.Num] = append(queue[client+tx1.Num], tx1)
		}
	}
	data := []byte(strconv.Itoa(new_txid + 1))
	ioutil.WriteFile("Txid.txt", data, 0664)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(num2)/float64(cnt)*100), 64)
	utils.Info("随机生成", cnt, "笔交易,其中跨链交易", num2, "笔,跨链交易占比为", value, "%")
	//utils.Info("随机生成", cnt, "笔交易")
	for i := 0; i < group_num; i++ {
		fmt.Println("第", i+1, "条链有", chain_num[i][0], "条片内交易,有", chain_num[i][1], "条第一阶段跨片交易,有", chain_num[i][2], "条第二阶段跨片交易")
	}
	return queue
}

//按跨片交易比例生成单笔交易
func Txmaker(r int, group_num int) (twochain bool, tx1 Tx, tx2 Tx) {
	n1 := rand.Intn(group_num)
	n2 := rand.Intn(group_num)
	random := rand.Intn(100)
	r1 := rand.Intn(2)
	r2 := rand.Intn(2)
	if random >= r {
		s := "transfer"
		c := 0
		id := strconv.Itoa(new_txid)
		new_txid++
		var tx1 Tx
		if r1%2 == 0 {
			tx1 = Tx{id, s, n1, "addr1", "addr2", c}
		} else {
			tx1 = Tx{id, s, n1, "addr2", "addr1", c}
		}
		return false, tx1, tx1
	} else {
		for ; n1 == n2; n2 = rand.Intn(group_num) {
		}
		s1 := "withdraw"
		s2 := "store"
		c := 0
		id1 := strconv.Itoa(new_txid)
		id2 := strconv.Itoa(new_txid + 1)
		new_txid = new_txid + 2
		var tx1, tx2 Tx
		tx1 = Tx{id1, s1, n1, "addr" + strconv.Itoa(r1+1), "", c}
		tx2 = Tx{id2, s2, n2, "addr" + strconv.Itoa(r2+1), "", c}
		return true, tx1, tx2
	}
}

//随机生成单笔交易
func Txmakerrand(r int, group_num int) (twochain bool, tx1 Tx, tx2 Tx) {
	n1 := rand.Intn(group_num)
	n2 := rand.Intn(group_num)
	r1 := rand.Intn(4)

	// for ; n2 >= group_num/2; n2 = rand.Intn(group_num) {
	// }
	if r == -2 && r1 < 2 {
		n2 = 0
	}
	if r == -3 && r1 < 3 {
		n2 = 0
	}
	if r == -4 {
		if r1 == 0 {
			n1 = 0
			n2 = 1
		} else if r1 == 1 {
			n1 = 1
			n2 = 2
		} else if r1 == 2 {
			n1 = 2
			n2 = 3
		} else {
			n1 = 3
			n2 = 0
		}
	}
	if n1 == n2 {
		s := "transfer"
		c := 0
		id := strconv.Itoa(new_txid)
		new_txid++
		var tx1 Tx
		if r1%2 == 0 {
			tx1 = Tx{id, s, n1, "addr1", "addr2", c}
		} else {
			tx1 = Tx{id, s, n1, "addr2", "addr1", c}
		}
		return false, tx1, tx1
	} else {
		s1 := "withdraw"
		s2 := "store"
		c := 0
		id1 := strconv.Itoa(new_txid)
		id2 := strconv.Itoa(new_txid + 1)
		new_txid = new_txid + 2
		var tx1, tx2 Tx
		tx1 = Tx{id1, s1, n1, "addr" + strconv.Itoa((r1+1)%2+1), "", c}
		tx2 = Tx{id2, s2, n2, "addr" + strconv.Itoa(r1%2+1), "", c}
		return true, tx1, tx2
	}
}
