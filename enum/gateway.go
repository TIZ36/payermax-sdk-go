package enum

// @doc
// @name    gateway
// @brief   payermax 网关常量

const (
	UatUrl            = "https://pay-gate-uat.payermax.com/aggregate-pay/api/gateway" // payermax uat 网关地址
	ProdUrl           = "https://pay-gate.payermax.com/aggregate-pay/api/gateway"     // payermax prod 网关地址
	NotifyTypePayment = "PAYMENT"                                                     // 通知类型：支付
	NotifyTypeRefund  = "REFUND"                                                      // 通知类型：退款
	version           = "1.4"                                                         // 版本号
	keyVersion        = "1"                                                           // 密钥版本
)
