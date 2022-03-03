package app

type UserMemberInfo struct {
	is_poi_menber bool   `json:"is_poi_menber"` //是否必须：否；描述：是否为会员用户，参考值：ture-是，false-否。
	card_code     string `json:"card_code"`     //是否必须：否；描述：会员卡卡号，商家系统里此会员的卡号。
	level_code    string `json:"level_code"`    //是否必须：否；描述：会员等级，商家维护的会员等级code，对应信息请商家参考自己中台系统的相关信息。
	poi_member_id string `json:"poi_member_id"` //是否必须：否；描述：商家会员id，商家系统里的会员id。
}
