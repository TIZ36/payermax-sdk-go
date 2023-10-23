package payermax

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tiz36/payermax-sdk-go/api"
	"github.com/tiz36/payermax-sdk-go/config"
	"github.com/tiz36/payermax-sdk-go/gateway"
	"time"
)

const (
	// DateFormat 符合RFC3339规范的时间格式
	DateFormat      = "2006-01-02T15:04:05.000Z07:00"
	HeaderSign      = "sign"
	HeaderAuthToken = "merchant_auth_token"
)

type DefaultPayerMaxClient struct {
	// http client resty client
	restyClient *resty.Client

	// payermax config
	config *config.PayMaxConfig

	// gateway
	gateway string
}

// NewDefaultPayerMaxClient 创建一个默认的 PayerMaxClient
func NewDefaultPayerMaxClient() *DefaultPayerMaxClient {
	return &DefaultPayerMaxClient{
		restyClient: resty.New(),
	}
}

// NewPayerMaxClientWithConfig 输入gateway和config 创建一个的 PayerMaxClient
func NewPayerMaxClientWithConfig(gateway string, config *config.PayMaxConfig) PayerMaxClient {
	return &DefaultPayerMaxClient{
		restyClient: resty.New(),
		config:      config,
		gateway:     gateway,
	}
}

// SetGateway 设置环境，目标 PayerMax 的网关
func (client *DefaultPayerMaxClient) SetGateway(gateway string) PayerMaxClient {
	client.gateway = gateway
	return client
}

// SetConfig 设置配置
func (client *DefaultPayerMaxClient) SetConfig(config *config.PayMaxConfig) PayerMaxClient {
	client.config = config
	return client
}

// SendRequest 发送请求
func (client *DefaultPayerMaxClient) SendRequest(api string, params any) (any, error) {
	return client.SendRequestWithConfig(api, params, client.config)
}

// SendRequestWithConfig 发送请求
func (client *DefaultPayerMaxClient) SendRequestWithConfig(api string, params any, config *config.PayMaxConfig) (any, error) {
	var reqString string
	var err error

	// 构建请求字符串
	if reqString, err = buildReqString(config, params); err != nil {
		return "", err
	}

	sign, _ := gateway.Sign(config, reqString)

	fmt.Println("sign:", sign)
	fmt.Println("reqString:", reqString)

	// 构建http请求
	req := client.restyClient.
		R().
		SetHeader(HeaderSign, sign).
		SetHeader("Content-Type", "application/json").
		SetBody(reqString)

	if client.config.AuthToken != "" {
		req.SetHeader(HeaderAuthToken, client.config.AuthToken)
	}

	resp, err := req.Post(client.gateway + api)

	if !resp.IsSuccess() {
		// 请求不成功
		return "", err
	} else {
		var payermaxResp gateway.PayerMaxResponse
		err := json.Unmarshal(resp.Body(), &payermaxResp)

		if err != nil {
			return "", err
		}

		// TODO: check signature

		if !payermaxResp.IsSuccess() {
			return nil, errors.New(
				fmt.Sprintf("请求失败, Code: %v, Msg:%v", payermaxResp.Code, payermaxResp.Msg))
		}

		return payermaxResp.Data, nil
	}
}

// VerifyNotification 验证回调通知
func (client *DefaultPayerMaxClient) VerifyNotification(body string, signature string) (bool, error) {
	return gateway.Verify(body, signature, client.config)
}

func (client *DefaultPayerMaxClient) ParseNotification(body string) (string, any, error) {
	var notify gateway.PayerMaxNotify

	if err := json.Unmarshal([]byte(body), &notify); err != nil {
		return "", nil, err
	}

	if !notify.IsSuccess() {
		return notify.NotifyType, nil, errors.New("notify is not success")
	}

	switch notify.NotifyType {
	case gateway.NotifyTypePayment:

		if paymentNotify, ok := notify.Data.(api.PaymentNotifyData); ok {
			return gateway.NotifyTypePayment, paymentNotify, nil
		} else {
			return gateway.NotifyTypePayment, nil, errors.New("invalid payment notify data")
		}
	case gateway.NotifyTypeRefund:
		if refundNotify, ok := notify.Data.(api.RefundNotifyData); ok {
			return gateway.NotifyTypeRefund, refundNotify, nil
		} else {
			return gateway.NotifyTypeRefund, nil, errors.New("invalid refund notify data")
		}
	default:
		return "", nil, errors.New("unknown notify type")
	}
}

func (client *DefaultPayerMaxClient) GetConfig() *config.PayMaxConfig {
	return client.config
}

func buildReqString(config *config.PayMaxConfig, apiData any) (string, error) {
	var req = gateway.PayerMaxRequest{
		Version:     config.ApiVersion,
		KeyVersion:  config.KeyVersion,
		RequestTime: time.Now().UTC().Format(DateFormat),
		AppId:       config.AppID,
		MerchantNo:  config.MchNo,
		Data:        apiData,
	}

	// 设置isv商户号
	if len(config.SpMchNo) > 0 {
		req.ISVMerchantNo = config.SpMchNo
	}

	bin, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	return string(bin), nil
}
