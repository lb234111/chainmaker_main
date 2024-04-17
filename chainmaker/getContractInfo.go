// coding:utf-8
// 获取合约信息
package chainmaker

import "chainmaker.org/chainmaker/pb-go/v2/common"

// 根据合约名获取合约
func (c *Client) GetContractByName(contractName string) (*common.Contract, error) {
	return c.ChainmakerSDK.GetContractInfo(contractName)
}

// 获取全部合约
func (c *Client) GetAllContract() ([]*common.Contract, error) {
	return c.ChainmakerSDK.GetContractList()
}
