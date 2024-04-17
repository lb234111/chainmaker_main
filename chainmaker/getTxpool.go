// coding:utf-8
// 查询交易池
package chainmaker

import (
	"chainmaker/utils"
	"fmt"
)

// 获取交易池中的交易
func (c *Client) GetTxPool() ([]string, error) {
	ls_result, err := c.ChainmakerSDK.GetTxIdsByTypeAndStage(3, 3)
	if err != nil {
		info := utils.InfoError("GetTxIdsByTypeAndStage", err)
		return []string{}, fmt.Errorf(info)
	}
	return ls_result, nil
}
