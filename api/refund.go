package api

import "encoding/json"

const (
	RefundApiPath = "/refund"
)

type RefundReqData struct {
	// [must] 商户退款申请号，唯一标识一笔退款申请交易
	OutRefundNo string `json:"outRefundNo"`

	// [must] 退款金额，金额的单位为元。 除以下国家外按照各国币种支持的小数点位（最大四位）上送。 注意：巴林、科威特、伊拉克，约旦、突尼斯、利比亚、奥马尔地区，本币只支持两位小数； 印尼、中国台湾、巴基斯坦、哥伦比亚地区，本币不支持带小数金额。
	RefundAmount json.Number `json:"refundAmount"`

	// [must] 退款币种（原交易币种）
	RefundCurrency string `json:"refundCurrency"`

	//[must] 商户订单号
	OutTradeNo string `json:"outTradeNo"`

	// [optional] 退款说明
	Comments string `json:"comments,omitempty"`

	// [optional] 退款回调地址，可后台配置【配置地址】
	RefundNotifyUrl string `json:"refundNotifyUrl,omitempty"`
}

type RefundRespData struct {
	// [must] 商户退款申请号
	OutRefundNo string `json:"outRefundNo"`
	// [must] 交易订单号
	TradeOrderNo string `json:"tradeOrderNo"`
	// [must] payermax 退款单号
	RefundTradeNo string `json:"refundTradeNo"`
	// [must] 退款状态：REFUND_PENDING，REFUND_FAILED
	Status string `json:"status"`
}
