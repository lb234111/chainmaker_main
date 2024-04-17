// coding:utf-8
// 服务端证书相关常量定义
package config

import (
	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// User用户结构定义了用户基础信息
type User struct {
	TlsKeyPath, TlsCrtPath   string
	SignKeyPath, SignCrtPath string
}

// 公钥模式下的用户结构
type PKUsers struct {
	SignKeyPath string
}

//证书sdk文件路径相关
const (
	CHAIN1_CLIENT1_SDK_CONFIG_PATH = "./config_files/sdkconfigs/chain1_sdkconfig1.yml" // chain1:user1
	CHAIN1_CLIENT2_SDK_CONFIG_PATH = "./config_files/sdkconfigs/chain1_sdkconfig2.yml" // chain1:user2
	CHAIN2_CLIENT1_SDK_CONFIG_PATH = "./config_files/sdkconfigs/chain2_sdkconfig1.yml" // chain2:user1
	CHAIN2_CLIENT2_SDK_CONFIG_PATH = "./config_files/sdkconfigs/chain2_sdkconfig2.yml" // chain2:user2
	CHAIN1_CLIENT3_SDK_CONFIG_PATH = "./config_files/sdkconfigs/chain1_sdkconfig3.yml" // chain1:user3
	CHAIN1_CLIENT4_SDK_CONFIG_PATH = "./config_files/sdkconfigs/chain1_sdkconfig4.yml" // chain1:user4
	CHAIN2_CLIENT3_SDK_CONFIG_PATH = "./config_files/sdkconfigs/chain2_sdkconfig3.yml" // chain2:user3
	CHAIN2_CLIENT4_SDK_CONFIG_PATH = "./config_files/sdkconfigs/chain2_sdkconfig4.yml" // chain2:user4
)

//合约调用相关
const (
	// invoke方法
	INVOKE_METHOD = "invoke_contract"
	//合约虚拟机参数
	DOCKER_GO = common.RuntimeType_DOCKER_GO //DOCKER_GO类型虚拟机
	WASM      = common.RuntimeType_WASMER    //WASMER类型虚拟机
	GASM      = common.RuntimeType_GASM      //GASM类型虚拟机
	//合约创建超时时间(单位: s)
	CREATE_CONTRACT_TIMEOUT = 10
)

//名称信息
const (
	//组织信息
	OrgId1 = "wx-org1.chainmaker.org" //组织1名称
	OrgId2 = "wx-org2.chainmaker.org" //组织2名称
	OrgId3 = "wx-org3.chainmaker.org" //组织3名称
	OrgId4 = "wx-org4.chainmaker.org" //组织4名称
	//组织管理员信息
	Chain1Org1Admin1 = "chain1org1admin1" //链1组织1的管理员名称
	Chain1Org2Admin1 = "chain1org2admin1" //链1组织2的管理员名称
	Chain1Org3Admin1 = "chain1org3admin1" //链1组织3的管理员名称
	Chain1Org4Admin1 = "chain1org4admin1" //链1组织4的管理员名称
	Chain2Org1Admin1 = "chain2org1admin1" //链2组织1的管理员名称
	Chain2Org2Admin1 = "chain2org2admin1" //链2组织2的管理员名称
	Chain2Org3Admin1 = "chain2org3admin1" //链2组织3的管理员名称
	Chain2Org4Admin1 = "chain2org4admin1" //链2组织4的管理员名称
	//客户端信息
	Chain1Org1Client1 = "chain1org1client1" //链1组织1的客户端名称
	Chain1Org2Client1 = "chain1org2client1" //链1组织2的客户端名称
	Chain1Org3Client1 = "chain1org3client1" //链1组织3的客户端名称
	Chain1Org4Client1 = "chain1org4client1" //链1组织4的客户端名称
	Chain2Org1Client1 = "chain2org1client1" //链2组织1的客户端名称
	Chain2Org2Client1 = "chain2org2client1" //链2组织2的客户端名称
	Chain2Org3Client1 = "chain2org3client1" //链2组织3的客户端名称
	Chain2Org4Client1 = "chain2org4client1" //链2组织4的客户端名称
	//测试用户信息
	User1A = "testuser1A"
	User1B = "testuser1B"
	User2A = "testuser2A"
	User2B = "testuser2B"
)

// 所有链1上的管理员用户信息
var Chain1Admins = []string{Chain1Org1Admin1, Chain1Org2Admin1, Chain1Org3Admin1, Chain1Org4Admin1}

// 所有链2上的管理员用户信息
var Chain2Admins = []string{Chain2Org1Admin1, Chain2Org2Admin1, Chain2Org3Admin1, Chain2Org4Admin1}
