package app

type AuditDetail struct {
	SubModule    string `json:"sub_module"`    //描述：子模块 1-基本信息 2-资质信息 3-法人信息 4-配送信息 5-结算信息
	ModuleStatus string `json:"module_status"` //描述：子模块审核状态 0-待录入 1-已录入 2-审核中 3-已驳回 4-签约中 5-签约失败 6-已生效 7-缺失
	Reason       string `json:"reason"`        //描述：原因
}
