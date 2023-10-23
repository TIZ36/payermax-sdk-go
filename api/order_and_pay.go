package api

import "encoding/json"

const (
	OrderAndPayApiPath = "/orderAndPay"
)

type OrderAndPayApiReqData struct {
	// [must] 商户订单号，唯一标识商户的一笔交易，不能重复，只能包含字母、数字、下划线。商户订单号生成规则，不支持大小写敏感。举例：AAA和AAa被认为是相同的。
	OutTradeNo string `json:"outTradeNo"`

	// [must] 使用PayerMax_Hosted_Payment_Page进行付款方支付信息收集及处理
	Integrate string `json:"integrate"`

	// [must] 订单标题或产品信息，避免使用纯数字，如使用PayerMax_Hosted_Payment_Page，会展示在用户支付页面 注：巴西钱包Pix不能超过43位
	Subject string `json:"subject"`

	// [must] 标价金额，金额的单位为元。【风控行业限额】 除以下国家外按照各国币种支持的小数点位（最大4位小数）上送。 注意：巴林、科威特、伊拉克，约旦、突尼斯、利比亚、奥马尔地区，本币只支持两位小数； 印尼、中国台湾、巴基斯坦、哥伦比亚地区，本币不支持带小数金额。详见【交易支持国家/地区与币种】
	TotalAmount json.Number `json:"totalAmount"`

	// [must] 标价币种，大写字母，参见【交易支持国家/地区与币种】
	Currency string `json:"currency"`

	// [optional] 国家代码，大写字母，如果所传的国家代码与币种不匹配，则以货币代码对应的地区展示收银台。如指定了支付方式，则国家必须上送，参见【交易支持国家/地区与币种】如果下单不传国家，PayerMax判断用户地区方案优先级为：参数所传币种>参数所传国家 >根据用户使用习惯，之前使用过的国家 > 用户ip > 让用户自己选择地区
	Country string `json:"country"`

	// [must] 商户内部的用户Id，需要保证每个ID唯一性，支付方式保存后会根据userId进行支付方式推荐
	UserId string `json:"userId"`

	// [optional] 收银台页面语言。【支持的国家与币种】
	//优先级备注：用户上次使用的语言 > 用户浏览器语言 > 用户ip国家语言 > 商户下单传的语言 > 默认EN
	Language string `json:"language,omitempty"`

	// [optional] 商户自定义附加数据，可支持商户自定义并在响应中返回
	ReferralCode string `json:"referralCode,omitempty"`

	// optional 商户自定义附加数据，可支持商户自定义并在响应中返回
	Reference string `json:"reference,omitempty"`

	// [must] 商户指定的跳转URL，用户完成支付后会被跳转到该地址，以http/https开头或者商户应用的scheme地址（目前仅支持Andriod）
	FrontCallbackURL string `json:"frontCallbackURL"`

	// [must] 服务端回调通知URL，以http/https开头 可以通过MerchantDashboard平台配置商户通知地址，详情见【配置异步通知地址】，如果交易中上送，则以交易为准
	NotifyUrl string `json:"notifyUrl"`

	// [optional] 指定关单时间(单位：秒)。最小30分钟，最大1天。若传入则以该时间为关单时间。默认30分钟
	ExpireTime string `json:"expireTime,omitempty"`

	// [optional] 支付信息，非必填 1、支持仅指定支付方式，收银台会拉取该支付方式支持的所有目标机构 2、支持指定支付方式+目标机构
	PaymentDetail *PaymentDetail `json:"paymentDetail,omitempty"`

	// [optional] 设备信息
	EnvInfo *EnvInfo `json:"envInfo,omitempty"`

	// [optional] 二级商户信息
	SubMerchant *SubMerchant `json:"subMerchant,omitempty"`

	// [must] 商品信息
	GoodsDetails []*GoodsDetail `json:"goodsDetails"`

	// [optional] 邮寄信息 电商场景下需要上送
	ShippingInfo *ShippingInfo `json:"shippingInfo,omitempty"`

	BillingInfo *BillingInfo `json:"billingInfo,omitempty"`

	RiskParams map[string]any `json:"riskParams,omitempty"`

	SeparateIndicate string `json:"separateIndicate,omitempty"`

	SeparateAccountInfo []*SeparateAccountInfo `json:"separateAccountInfo,omitempty"`
}

type OrderAndPayApiRespData struct {
	// [must] 收银台支付跳转链接
	RedirectUrl string `json:"redirectUrl"`
	// [must] 商户订单号
	OutTradeNo string `json:"outTradeNo"`
	// [must] payermax订单号
	TradeToken string `json:"tradeToken"`
	// [must] 交易状态
	Status string `json:"status"`
}

