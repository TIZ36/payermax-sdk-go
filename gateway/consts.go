package gateway

const (
	UatUrl  = "https://pay-gate-uat.payermax.com/aggregate-pay/api/gateway"
	ProdUrl = "https://pay-gate.payermax.com/aggregate-pay/api/gateway"

	NotifyTypePayment = "PAYMENT"
	NotifyTypeRefund  = "REFUND"

	version    = "1.4"
	keyVersion = "1"

	SuccessReturn = "APPLY_SUCCESS"

	PayStatusSucc  = "SUCCESS"
	PayStatusFail  = "FAIL"
	PayStatusClose = "CLOSED"

	RefundStatusSucc = "REFUND_SUCCESS"
	RefundStatusFail = "REFUND_FAILED"
)
