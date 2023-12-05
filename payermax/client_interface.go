package payermax

import (
	"github.com/tiz36/payermax-sdk-go/config"
)

type PayerMaxClient interface {
	// SetGateway 设置环境，目标 PayerMax 的网关
	SetGateway(gateway string) PayerMaxClient

	// SetConfig 设置配置
	SetConfig(config *config.PayMaxConfig) PayerMaxClient

	// SendRequest 发送请求
	SendRequest(apiPath string, req any) (any, error)

	// SendRequestWithConfig 发送请求
	SendRequestWithConfig(api string, req any, config *config.PayMaxConfig) (any, error)

	// VerifyNotification 验证回调通知
	VerifyNotification(signData []byte, signature string) (bool, error)

	ParseNotification(signData []byte) (string, any, error)

	GetConfig() *config.PayMaxConfig

	GetErr() error
}
