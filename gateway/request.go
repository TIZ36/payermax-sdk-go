package gateway

// PayerMaxRequest PayerMax 请求， 参考：https://docs.payermax.com/#/30?page_id=784&lang=zh-cn
type PayerMaxRequest struct {
	// Version 接口版本, 当前值为 1.4
	Version string `json:"version"`

	// keyVersion 密钥版本, 当前值为 1
	KeyVersion string `json:"keyVersion"`

	// RequestTime 请求时间, 格式为 符合RFC3339规范的时间格式
	RequestTime string `json:"requestTime"`

	// AppId 应用ID
	AppId string `json:"appId"`

	// MerchantNo 商户号
	MerchantNo string `json:"merchantNo,omitempty"`

	// IsvMerchantNo ISV商户号
	ISVMerchantNo string `json:"spMerchantNo,omitempty"`

	// Data	 请求数据体
	Data any `json:"data"`
}
