// coding:utf-8
// 部署合约
package chainmaker

import (
	"chainmaker/config"
	"chainmaker/utils"
	"errors"
	"fmt"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
)

// 部署合约
// @param client: 客户端
// @param contractName: 合约名称
// @param version: 合约版本
// @param byteCodePath: 编译好的合约文件路径
// @param runtime: 运行方式WASM,DOCKER_GO,GASM等
// @param kvs: 合约初始化所需要的参数，以key-value形式传入
// @param withSyncResult: 是否需要同步获取交易结果
// @param usernames: 链节点数(需要节点背书同意)
func (c *Client) DeployContract(
	contractName string,
	version string,
	byteCodePath string,
	runtime common.RuntimeType,
	kvs []*common.KeyValuePair,
	withSyncResult bool,
	createContractTimeout int64,
	usernames ...string,
) (*common.TxResponse, error) {
	// 测试看看这个合约在不在
	contractInfo, err := c.ChainmakerSDK.GetContractInfo(contractName)
	if err == nil {
		utils.InfoTips(contractName, "合约存在:")
		utils.Info(contractInfo)
		return nil, nil
	}

	payload, err := c.ChainmakerSDK.CreateContractCreatePayload(contractName, version, byteCodePath, runtime, kvs)
	if err != nil {
		info := utils.MyLogPayloadErr("client.CreateContractCreatePayload", payload, err)
		return nil, fmt.Errorf(info)
	}
	// 根据不同的模式，获取背书节点
	endorsers, err := getEndorsers(c.ChainmakerSDK.GetHashType(),
		c.ChainmakerSDK.GetAuthType(), payload, usernames...)
	if err != nil {
		return nil, fmt.Errorf("getEndorsers error: " + err.Error())
	}
	resp, err := c.ChainmakerSDK.SendContractManageRequest(payload, endorsers, createContractTimeout, withSyncResult)
	if err != nil {
		return resp, err
	}
	// 检查回执
	err = checkProposalRequestResp(resp, true)
	if err != nil {
		return resp, err
	}

	// log
	utils.Info("deploy contract success\n[code]:" + string(resp.Code) + "\n[message]:" + resp.Message + "\n[contractResult]:" + resp.ContractResult.String() + "\n[txid]:" + resp.TxId)
	return resp, nil
}

// 检查回执
func checkProposalRequestResp(resp *common.TxResponse, needContractResult bool) error {
	if resp.Code != common.TxStatusCode_SUCCESS {
		if resp.Message == "" {
			resp.Message = resp.Code.String()
		}
		return fmt.Errorf(resp.Message)
	}
	// todo:这是检查什么
	if needContractResult && resp.ContractResult == nil {
		return fmt.Errorf("contract result is nil")
	}
	// 0是成功代码
	if resp.ContractResult != nil && resp.ContractResult.Code != 0 {
		return fmt.Errorf(resp.ContractResult.Message)
	}
	return nil
}

// 自动根据客户端的连接模式获取背书节点
func getEndorsers(hashType crypto.HashType,
	authType sdk.AuthType,
	payload *common.Payload,
	usernames ...string,
) ([]*common.EndorsementEntry, error) {
	var endorsers []*common.EndorsementEntry

	for _, name := range usernames {
		var entry *common.EndorsementEntry
		var err error
		users := config.GetAdmins()
		pkUsers := config.GetPkUser()
		switch authType {
		case sdk.PermissionedWithCert:
			u, ok := users[name]
			if !ok {
				return nil, errors.New("user not found")
			}
			entry, err = sdkutils.MakeEndorserWithPath(u.SignKeyPath, u.SignCrtPath, payload)
			if err != nil {
				return nil, err
			}

		case sdk.Public:
			u, ok := pkUsers[name]
			if !ok {
				return nil, errors.New("user not found")
			}
			entry, err = sdkutils.MakePkEndorserWithPath(u.SignKeyPath, hashType, "", payload)
			if err != nil {
				return nil, err
			}

		default:
			return nil, errors.New("invalid authType")
		}
		endorsers = append(endorsers, entry)
	}

	return endorsers, nil
}
