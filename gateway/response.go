package gateway

type PayerMaxResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func (resp *PayerMaxResponse) IsSuccess() bool {
	return resp.Code == SuccessReturn
}
