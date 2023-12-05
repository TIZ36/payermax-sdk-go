package payermax

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tiz36/payermax-sdk-go/api"
	"github.com/tiz36/payermax-sdk-go/config"
	"github.com/tiz36/payermax-sdk-go/enum"
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

	Err error
}

// NewDefaultPayerMaxClient 创建一个默认的 PayerMaxClient
func NewDefaultPayerMaxClient() *DefaultPayerMaxClient {
	return &DefaultPayerMaxClient{
		restyClient: resty.New(),
	}
}

// NewPayerMaxClientWithConfig 输入gateway和config 创建一个的 PayerMaxClient
func NewPayerMaxClientWithConfig(gateway string, config *config.PayMaxConfig) (PayerMaxClient, error) {
	if gateway == "" {
		return nil, errors.New("gateway url is empty")
	}
	if config.MchNo == "" {
		return nil, errors.New("merchant_no is empty")
	}
	if config.AppID == "" {
		return nil, errors.New("app_id is empty")
	}
	if config.PayMaxPublicKey == "" {
		return nil, errors.New("paymax_public_key is empty")
	}
	if config.MerchantPrivateKey == "" {
		return nil, errors.New("merchant_private_key is empty")
	}

	if config.SecretType == "" {
		config.SecretType = enum.DefaultSecretType
	}

	if config.ApiVersion == "" {
		config.ApiVersion = enum.DefaultApiVersion
	}
	if config.KeyVersion == "" {
		config.KeyVersion = enum.DefaultKeyVersion
	}

	return &DefaultPayerMaxClient{
		restyClient: resty.New(),
		config:      config,
		gateway:     gateway,
	}, nil
}

// SetGateway 设置环境，目标 PayerMax 的网关
func (client *DefaultPayerMaxClient) SetGateway(gateway string) PayerMaxClient {
	if client.Err != nil {
		return client
	}

	if gateway == "" {
		client.Err = errors.New("gateway url is empty")
		return client
	}

	client.gateway = gateway
	return client
}

// SetConfig 设置配置
func (client *DefaultPayerMaxClient) SetConfig(config *config.PayMaxConfig) PayerMaxClient {
	if client.Err != nil {
		return client
	}

	if config.MchNo == "" {
		client.Err = errors.New("merchant_no is empty")
		return client
	}
	if config.AppID == "" {
		client.Err = errors.New("app_id is empty")
		return client
	}
	if config.PayMaxPublicKey == "" {
		client.Err = errors.New("paymax_public_key is empty")
		return client
	}
	if config.MerchantPrivateKey == "" {
		client.Err = errors.New("merchant_private_key is empty")
		return client
	}

	if config.SecretType == "" {
		config.SecretType = enum.DefaultSecretType
	}
	if config.ApiVersion == "" {
		config.ApiVersion = enum.DefaultApiVersion
	}
	if config.KeyVersion == "" {
		config.KeyVersion = enum.DefaultKeyVersion
	}

	client.config = config
	return client
}

func (client *DefaultPayerMaxClient) GetErr() error {
	return client.Err
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

		if !payermaxResp.IsSuccess() {
			return nil, errors.New(
				fmt.Sprintf("请求失败, Code: %v, Msg:%v", payermaxResp.Code, payermaxResp.Msg))
		}

		return payermaxResp.Data, nil
	}
}

// VerifyNotification 验证回调通知
func (client *DefaultPayerMaxClient) VerifyNotification(signData []byte, signature string) (bool, error) {
	return gateway.Verify(signData, signature, client.config)
}

func (client *DefaultPayerMaxClient) ParseNotification(signData []byte) (string, any, error) {
	var notify gateway.PayerMaxNotify
	var dataBin []byte
	var err error

	if err := json.Unmarshal(signData, &notify); err != nil {
		return "", nil, err
	}

	if !notify.IsSuccess() {
		return notify.NotifyType, nil, errors.New("notify is not success")
	}

	if dataBin, err = json.Marshal(notify.Data); err != nil {
		return "", nil, errors.New("unknown notify data: " + err.Error())
	}

	switch notify.NotifyType {
	case enum.NotifyTypePayment:
		var paymentNotify api.PaymentNotifyData
		if err := json.Unmarshal(dataBin, &paymentNotify); err != nil {
			return enum.NotifyTypePayment, nil, errors.New("invalid payment notify data")
		} else {
			return enum.NotifyTypePayment, paymentNotify, nil
		}
	case enum.NotifyTypeRefund:
		var refundNotify api.RefundNotifyData
		if err := json.Unmarshal(dataBin, &refundNotify); err != nil {
			return enum.NotifyTypeRefund, nil, errors.New("invalid refund notify data")
		} else {
			return enum.NotifyTypeRefund, refundNotify, nil
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
