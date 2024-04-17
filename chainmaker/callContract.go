// coding:utf-8
// 调用合约
package chainmaker

import (
	"chainmaker/utils"
	"strconv"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 调用donothing合约
func (client *Client) CallDonothing(Txid string, withSyncResult bool) {
	//utils.Info("调用donothing合约")
	resp, err := client.InvokeContract("donothing", Txid, nil, withSyncResult)
	if err != nil {
		info := utils.InfoError("invoke contract", err)
		panic(info)
	}
	utils.Info("调用成功, 结果:", resp)
}

// 调用store合约get方法
func (client *Client) CallGet(Txid string, addr string, withSyncResult bool) *common.TxResponse {
	//utils.Info("调用合约")
	// 参数
	kvs := []*common.KeyValuePair{
		{Key: "method", Value: []byte("get")},
		{Key: "addr", Value: []byte(addr)},
	}
	resp, err := client.InvokeContract("store", Txid, kvs, withSyncResult)
	if err != nil {
		info := utils.InfoError("invoke contract", err)
		panic(info)
	}
	utils.Info("调用成功, 结果:", resp)
	return resp
}

// 调用store合约store方法
func (client *Client) CallStore(Txid string, addr string, value int, withSyncResult bool) *common.TxResponse {
	//utils.Info("调用合约")
	// 参数
	kvs := []*common.KeyValuePair{
		{Key: "method", Value: []byte("store")},
		{Key: "addr", Value: []byte(addr)},
		{Key: "value", Value: []byte(strconv.Itoa(value))},
	}
	resp, err := client.InvokeContract("store", Txid, kvs, withSyncResult)
	if err != nil {
		info := utils.InfoError("invoke contract", err)
		panic(info)
	}
	utils.Info("调用成功, 结果:", resp)
	return resp
}

// 调用store合约withdraw方法
func (client *Client) CallWithdraw(Txid string, addr string, value int, withSyncResult bool) *common.TxResponse {
	//utils.Info("调用合约")
	// 参数
	kvs := []*common.KeyValuePair{
		{Key: "method", Value: []byte("withdraw")},
		{Key: "addr", Value: []byte(addr)},
		{Key: "value", Value: []byte(strconv.Itoa(value))},
	}
	resp, err := client.InvokeContract("store", Txid, kvs, withSyncResult)
	if err != nil {
		info := utils.InfoError("invoke contract", err)
		panic(info)
	}
	utils.Info("调用成功, 结果:", resp)
	return resp
}

// 调用store合约transfer方法
func (client *Client) CallTransfer(Txid string, addr1 string, addr2 string, value int, withSyncResult bool) *common.TxResponse {
	//utils.Info("调用合约")
	// 参数
	kvs := []*common.KeyValuePair{
		{Key: "method", Value: []byte("transfer")},
		{Key: "addr_from", Value: []byte(addr1)},
		{Key: "addr_to", Value: []byte(addr2)},
		{Key: "value", Value: []byte(strconv.Itoa(value))},
	}
	resp, err := client.InvokeContract("store", Txid, kvs, withSyncResult)
	if err != nil {
		info := utils.InfoError("invoke contract", err)
		panic(info)
	}
	utils.Info("调用成功, 结果:", resp)
	return resp
}
