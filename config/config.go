package config

type PayMaxConfig struct {
	// 商户号
	MchNo string

	// ISV 商户号
	SpMchNo string

	// 授权码
	AuthToken string

	// AppID
	AppID string

	// SecretType
	SecretType string

	// MerchantPrivateKey
	MerchantPrivateKey string

	// PayMaxPublicKey
	PayMaxPublicKey string

	// 接口版本
	ApiVersion string

	// 密钥版本
	KeyVersion string

	// Notification callback urls
	Notify Notify
}

type Notify struct {
	// 收单
	ChargeCallbackUrl string

	// 退款
	RefundCallbackUrl string

	// 争议
	DisputeCallbackUrl string

	// 授权
	AuthCallbackUrl string

	// 账户服务
	AccountCallbackUrl string

	// 出款
	WithdrawCallbackUrl string
}
