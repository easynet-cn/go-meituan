package poi

type BankInfo struct {
	BankId   int    `json:"bank_id"`   //描述：银行or支行id
	BankName string `json:"bank_name"` //描述：银行or支行名称
}
