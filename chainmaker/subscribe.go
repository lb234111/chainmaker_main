// coding:utf-8
// 订阅事件
package chainmaker

import (
	"chainmaker/utils"
	"context"
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 监听事件
// @param cc 发起监听的客户端
// @param contractName 合约名
// @param eventName 事件名
// @param resultChan 保存事件的通道
func (c *Client) SubscribeEvent(contractName, eventName string, resultChan chan<- any) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ec, err := c.ChainmakerSDK.SubscribeContractEvent(ctx, -1, -1, contractName, eventName)
	if err != nil {
		utils.InfoError("subscribeHtlcProposal.cc.SubscribeContractEvent", err)
		return
	}
	for {
		select {
		case event, ok := <-ec:
			if !ok || event == nil {
				utils.InfoError("chainmaker subscribeEvent.event", err)
				close(resultChan)
				return
			}
			contractEvent, ok := event.(*common.ContractEventInfo)
			if !ok {
				utils.Info("event type convert error")
				close(resultChan)
				return
			}
			resultChan <- contractEvent
		case <-ctx.Done():
			fmt.Println("上下文关闭,结束监听")
			close(resultChan)
			return
		}
	}
}
