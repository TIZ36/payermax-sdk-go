package api

import "encoding/json"

type RefundNotifyData struct {
	// must 商户退款申请号，唯一标识一笔退款申请交易
	OutRefundNo string `json:"outRefundNo"`

	// must PayerMax退款流水号
	RefundTradeNo string `json:"refundTradeNo"`

	// must 商户订单号
	OutTradeNo string `json:"outTradeNo"`

	// must 退款金额
	RefundAmount json.Number `json:"refundAmount"`

	// must 退款币种
	RefundCurrency string `json:"refundCurrency"`

	// must 退款状态
	Status string `json:"status"`
}
