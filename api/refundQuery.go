package api

const (
	RefundQueryApiPath = "/refundQuery"
)

type RefundQueryReqData struct {
	//[optional] 商户退款申请号，唯一标识一笔退款申请交易
	OutRefundNo string `json:"outRefundNo,omitempty"`

	//[optional] PayerMax退款流水号，与原商户退款单号二选一上送
	RefundTradeNo string `json:"refundTradeNo,omitempty"`
}

type RefundQueryRespData struct {
	// [optional] 商户退款申请号
	OutRefundNo string `json:"outRefundNo"`
	// [optional] PayerMax退款流水号
	RefundTradeNo string `json:"refundTradeNo"`
	// [optional] 退款金额，金额的单位为元
	RefundAmount string `json:"refundAmount"`
	// [optional] 退款币种（原交易币种）
	RefundCurrency string `json:"refundCurrency"`
	// [optional] 原商户订单号
	OutTradeNo string `json:"outTradeNo"`
	// [optional] 退款单状态，详见【退款状态】
	//
	//<= 32 字符
	//枚举值:
	//REFUND_SUCCESS
	//REFUND_PENDING
	//REFUND_FAILED
	Status string `json:"status"`
	// [optional] 交易状态结果描述，仅失败时有值
	ResultMsg string `json:"resultMsg,omitempty"`
}
