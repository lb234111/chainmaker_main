// coding:utf-8
// 新建长安链客户端
package chainmaker

import (
	"chainmaker/utils"
	"fmt"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

// 封装的长安链客户端
type Client struct {
	// 长安链SDK客户端
	ChainmakerSDK *sdk.ChainClient
}

// 创建链客户端
// @param sdkConfPath:客户端sdk路径
func NewChainmakerClient(sdkConfPath string) (*Client, error) {
	result := Client{}
	cc, err := sdk.NewChainClient(
		sdk.WithConfPath(sdkConfPath),
		sdk.WithEnableTxResultDispatcher(true),
	)
	if err != nil {
		info := utils.InfoError("sdk.NewChainClient", err)
		return nil, fmt.Errorf(info)
	}
	if cc.GetAuthType() == sdk.PermissionedWithCert {
		if err := cc.EnableCertHash(); err != nil {
			info := utils.InfoError("cc.EnableCertHash", err)
			return nil, fmt.Errorf(info)
		}
	}
	result.ChainmakerSDK = cc
	return &result, nil
}
