package shangou

type SinglePoiCertInfo struct {
	TypeLevelOne             int    `json:"type_level_one"`               //是否必须：是；描述：一级证件类型，参考入驻流程附录：单店资质信息证件类型
	TypeLevelTwo             int    `json:"type_level_two"`               //是否必须：是；描述：二级证件类型，参考入驻流程附录：单店资质信息证件类型
	Name                     string `json:"name"`                         //是否必须：是；描述：证件名称
	Number                   string `json:"number"`                       //是否必须：是；描述：证件号码
	ValidStartDate           int64  `json:"valid_start_date"`             //是否必须：否；描述：证件有效起始时间，1代表长期有效，传入秒级时间戳
	ValidDate                int64  `json:"valid_date"`                   //是否必须：否；描述：证件有效期，1代表长期有效，传入秒级时间戳
	Address                  string `json:"address"`                      //是否必须：否；描述：证件注册地
	CertPic                  string `json:"cert_pic"`                     //是否必须：是；描述：资质图片 此字段支持直接传入图片URL地址也可通过【图片上传API】接口上传图片至美团系统，得到对应的图片ID上传。
	LegalPerson              string `json:"legal_person"`                 //是否必须：否；描述：法人姓名，资质类型为营业执照时必填
	SetUpDate                string `json:"set_up_date"`                  //是否必须：否；描述：证件注册机日期，秒级时间戳
	IsssuingAuthority        string `json:"isssuing_authority"`           //是否必须：否；描述：发证机关
	ApprovedDate             string `json:"approved_date"`                //是否必须：否；描述：核算日期，秒级时间戳
	BusinessScope            string `json:"business_scope"`               //是否必须：否；描述：经营范围
	Capital                  string `json:"capital"`                      //是否必须：否；描述：注册资本
	CreditCode               string `json:"credit_code"`                  //是否必须：否；描述：许可证社会信用代码
	SinglePoiLegalPersonInfo string `json:"single_poi_legal_person_info"` //是否必须：是；描述：法人信息，json对象
	SinglePoiCooperationInfo string `json:"single_poi_cooperation_info"`  //是否必须：是；描述：配送方式
}
