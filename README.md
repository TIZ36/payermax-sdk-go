# payermax-sdk-go
payermax golang sdk

1. new payermax client
```go
// new a payermax client
defaultPayermaxClient := payermax.
    NewDefaultPayerMaxClient().
    SetGateway(gateway.UatUrl).
    SetConfig(&config.PayMaxConfig{
        MchNo:              "",
        AppID:              "",
		SecretType:         "",
        MerchantPrivateKey: "",
        PayMaxPublicKey:    "",
        ApiVersion:         "",
        KeyVersion:         "",
        Notify: config.Notify{
            ChargeCallbackUrl:   "",
            RefundCallbackUrl:   "",
            DisputeCallbackUrl:  "",
            AuthCallbackUrl:     "",
            AccountCallbackUrl:  "",
            WithdrawCallbackUrl: "",
        },
    })

// wrap with request-ctx to do request
reqCtx := payermax.RequestCtx[api.OrderQueryReqData, api.OrderQueryRespData]{
    PayerMaxClient: client,
}

if resp, err := reqCtx.Request(api.OrderQueryApiPath, api.OrderQueryReqData{
    OutTradeNo: "",
}); err != nil {
    panic(err)
} else {
    fmt.Println(resp)
}
```
