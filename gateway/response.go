package gateway

import "github.com/tiz36/payermax-sdk-go/enum"

type PayerMaxResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func (resp *PayerMaxResponse) IsSuccess() bool {
	return resp.Code == enum.ApplySuccess.Code
}
