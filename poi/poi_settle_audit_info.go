package poi

type PoiSettleAuditInfo struct {
	AppPoiCode           string `json:"app_poi_code"`         //描述：门店编码
	Status               int    `json:"status"`               //描述：任务状态：0-待完善；1-审核中；2-已驳回；4-待提审；5-审核通过；7-暂不合作。无论传入的audit_type为何值，该字段均有值
	BaseStatus           int    `json:"base_status"`          //描述：基本信息模块：0-待录入；1-已录入；2-审核中；3-已驳回；4-签约中；5-签约失败；6-已生效；7-缺失。仅当audit_type=1时有效。
	LegalPersonStatus    int    `json:"legal_person_status"`  //描述：法人信息模块（仅单店）：0-待录入；1-已录入；2-审核中；3-已驳回；4-签约中；5-签约失败；6-已生效；7-缺失。仅当audit_type=1时有效。
	ShippingStatus       int    `json:"shipping_status"`      //描述：配送信息模块： 0-待录入；1-已录入；2-审核中；3-已驳回；4-签约中；5-签约失败；6-已生效；7-缺失。仅当audit_type=1时有效。
	QualificatsionStatus int    `json:"qualification_status"` //描述：资质信息模块：0-待录入；1-已录入；2-审核中；3-已驳回；4-签约中；5-签约失败；6-已生效；7-缺失。仅当audit_type=1时有效。
	Settlementtatus      int    `json:"settlement_status"`    //描述：结算信息模块（仅连锁门店）：0-待录入；1-已录入；2-审核中；3-已驳回；4-签约中；5-签约失败；6-已生效；7-缺失。仅当audit_type=1时有效。
	LogoStatus           int    `json:"logo_status"`          //描述：门店logo审核结果：1-审核中；2-审核通过；3-已驳回。仅当audit_type=2时有效。
	Reason               string `json:"reason"`               //描述：审核描述
}
