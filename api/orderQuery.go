package api

const (
	OrderQueryApiPath = "/orderQuery"
)

// P16424103521913
type OrderQueryReqData struct {
	// [must] 商户订单号
	OutTradeNo string `json:"outTradeNo"`
}

type OrderQueryRespData = map[string]any
