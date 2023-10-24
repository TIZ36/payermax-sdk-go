package enum

type PayermaxErr struct {
	Code    string `json:"code"`    // 结果码
	Msg     string `json:"msg"`     // 结果描述
	Comment string `json:"comment"` // 备注
}
