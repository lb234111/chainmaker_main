// coding:utf-8
// 调用合约
package chainmaker

import (
	"chainmaker/config"
	"chainmaker/utils"
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 调用合约方法
// @param client: 客户端
// @param contractName: 合约名称
// @param method: 调用或者查询，这里一般是设置成invoke_contract
// @param txId: 可以置空
// @param kvs: 参数列表
// @param withSyncResult: 是否同步获取结果
func (c *Client) InvokeContract(
	contractName,
	txId string,
	kvs []*common.KeyValuePair,
	withSyncResult bool,
) (*common.TxResponse, error) {
	// inTime := time.Now()
	resp, err := c.ChainmakerSDK.InvokeContract(contractName, config.INVOKE_METHOD, txId, kvs, -1, withSyncResult)
	// txTime := time.Since(inTime)
	// fmt.Println("交易花费的时间:", txTime)

	if err != nil {
		info := utils.MyLogRespErr("client.InvokeContract", resp, err)
		return resp, fmt.Errorf(info)
	}
  //fmt.Printf("%+v\n",resp)
  
	if !withSyncResult {
		// 赋值
		resp.Code = common.TxStatusCode_SUCCESS
		resp.Message = "200"
		resp.ContractResult = &common.ContractResult{}
		resp.TxTimestamp = 0
		resp.TxBlockHeight = 0
		//resp.TxId = "id"
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		info := utils.MyLogRespErr("client.InvokeContract", resp, err)
		return resp, fmt.Errorf(info)
	}

	return resp, nil
}
