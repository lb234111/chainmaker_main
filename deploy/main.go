// coding:utf-8
// 部署合约
package main

import (
	"chainmaker/chainmaker"
	"chainmaker/config"
	"chainmaker/utils"
	"fmt"
	"os"
	"strconv"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// @param contractName 合约名
func deployContract(c *chainmaker.Client, contractName string) error {
	aProposalArgs := []*common.KeyValuePair{}
	_, err := c.DeployContract(
		contractName, "v0",
		fmt.Sprintf("./contracts/%s/%s.7z", contractName, contractName),
		config.DOCKER_GO,
		aProposalArgs,
		true,
		config.CREATE_CONTRACT_TIMEOUT,
		config.Chain1Admins...)
	if err != nil {
		info := utils.InfoError("deploy "+contractName+" contract", err)
		return fmt.Errorf(info)
	}
	utils.Info("deploy", contractName, "contract success")
	return nil
}

// 部署长安链合约
func deployChainmakerContract(configPath string) {
	utils.InfoTips("开始部署chainmaker合约")
	chainmaker_client, err := chainmaker.NewChainmakerClient(configPath)
	if err != nil {
		utils.InfoError("deploy chainmaker client", err)
		os.Exit(1)
	}
	// 部署donothing合约
	if err := deployContract(chainmaker_client, "donothing"); err != nil {
		utils.InfoError("deploy contract", err)
		os.Exit(1)
	}
	// 部署store合约
	if err := deployContract(chainmaker_client, "store"); err != nil {
		utils.InfoError("deploy contract", err)
		os.Exit(1)
	}
}

func main() {
	var group_num int = 1
	utils.Info("加载配置文件")
	for i := 0; i < group_num; i++ {
		// 证书路径
		path := "./config_files/sdkconfigs/chain" + strconv.Itoa(i+1+4) + "_sdkconfig1.yml"
		// 部署chainmaker合约
		deployChainmakerContract(path)
	}
}
