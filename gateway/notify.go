package gateway

// PayerMaxNotify 异步通知
type PayerMaxNotify struct {
	// 返回码，’APPLY_SUCCESS’代表成功
	Code string `json:"code"`

	// 返回描述，’Success.’
	Msg string `json:"msg"`

	//密钥版本。当前值为：1
	KeyVersion string `json:"keyVersion"`

	//商户app id
	AppId string `json:"appId"`

	//商户Id
	MerchantNo string `json:"merchantNo"`

	//通知时间，符合rfc3339规范，格式：yyyy-MM-dd'T'HH:mm:ss.SSSXXX
	NotifyTime string `json:"notifyTime"`

	//通知类型 PAYMENT
	NotifyType string `json:"notifyType"`

	// 通知数据体
	Data any `json:"data"`
}

func (notify *PayerMaxNotify) IsSuccess() bool {
	return notify.Code == SuccessReturn
}

func NotifySuccess() map[string]any {
	return map[string]any{
		"code": "SUCCESS",
		"msg":  "Success",
	}
}
