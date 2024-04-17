// coding:utf-8
// 获取区块
package chainmaker

import (
	"chainmaker/utils"
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 获取区块
func (c *Client) GetBlockByHeight(height int64) (*common.BlockInfo, error) {
	block, err := c.ChainmakerSDK.GetBlockByHeight(uint64(height), true)
	if err != nil {
		info := utils.InfoError("getBlockByHeight", err)
		return nil, fmt.Errorf(info)
	}
	return block, nil
}

//获取最新区块高度
func (c *Client) GetCurrentBlockHeight() int64 {
	height, err := c.ChainmakerSDK.GetCurrentBlockHeight()
	if err != nil {
		utils.InfoError("getBlockByHeight", err)
		return 0
	}
	return int64(height)
}

//获取指定高度区块的交易数量
func (c *Client) GetBlockTxNumByHeight(height int64) int {
	header, _ := c.ChainmakerSDK.GetBlockHeaderByHeight(uint64(height))
	cnt := header.TxCount
	return int(cnt)
}

//获取指定高度区块的所有交易id
func (c *Client) GetBlockTxIdByHeight(height int64) []string {
	info, _ := c.GetBlockByHeight(height)
	cnt := len(info.Block.Txs)
	id_list := []string{}
	for i := 0; i < cnt; i++ {
		id_list = append(id_list, info.Block.Txs[i].Payload.TxId)
	}
	return id_list
}

//获取最新区块的时间戳
func (c *Client) GetCurrentBlockTime() int64 {
	height, _ := c.ChainmakerSDK.GetCurrentBlockHeight()
	header, _ := c.ChainmakerSDK.GetBlockHeaderByHeight(height)
	time := header.BlockTimestamp
	//fmt.Printf("time:%+v\n",time)
	return time
}

//获取指定高度区块的时间戳
func (c *Client) GetBlockTimeByHeight(height int64) int64 {
	header, _ := c.ChainmakerSDK.GetBlockHeaderByHeight(uint64(height))
	time := header.BlockTimestamp
	//fmt.Printf("time:%+v\n",time)
	return time
}

//获取交易总数
func (c *Client) GetTxCount() int64 {
	var cnt uint32 = 0
	height, _ := c.ChainmakerSDK.GetCurrentBlockHeight()
	var i uint64
	for i = 1; i <= height; i++ {
		header, _ := c.ChainmakerSDK.GetBlockHeaderByHeight(i)
		cnt += header.TxCount
	}
	return int64(cnt)
}

//获取链的配置信息
func (c *Client) GetChainInfo() {
	info, _ := c.ChainmakerSDK.GetChainInfo()
	fmt.Printf("%+v\n\n", info)
}

//获取交易信息
func (c *Client) GetTxById(id string) *common.TransactionInfo {
	info, _ := c.ChainmakerSDK.GetTxByTxId(id)
	fmt.Println(info)
	return info
}
