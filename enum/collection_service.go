package enum

// @doc
// @name collection_service 收款服务
const (
	CollectionServiceTransactionStatusSuccess = "SUCCESS" // 交易成功
	CollectionServiceTransactionStatusPending = "PENDING" // 交易处理中
	CollectionServiceTransactionStatusFailed  = "FAILED"  // 交易失败
	CollectionServiceTransactionStatusClosed  = "CLOSED"  // 交易关闭

	TradeServiceRefundStatusSuccess = "REFUND_SUCCESS" // 退款成功
	TradeServiceRefundStatusPending = "REFUND_PENDING" // 退款处理中
	TradeServiceRefundStatusFailed  = "REFUND_FAILED"  // 退款失败
)

// error code list https://docs.payermax.com/#/30?page_id=662&lang=en
var (
	ApplySuccess                = PayermaxErr{Code: "APPLY_SUCCESS", Msg: "Successful application", Comment: "It only means that the request is successful. Please update the transaction result according to the status."}
	SystemError                 = PayermaxErr{Code: "SYSTEM_ERROR", Msg: "System error"}
	SignVerifyFailed            = PayermaxErr{Code: "SIGN_VERIFY_FAILED", Msg: "Signature error"}
	RequestTimeout              = PayermaxErr{Code: "REQUEST_TIMEOUT", Msg: "Request timed out"}
	PaymentRejected             = PayermaxErr{Code: "PAYMENT_REJECTED", Msg: "Request denied"}
	AuthenticateFailed          = PayermaxErr{Code: "AUTHENTICATE_FAILED", Msg: "Authorization failed"}
	IssuerPaymentRejected       = PayermaxErr{Code: "ISSUER_PAYMENT_REJECTED", Msg: "Request rejected by third party"}
	ParamsInvalid               = PayermaxErr{Code: "PARAMS_INVALID", Msg: "Invalid parameter"}
	MerchantInvalid             = PayermaxErr{Code: "MERCHANT_INVALID", Msg: "Invalid merchant"}
	MerchantAppInvalid          = PayermaxErr{Code: "MERCHANT_APP_INVALID", Msg: "Invalid Merchant APP"}
	ContractInvalid             = PayermaxErr{Code: "CONTRACT_INVALID", Msg: "Invalid contract"}
	OrderNotExist               = PayermaxErr{Code: "ORDER_NOT_EXIST", Msg: "Order number does not exist"}
	OrderClosed                 = PayermaxErr{Code: "ORDER_CLOSED", Msg: "Order closed"}
	OrderRepeat                 = PayermaxErr{Code: "ORDER_REPEAT", Msg: "Duplicate order"}
	PaymentMethodNotExist       = PayermaxErr{Code: "PAYMENT_METHOD_NOT_EXIST", Msg: "Payment method does not exist"}
	PaymentMethodSuspend        = PayermaxErr{Code: "PAYMENT_METHOD_SUSPEND", Msg: "The payment method is temporarily unavailable"}
	OnboardError                = PayermaxErr{Code: "ONBOARD_ERROR", Msg: "The merchant has not reported"}
	UnexpectedError             = PayermaxErr{Code: "UNEXPECTED_ERROR", Msg: "unknown mistake"}
	BalanceInsufficient         = PayermaxErr{Code: "BALANCE_INSUFFICIENT", Msg: "Insufficient balance"}
	CountryInvalid              = PayermaxErr{Code: "COUNTRY_INVALID", Msg: "Invalid country"}
	CurrencyInvalid             = PayermaxErr{Code: "CURRENCY_INVALID", Msg: "Invalid currency"}
	AmountLimit                 = PayermaxErr{Code: "AMOUNT_LIMIT", Msg: "Amount limit"}
	AmountLimitMinimum          = PayermaxErr{Code: "AMOUNT_LIMIT_MINIMUM", Msg: "Minimum amount limit"}
	AmountLimitMaximum          = PayermaxErr{Code: "AMOUNT_LIMIT_MAXIMUM", Msg: "maximum amount limit"}
	AmountInvalid               = PayermaxErr{Code: "AMOUNT_INVALID", Msg: "Invalid amount"}
	OTPVerifyLimit              = PayermaxErr{Code: "OTP_VERIFY_LIMIT", Msg: "OTP verification exceeds limit"}
	OTPVerifyFailed             = PayermaxErr{Code: "OTP_VERIFY_FAILED", Msg: "OTP verification failed"}
	OverVerifyLimit             = PayermaxErr{Code: "OVER_VERIFY_LIMIT", Msg: "Exceeded the number of verifications"}
	AuthFailed                  = PayermaxErr{Code: "AUTH_FAILED", Msg: "Authorization failed or does not exist"}
	AuthExpired                 = PayermaxErr{Code: "AUTH_EXPIRED", Msg: "Authorization expired"}
	BarcodeRefreshLimit         = PayermaxErr{Code: "BARCODE_REFRESH_LIMIT", Msg: "Barcode refresh limit"}
	BarcodeRefreshFailed        = PayermaxErr{Code: "BARCODE_REFRESH_FAILED", Msg: "Barcode refresh failed"}
	CardInvalid                 = PayermaxErr{Code: "CARD_INVALID", Msg: "Invalid card number"}
	CardExpireDateInvalid       = PayermaxErr{Code: "CARD_EXPIRE_DATE_INVALID", Msg: "Invalid card number validity period"}
	CardHolderNameInvalid       = PayermaxErr{Code: "CARD_HOLDER_NAME_INVALID", Msg: "Invalid cardholder name"}
	CVVInvalid                  = PayermaxErr{Code: "CVV_INVALID", Msg: "Invalid CVV"}
	UnsupportCard               = PayermaxErr{Code: "UNSUPPORT_CARD", Msg: "Card is not support"}
	AccountInvalid              = PayermaxErr{Code: "ACCOUNT_INVALID", Msg: "Invalid account"}
	PhoneNumInvalid             = PayermaxErr{Code: "PHONE_NUM_INVALID", Msg: "Invalid phone number"}
	UPIInvalid                  = PayermaxErr{Code: "UPI_INVALID", Msg: "Invalid UPI"}
	PINVerifyLimit              = PayermaxErr{Code: "PIN_VERIFY_LIMIT", Msg: "Pin verification exceeds limit"}
	PINInvalid                  = PayermaxErr{Code: "PIN_INVALID", Msg: "Pin is invalid"}
	BankCodeInvalid             = PayermaxErr{Code: "BANKCODE_INVALID", Msg: "Invalid bank card number"}
	IDNumInvalid                = PayermaxErr{Code: "ID_NUM_INVALID", Msg: "Invalid ID number"}
	EmailInvalid                = PayermaxErr{Code: "EMAIL_INVALID", Msg: "Invalid Email"}
	DocumentInvalid             = PayermaxErr{Code: "DOCUMENT_INVALID", Msg: "Invalid document"}
	TCKInvalid                  = PayermaxErr{Code: "TCK_INVALID", Msg: "Invalid TC Kimlik No."}
	DateInvalid                 = PayermaxErr{Code: "DATE_INVALID", Msg: "Invalid date"}
	PayeeNameInvalid            = PayermaxErr{Code: "PAYEE_NAME_INVALID", Msg: "Invalid payer name"}
	RemarkInvalid               = PayermaxErr{Code: "REMARK_INVALID", Msg: "Invalid remark"}
	CNICInvalid                 = PayermaxErr{Code: "CNIC_INVALID", Msg: "Invalid Cnic"}
	PaymentProcessing           = PayermaxErr{Code: "PAYMENT_PROCESSING", Msg: "Payment processing"}
	PaymentFailed               = PayermaxErr{Code: "PAYMENT_FAILED", Msg: "Payment failed"}
	RefundFailed                = PayermaxErr{Code: "REFUND_FAILED", Msg: "Refund failed"}
	RefundNoInvalid             = PayermaxErr{Code: "REFUND_NO_INVALID", Msg: "Invalid refund order number"}
	RefundInsufficientBalance   = PayermaxErr{Code: "REFUND_INSUFFICIENT_BALANCE", Msg: "Insufficient balance to refund"}
	AccountBlocked              = PayermaxErr{Code: "ACCOUNT_BLOCKED", Msg: "Account locked/frozen"}
	ParticipantInvalid          = PayermaxErr{Code: "PARTICIPANT_INVALID", Msg: "Invalid participant"}
	DeclinedByMerchantBlacklist = PayermaxErr{Code: "DECLINED_BY_MERCHANT_BLACKLIST", Msg: "Merchant blacklist blocking"}
)
