package payermax

import (
	"encoding/json"
	"fmt"
)

type RequestCtx[T any, K any] struct {
	PayerMaxClient
}

func (ctx *RequestCtx[T, K]) Request(api string, params T) (K, error) {
	var resp K
	if respAny, err := ctx.PayerMaxClient.SendRequest(api, params); err != nil {
		return resp, err
	} else {
		fmt.Println("xxx ==>", respAny)

		respBin, _ := json.Marshal(respAny)
		json.Unmarshal(respBin, &resp)
		return resp, nil
	}
}
