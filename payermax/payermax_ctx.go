package payermax

type RequestCtx[T any, K any] struct {
	PayerMaxClient
}

func (ctx *RequestCtx[T, K]) Request(api string, params T) (K, error) {
	var resp K
	if respAny, err := ctx.PayerMaxClient.SendRequest(api, params); err != nil {
		return resp, err
	} else {
		return respAny.(K), nil
	}
}
