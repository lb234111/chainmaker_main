// coding:utf-8
// 升级合约
package chainmaker

import (
	"fmt"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// 升级合约
// @param client: 客户端
// @param contractName: 合约名称
// @param upgradeVersion: 合约版本
// @param upgradeByteCodePath: 编译好的合约文件路径
// @param runtime: 运行方式WASM,DOCKER_GO,GASM等
// @param kvs: 合约初始化所需要的参数，以key-value形式传入
// @param withSyncResult: 是否需要同步获取交易结果
// @param usernames: 链节点数(需要节点背书同意)
func (c *Client) UpgradeContract(
	contractName string,
	upgradeVersion string,
	upgradeByteCodePath string,
	runtime common.RuntimeType,
	kvs []*common.KeyValuePair,
	withSyncResult bool,
	createContractTimeout int64,
	usernames ...string,
) (*common.TxResponse, error) {
	// 和部署合约的操作基本一样
	payload, err := c.ChainmakerSDK.CreateContractUpgradePayload(contractName, upgradeVersion, upgradeByteCodePath, runtime, kvs)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	// 获取背书节点
	endorsers, err := getEndorsers(c.ChainmakerSDK.GetHashType(),
		c.ChainmakerSDK.GetAuthType(), payload, usernames...)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	resp, err := c.ChainmakerSDK.SendContractManageRequest(payload, endorsers, createContractTimeout, withSyncResult)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	// ??检查
	err = checkProposalRequestResp(resp, false)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	fmt.Printf("upgrade contract success\n[code]:%v\n[message]:%v\n[contractResult]:%v\n[txid]:%v \n", resp.Code, resp.Message, resp.ContractResult, resp.TxId)
	return resp, nil
}
