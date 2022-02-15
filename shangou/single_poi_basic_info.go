package shangou

//门店基本信息
type SinglePoiBasicInfo struct {
	FirstTag       int     `json:"first_tag"`       //是否必须：是；描述：一级品类，可以通过查询门店经营品类列表接口获取
	SecondTag      int     `json:"second_tag"`      //是否必须：是；描述：二号品类，可以通过查询门店经营品类列表接口获取
	ContractName   string  `json:"contract_name"`   //是否必须：是；描述：联系人
	ContractTel    string  `json:"contract_tel"`    //是否必须：是；描述：联系电话，用于接收商家账号信息。
	ShopName       string  `json:"shop_name"`       //是否必须：是；描述：门店名称
	Longitude      float64 `json:"longitude"`       //是否必须：是；描述：经度
	Latitude       float64 `json:"latitude"`        //是否必须：是；描述：维度
	Address        string  `json:"address"`         //是否必须：是；描述：门店详细位置
	ShopFrontPic   string  `json:"shop_front_pic"`  //是否必须：是；描述：门脸图 此字段支持直接传入图片URL地址也可通过【图片上传API】接口上传图片至美团系统，得到对应的图片ID上传。
	EnvironmentPic string  `json:"environment_pic"` //是否必须：是；描述：环境图 此字段支持直接传入图片URL地址也可通过【图片上传API】接口上传图片至美团系统，得到对应的图片ID上传。
	ShopLogoPic    string  `json:"shop_logo_pic"`   //是否必须：是；描述：门店logo图 此字段支持直接传入图片URL地址也可通过【图片上传API】接口上传图片至美团系统，得到对应的图片ID上传。
}
