package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

// 获取固定路径下管理员证书信息
func GetAdmins() map[string]*User {
	var CONFIG_PATH string
	_, filename, _, _ := runtime.Caller(0)
	CONFIG_PATH = filepath.Dir(filepath.Dir(filename)) + "/config_files/"
	var admins = map[string]*User{
		Chain1Org1Admin1: {
			CONFIG_PATH + "chain1/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.tls.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.tls.crt",
			CONFIG_PATH + "chain1/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.sign.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.sign.crt",
		},
		Chain1Org2Admin1: {
			CONFIG_PATH + "chain1/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.tls.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.tls.crt",
			CONFIG_PATH + "chain1/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.sign.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.sign.crt",
		},
		Chain1Org3Admin1: {
			CONFIG_PATH + "chain1/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.tls.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.tls.crt",
			CONFIG_PATH + "chain1/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.sign.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.sign.crt",
		},
		Chain1Org4Admin1: {
			CONFIG_PATH + "chain1/crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.tls.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.tls.crt",
			CONFIG_PATH + "chain1/crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.sign.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.sign.crt",
		},
		Chain2Org1Admin1: {
			CONFIG_PATH + "chain2/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.tls.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.tls.crt",
			CONFIG_PATH + "chain2/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.sign.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.sign.crt",
		},
		Chain2Org2Admin1: {
			CONFIG_PATH + "chain2/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.tls.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.tls.crt",
			CONFIG_PATH + "chain2/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.sign.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.sign.crt",
		},
		Chain2Org3Admin1: {
			CONFIG_PATH + "chain2/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.tls.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.tls.crt",
			CONFIG_PATH + "chain2/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.sign.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.sign.crt",
		},
		Chain2Org4Admin1: {
			CONFIG_PATH + "chain2/crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.tls.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.tls.crt",
			CONFIG_PATH + "chain2/crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.sign.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.sign.crt",
		},
	}
	return admins
}

// 获取公钥模式下的用户私钥地址
func GetPkUser() map[string]*PKUsers {
	var CONFIG_PATH string
	_, filename, _, _ := runtime.Caller(0)
	CONFIG_PATH = filepath.Dir(filepath.Dir(filename)) + "/config_files/"
	var admins = map[string]*PKUsers{
		Chain1Org1Admin1: {
			CONFIG_PATH + "chain1/crypto-config/node1/admin/admin1/admin1.key",
		},
		Chain1Org2Admin1: {
			CONFIG_PATH + "chain1/crypto-config/node2/admin/admin1/admin1.key",
		},
		Chain1Org3Admin1: {
			CONFIG_PATH + "chain1/crypto-config/node3/admin/admin1/admin1.key",
		},
		Chain1Org4Admin1: {
			CONFIG_PATH + "chain1/crypto-config/node4/admin/admin1/admin1.key",
		},
	}
	return admins
}

