# payermax-sdk-go
payermax golang sdk

Usage
-----
```go
package main

import (
	"fmt"
	iso "github.com/kzmnbrs/iso/go"
	"github.com/ladydascalie/currency"
	"github.com/tiz36/payermax-sdk-go/api"
	"github.com/tiz36/payermax-sdk-go/config"
	"github.com/tiz36/payermax-sdk-go/enum"
	"github.com/tiz36/payermax-sdk-go/payermax"
)

func main() {
	defaultPayermaxClient := payermax.
		NewDefaultPayerMaxClient().
		SetGateway(enum.UatUrl).
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

	if defaultPayermaxClient.GetErr() != nil {
		panic(defaultPayermaxClient.GetErr())
	}

	orderAndPay(defaultPayermaxClient)
}

func orderAndPay(client payermax.PayerMaxClient) {
	// 收银台-下单
	reqCtx := payermax.RequestCtx[api.OrderAndPayApiReqData, api.OrderAndPayApiRespData]{
		PayerMaxClient: client,
	}

	if resp, err := reqCtx.Request(api.OrderAndPayApiPath, api.OrderAndPayApiReqData{
		OutTradeNo:       "",
		Integrate:        "",
		Subject:          "iphone15promax",
		TotalAmount:      "0.09",
		Currency:         "USD",
		Country:          "",
		UserId:           "test-111",
		Language:         "en",
		ReferralCode:     "",
		FrontCallbackURL: "http://www.frontCallbackURL.com",
		NotifyUrl:        "http://www.notifyUrl.com",
		ExpireTime:       "1800",
		PaymentDetail:    &api.PaymentDetail{},
		GoodsDetails: []*api.GoodsDetail{
			{
				GoodsId:       "",
				GoodsName:     "",
				Quantity:      "",
				Price:         "",
				GoodsCurrency: "",
				ShowUrl:       "",
				GoodsCategory: "",
			},
		},
		ShippingInfo:        nil,
		BillingInfo:         nil,
		RiskParams:          nil,
		SeparateIndicate:    "",
		SeparateAccountInfo: nil,
	}); err != nil {
		panic(err)
	} else {
		fmt.Println(resp)
	}
}

```
