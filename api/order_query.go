package api

import "encoding/json"

const (
	OrderQueryApiPath = "/orderQuery"
)

// P16424103521913
type OrderQueryReqData struct {
	// [must] 商户订单号
	OutTradeNo string `json:"outTradeNo"`
}

type OrderQueryRespData struct {
	// [optional] 附加数据，商户上送的值
	Reference string `json:"reference,omitempty"`
	// [optional] 商户订单号
	OutTradeNo string `json:"outTradeNo"`

	// [optional] PayerMax订单号
	TradeToken string `json:"tradeToken"`

	// [must] 标价金额，金额的单位为元
	TotalAmount json.Number `json:"totalAmount"`

	// [must] 标价币种，大写字母
	Currency string `json:"currency"`

	// [optional] 渠道订单号
	ChannelNo string `json:"channelNo,omitempty"`

	// [optional] 三方单号
	ThirdChannelNo string `json:"thirdChannelNo,omitempty"`

	// [optional] VA单号
	PaymentCode string `json:"paymentCode,omitempty"`
	// [must] 用户支付的国家代码，大写字母
	Country string `json:"country"`
	// [optional] 交易状态
	Status string `json:"status,omitempty"`
	// [must] 完成时间
	CompleteTime string `json:"completeTime"`
	// [optional] 交易状态结果描述,仅失败时有值
	ResultMsg string `json:"resultMsg,omitempty"`
	// [optional] 支付信息，只有交易成功会返回支付方式，当交易失败时返回[ ]
	PaymentDetails []*PaymentDetail `json:"paymentDetails,omitempty"`
	// [optional]费用信息，支付成功且费用时才返回
	Fees []*Fee `json:"fees,omitempty"`
}