// 获取固定路径下普通客户端的证书信息
func GetClients() map[string]*User {
	var CONFIG_PATH string
	_, filename, _, _ := runtime.Caller(0)
	CONFIG_PATH = filepath.Dir(filepath.Dir(filename)) + "/config_files/"
	var clients = map[string]*User{
		Chain1Org1Client1: {
			CONFIG_PATH + "chain1/crypto-config/wx-org1.chainmaker.org/user/client1/client1.tls.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org1.chainmaker.org/user/client1/client1.tls.crt",
			CONFIG_PATH + "chain1/crypto-config/wx-org1.chainmaker.org/user/client1/client1.sign.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org1.chainmaker.org/user/client1/client1.sign.crt",
		},
		Chain1Org2Client1: {
			CONFIG_PATH + "chain1/crypto-config/wx-org2.chainmaker.org/user/client1/client1.tls.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org2.chainmaker.org/user/client1/client1.tls.crt",
			CONFIG_PATH + "chain1/crypto-config/wx-org2.chainmaker.org/user/client1/client1.sign.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org2.chainmaker.org/user/client1/client1.sign.crt",
		},
		Chain1Org3Client1: {
			CONFIG_PATH + "chain1/crypto-config/wx-org3.chainmaker.org/user/client1/client1.tls.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org3.chainmaker.org/user/client1/client1.tls.crt",
			CONFIG_PATH + "chain1/crypto-config/wx-org3.chainmaker.org/user/client1/client1.sign.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org3.chainmaker.org/user/client1/client1.sign.crt",
		},
		Chain1Org4Client1: {
			CONFIG_PATH + "chain1/crypto-config/wx-org4.chainmaker.org/user/client1/client1.tls.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org4.chainmaker.org/user/client1/client1.tls.crt",
			CONFIG_PATH + "chain1/crypto-config/wx-org4.chainmaker.org/user/client1/client1.sign.key",
			CONFIG_PATH + "chain1/crypto-config/wx-org4.chainmaker.org/user/client1/client1.sign.crt",
		},
		Chain2Org1Client1: {
			CONFIG_PATH + "chain2/crypto-config/wx-org1.chainmaker.org/user/client1/client1.tls.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org1.chainmaker.org/user/client1/client1.tls.crt",
			CONFIG_PATH + "chain2/crypto-config/wx-org1.chainmaker.org/user/client1/client1.sign.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org1.chainmaker.org/user/client1/client1.sign.crt",
		},
		Chain2Org2Client1: {
			CONFIG_PATH + "chain2/crypto-config/wx-org2.chainmaker.org/user/client1/client1.tls.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org2.chainmaker.org/user/client1/client1.tls.crt",
			CONFIG_PATH + "chain2/crypto-config/wx-org2.chainmaker.org/user/client1/client1.sign.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org2.chainmaker.org/user/client1/client1.sign.crt",
		},
		Chain2Org3Client1: {
			CONFIG_PATH + "chain2/crypto-config/wx-org3.chainmaker.org/user/client1/client1.tls.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org3.chainmaker.org/user/client1/client1.tls.crt",
			CONFIG_PATH + "chain2/crypto-config/wx-org3.chainmaker.org/user/client1/client1.sign.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org3.chainmaker.org/user/client1/client1.sign.crt",
		},
		Chain2Org4Client1: {
			CONFIG_PATH + "chain2/crypto-config/wx-org4.chainmaker.org/user/client1/client1.tls.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org4.chainmaker.org/user/client1/client1.tls.crt",
			CONFIG_PATH + "chain2/crypto-config/wx-org4.chainmaker.org/user/client1/client1.sign.key",
			CONFIG_PATH + "chain2/crypto-config/wx-org4.chainmaker.org/user/client1/client1.sign.crt",
		},
	}
	return clients
}

// 创建链客户端
// @param sdkConfPath:客户端sdk路径
// @param SugaredLogger:客户端日志选项
func CreateClient(sdkConfPath string) (*sdk.ChainClient, error) {
	cc, err := sdk.NewChainClient(sdk.WithConfPath(sdkConfPath))
	if err != nil {
		return nil, fmt.Errorf("sdk.NewChainClient: " + err.Error())
	}
	if cc.GetAuthType() == sdk.PermissionedWithCert {
		if err := cc.EnableCertHash(); err != nil {
			return nil, fmt.Errorf("cc.EnableCertHash: " + err.Error())
		}
	}
	return cc, nil
}

//初始化测试用户信息
func GetTestUser() (map[string]*sdk.ChainClient, error) {
	testUser := make(map[string]*sdk.ChainClient)
	c1A, err := CreateClient(CHAIN1_CLIENT1_SDK_CONFIG_PATH)
	if err != nil {
		return nil, fmt.Errorf("create testUser1A error: " + err.Error())
	}
	c1B, err := CreateClient(CHAIN1_CLIENT2_SDK_CONFIG_PATH)
	if err != nil {
		return nil, fmt.Errorf("create testUser1B error: " + err.Error())
	}
	c2A, err := CreateClient(CHAIN2_CLIENT3_SDK_CONFIG_PATH)
	if err != nil {
		return nil, fmt.Errorf("create testUser2A error: " + err.Error())
	}
	c2B, err := CreateClient(CHAIN2_CLIENT4_SDK_CONFIG_PATH)
	if err != nil {
		return nil, fmt.Errorf("create testUser2B error: " + err.Error())
	}
	testUser[User1A] = c1A
	testUser[User1B] = c1B
	testUser[User2A] = c2A
	testUser[User2B] = c2B
	return testUser, nil
}
