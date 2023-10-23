package api

import "encoding/json"

type PaymentNotifyData struct {
	// must 商户订单号
	OutTradeNo string `json:"outTradeNo"`

	//must PayerMax交易流水号
	TradeToken string `json:"tradeToken"`

	//must
	TotalAmount json.Number `json:"totalAmount"`

	// must
	Currency string `json:"currency"`

	//optional
	Country string `json:"country,omitempty"`

	// must
	Status string `json:"status"`

	//must
	CompleteTime string `json:"completeTime"`

	//optional
	PaymentDetails []PaymentDetail `json:"paymentDetails,omitempty"`

	//optional 费用信息，支付成功且费用时才返回
	Fees Fee `json:"fees,omitempty"`

	//optional 附加信息
	Reference string `json:"reference,omitempty"`

	// optional 渠道订单号
	ChannelNo string `json:"channelNo,omitempty"`

	//optional 三方单号
	ThirdChannelNo string `json:"thirdChannelNo,omitempty"`

	// optional VA单号
	PaymentCode string `json:"paymentCode,omitempty"`
}

type Fee struct {
	//optional 商户费用
	MerFee MerFee `json:"merFee,omitempty"`
}

type MerFee struct {
	//must
	Url string `json:"url"`
	//must
	Amount string `json:"amount"`
	//must
	Currency string `json:"currency"`
}
