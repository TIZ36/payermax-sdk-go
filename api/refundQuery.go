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

type RefundQueryRespData = map[string]any
