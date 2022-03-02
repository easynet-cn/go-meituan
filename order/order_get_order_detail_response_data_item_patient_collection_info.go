package order

type OrderGetOrderDetailResposneDataItemPatientCollectionInfo struct {
	HealthCodeAuditStatus int    `json:"health_code_audit_status"` //描述：健康码洁审核状态 -1：没有上传健康码 0：待审核 1：审核通过 2：审核不通过
	HealthCodePic         string `json:"health_code_pic"`          //描述：健康码图片
	IsDangerousArea       string `json:"is_dangerous_area"`        //描述：是否有危险地区居留史 1-是 0-否
	RouteCodeAuditStatus  int    `json:"route_code_audit_status"`  //描述：行程码审核状态 -1：没有上传健康码 0：待审核 1：审核通过 2：审核不通过
	RouteCodePic          string `json:"route_code_pic"`           //描述：行程码图片
	SymptomDesc           int    `json:"symptom_desc"`             //描述：症状描述
	Temperature           int    `json:"temperature"`              //描述：体温
}
