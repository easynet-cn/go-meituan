package poi

//法人信息
type SinglePoiLegalPersonInfo struct {
	Name             string `json:"name"`              //是否必须：是；描述：身份证姓名
	Number           string `json:"number"`            //是否必须：是；描述：身份证号
	ValidStartDate   int64  `json:"valid_start_date"`  //是否必须：是；描述：证件有效起始时间，秒级时间戳，1为长期
	ValidDate        int64  `json:"valid_date"`        //是否必须：是；描述：证件有效期，秒级时间戳，1为长期
	Address          string `json:"address"`           //是否必须：是；描述：住址
	Birthday         int64  `json:"birthday"`          //是否必须：是；描述：出生日期，秒级时间戳
	Nation           string `json:"nation"`            //是否必须：是；描述：民族
	IssuingAuthority string `json:"issuing_authority"` //是否必须：是；描述：签发机关
	IdCardFrontPic   string `json:"id_card_front_pic"` //是否必须：是；描述：手持身份证正面照片 需露出五官、证件信息清晰，完整（建议手持身份证往拍摄镜头前伸，人物只需露出脸部五官即可，相机聚焦到身份证上） 此字段支持直接传入图片URL地址也可通过【图片上传API】接口上传图片至美团系统，得到对应的图片ID上传。
	IdCardBackPic    string `json:"id_card_back_pic"`  //是否必须：是；描述：手持身份证背面照片 需露出五官、证件信息清晰，完整（建议手持身份证往拍摄镜头前伸，人物只需露出脸部五官即可，相机聚焦到身份证上） 此字段支持直接传入图片URL地址也可通过【图片上传API】接口上传图片至美团系统，得到对应的图片ID上传。
	PhoneNumber      string `json:"phone_number"`      //是否必须：是；描述：身份证对应的手机号（后期签约短信会发送到该手机）
	Sex              int    `json:"sex"`               //是否必须：是；描述：性别 1男性 2女性
}