type PaymentDetail struct {
	// [optional] 卡token支付时，此字段为必填
	PaymentTokenID string `json:"paymentTokenId,omitempty"`

	// [optional] 支付方式类型，为空时，拉取所有可用支付方式，参见【收银台支付-支付方式列表】
	PaymentMethodType string `json:"paymentMethodType,omitempty"`

	// [optional] 目标机构，可以为空，如果指定目标机构，则支付方式也必须指定。
	TargetOrg string `json:"targetOrg,omitempty"`

	// [optional] 支付账号，当指定支付方式和目标机构时，支持上送用户在该目标机构的支付账号信息。参见【收银台支付-支付方式列表】下的支持上送账号类型列
	PayAccountInfo *PaymentAccountInfo `json:"payAccountInfo,omitempty"`

	CardInfo *CardInfo `json:"cardInfo,omitempty"`

	// [optional] 用户基本信。备注：巴西国家下该字段必填
	BuyerInfo *BuyerInfo `json:"buyerInfo,omitempty"`
}

type PaymentAccountInfo struct {
	// [optional] 用户在该支付方式下的目标机构注册的支付账号,详见每个支付方式介绍
	AccountNo string `json:"accountNo"`

	// [optional] 账号类型：EMAIL，PHONE，ACCOUNT
	AccountNoType string `json:"accountNoType"`
}

type CardInfo struct {
	// [optional] 卡组织，当paymentMethod为CARD时可以指定具体的卡组织, 具体参见【收银台支付-支付方式列表https://docs.payermax.com/#/30?page_id=665&lang=zh-cn】
	CardOrg string `json:"cardOrg"`
}

type BuyerInfo struct {
	// [optional]
	FirstName string `json:"firstName,omitempty"`
	// [optional]
	MiddleName string `json:"middleName,omitempty"`
	// [optional]
	LastName string `json:"lastName,omitempty"`
	// [optional]
	PhoneNo string `json:"phoneNo,omitempty"`
	// [optional]
	IdType string `json:"idType,omitempty"`
	// [optional]
	IdNo string `json:"idNo,omitempty"`
	// [optional]
	TaxType string `json:"taxType,omitempty"`
	// [optional]
	TaxNo string `json:"taxNo,omitempty"`
	// [optional]
	Address string `json:"address,omitempty"`
	// [optional]
	City string `json:"city,omitempty"`
	// [optional]
	Region string `json:"region,omitempty"`
	// [optional]
	ZipCode string `json:"zipCode,omitempty"`
	// [optional]
	ClientIp string `json:"clientIp,omitempty"`
	// [optional]
	UserAgent string `json:"userAgent,omitempty"`
	// [optional]
	Email string `json:"email,omitempty"`
}

type EnvInfo struct {
	// [optional]
	DeviceId string `json:"deviceId,omitempty"`
	// [optional]
	DeviceLanguage string `json:"deviceLanguage,omitempty"`
	// [optional]
	ScreenHeight string `json:"screenHeight,omitempty"`
	// [optional]
	ScreenWidth string `json:"screenWidth,omitempty"`
}

type SubMerchant struct {
	// [optional] 二级商户号
	SubMerchantNo string `json:"subMerchantNo,omitempty"`
	// [optional] 二级商户名
	SubMerchantName string `json:"subMerchantName,omitempty"`
}

type GoodsDetail struct {
	// [must] 商品id
	GoodsId string `json:"goodsId"`

	// [must] 商品名称
	GoodsName string `json:"goodsName"`

	// [must] 	商品数量
	Quantity string `json:"quantity"`

	// [must] 商品单价
	Price string `json:"price"`

	// [optional] 商品报价币种，大写字母，参见【交易支持国家/地区与币种https://docs.payermax.com/#/30?page_id=677&lang=zh-cn】
	GoodsCurrency string `json:"goodsCurrency,omitempty"`

	//[optional] 商品展示地址(电商场景必传)
	ShowUrl string `json:"showUrl,omitempty"`

	//[must] 商品类别(电商场景必传)
	GoodsCategory string `json:"goodsCategory"`
}

type ShippingInfo struct {
	FirstName  string `json:"firstname"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName"`
	PhoneNo    string `json:"phoneNo"`
	Email      string `json:"email"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2,omitempty"`
	City       string `json:"city"`
	Region     string `json:"region,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country"`
	ZipCode    string `json:"zipCode"`
}
type BillingInfo struct {
	FirstName  string `json:"firstname,omitempty"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	PhoneNo    string `json:"phoneNo,omitempty"`
	Email      string `json:"email"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2,omitempty"`
	City       string `json:"city"`
	Region     string `json:"region,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country"`
	ZipCode    string `json:"zipCode"`
}

type SeparateAccountInfo struct {
	// [must] 分账参与方(分账场景必传)
	ParticipantId string `json:"participantId"`

	// [must] separateAccountDesc
	SeparateAccountDesc string `json:"separateAccountDesc"`
}
