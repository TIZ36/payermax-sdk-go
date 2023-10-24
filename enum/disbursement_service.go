package enum

// @doc
// @name disbursement_service
// @brief disbursement service enum

const (
	DisbursementserviceTransactionStatusSuccess    = "SUCCESS"    // 交易状态：成功
	DisbursementserviceTransactionStatusPending    = "PENDING"    // 交易状态：处理中
	DisbursementserviceTransactionStatusFail       = "FAILED"     // 交易状态：失败
	DisbursementserviceTransactionStatusBounceback = "BOUNCEBACK" // 交易状态：退票
)

var (
	PARAMETER_ERROR              = PayermaxErr{Code: "PARAMETER_ERROR", Msg: "The 【field】is incorrect"}
	INSUFFICIENT_QUOTA           = PayermaxErr{Code: "INSUFFICIENT_QUOTA", Msg: "Insufficient quota"}
	ORDER_REPEAT                 = PayermaxErr{Code: "ORDER_REPEAT", Msg: "The order number repeat."}
	BALANCE_INSUFFICIENT         = PayermaxErr{Code: "BALANCE_INSUFFICIENT", Msg: "Insufficient balance."}
	AMOUNT_LIMIT_UNMATCHED       = PayermaxErr{Code: "AMOUNT_LIMIT_UNMATCHED", Msg: "The single transfer amount does not meet the limit requirements"}
	CONTRACT_NOT_SIGNED          = PayermaxErr{Code: "CONTRACT_NOT_SIGNED", Msg: "Merchant didn’t sign the contract or the payment method."}
	RISK_INTERCEPT               = PayermaxErr{Code: "RISK_INTERCEPT", Msg: ""}
	ACCOUNT_ABNORMAL             = PayermaxErr{Code: "ACCOUNT_ABNORMAL", Msg: "Account is abnormal."}
	ACCOUNT_NOT_EXIST            = PayermaxErr{Code: "ACCOUNT_NOT_EXIST", Msg: "Merchant account does not exist."}
	PAYMENT_FAILED               = PayermaxErr{Code: "PAYMENT_FAILED", Msg: "The amount is not supported."}
	PAYMENT_METHOD_NOT_AVAILABLE = PayermaxErr{Code: "PAYMENT_METHOD_NOT_AVAILABLE", Msg: "The payment method is not avaliable."}
	REQUEST_PARAM_INVALID        = PayermaxErr{Code: "REQUEST_PARAM_INVALID", Msg: "Request param invalid"}
	ACCOUNT_BLOCKED              = PayermaxErr{Code: "ACCOUNT_BLOCKED", Msg: "Account blocked or frozen"}
	BANK_NOT_AVAILABLE           = PayermaxErr{Code: "BANK_NOT_AVAILABLE", Msg: "Wrong bank details or bank not available"}
	DECLINED_BY_RISK             = PayermaxErr{Code: "DECLINED_BY_RISK", Msg: "Beneficiary bank rejected credit for a reason."}
	INVALID_ACCOUNT              = PayermaxErr{Code: "INVALID_ACCOUNT", Msg: "Invalid account or not active"}
	INVALID_BANK_BRANCH          = PayermaxErr{Code: "INVALID_BANK_BRANCH", Msg: "Invalid bank branch"}
	INVALID_BANKCODE             = PayermaxErr{Code: "INVALID_BANKCODE", Msg: "invalid IFSC Code"}
	INVALID_DOCUMENT             = PayermaxErr{Code: "INVALID_DOCUMENT", Msg: "Invalid payee document"}
	INVALID_PAYEE_NAME           = PayermaxErr{Code: "INVALID_PAYEE_NAME", Msg: "Beneficiary bank rejected credit for a reason."}
	INVALID_PHONE_NO             = PayermaxErr{Code: "INVALID_PHONE_NO", Msg: "Invalid mobile number"}
	KYC_FAILED                   = PayermaxErr{Code: "KYC_FAILED", Msg: "Beneficiary bank rejected credit for a reason."}
	OTHER_ERRORS                 = PayermaxErr{Code: "OTHER_ERRORS", Msg: "Something went wrong! Please inform us."}
	PAYEE_AMOUNT_EXCEED_LIMIT    = PayermaxErr{Code: "PAYEE_AMOUNT_EXCEED_LIMIT", Msg: "Exceed user amount limit"}
	PAYMENT_PENDING              = PayermaxErr{Code: "PAYMENT_PENDING", Msg: "Payment under process, please check later."}
	PROVIDER_FAILED_PROCESS      = PayermaxErr{Code: "PROVIDER_FAILED_PROCESS", Msg: "Provider failed to process, please retry later"}
	PROVIDER_REFUSED_PROCESS     = PayermaxErr{Code: "PROVIDER_REFUSED_PROCESS", Msg: "Provider refused to process"}
	SYSTEM_ERROR                 = PayermaxErr{Code: "SYSTEM_ERROR", Msg: "System is busy, Please try again later."}
	AMOUNT_UNMATCHED             = PayermaxErr{Code: "AMOUNT_UNMATCHED", Msg: "The single transfer amount does not meet the limit requirements"}
	INVALID_ADDRESS              = PayermaxErr{Code: "INVALID_ADDRESS", Msg: "Invalid Address"}
	INVALID_EMAIL                = PayermaxErr{Code: "INVALID_EMAIL", Msg: "Email address is incorrect."}
	INVALID_IFSC_CODE            = PayermaxErr{Code: "INVALID_IFSC_CODE", Msg: "Invalid bank Code"}
	INVALID_OPERATOR             = PayermaxErr{Code: "INVALID_OPERATOR", Msg: "Invalid operator"}
	REDEEMCODE_EXPIRED           = PayermaxErr{Code: "REDEEMCODE_EXPIRED", Msg: "RedeemCode has expired"}
	TRANSACTION_EXCEED_LIMIT     = PayermaxErr{Code: "TRANSACTION_EXCEED_LIMIT", Msg: "Transaction exceed limit, please retry later.Carrier Billing disbursement,please check the restrictions:Carrier billing disbursement definition"}
	AMOUNT_NOT_ENOUGH            = PayermaxErr{Code: "AMOUNT_NOT_ENOUGH", Msg: "Insufficient balance"}
	BUSINESS_ERROR               = PayermaxErr{Code: "BUSINESS_ERROR", Msg: "Business error"}
	UNDEFINED_ERROR              = PayermaxErr{Code: "UNDEFINED_ERROR", Msg: "Undefined error"}
	DECLINED_BY_COMPLIANCE       = PayermaxErr{Code: "DECLINED_BY_COMPLIANCE", Msg: "Beneficiary bank rejected credit for a reason. Eg:Declined by compliance"}
	DECLINED_PAYEE_BLACKLISTED   = PayermaxErr{Code: "DECLINED_PAYEE_BLACKLISTED", Msg: "Beneficiary bank rejected credit for a reason. Eg:Payee is in the blacklist"}
	INVALID_ACCOUNT_FORMAT       = PayermaxErr{Code: "INVALID_ACCOUNT_FORMAT", Msg: "Invalid account format"}
	INVALID_ACCOUNT_TYPE         = PayermaxErr{Code: "INVALID_ACCOUNT_TYPE", Msg: "Invalid account type"}
	MISMATCHED_BANK              = PayermaxErr{Code: "MISMATCHED_BANK", Msg: "Beneficiary bank rejected credit for a reason. Eg:Account and bank mismatch"}
	MISMATCHED_DOCUMENTS         = PayermaxErr{Code: "MISMATCHED_DOCUMENTS", Msg: "Beneficiary bank rejected credit for a reason. Eg:Mismatched documents"}
	CURRENCY_UNMATCHED           = PayermaxErr{Code: "CURRENCY_UNMATCHED", Msg: "Beneficiary bank rejected credit for a reason. Eg: account currency does not match"}
)
