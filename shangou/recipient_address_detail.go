package shangou

type RecipientAddressDetail struct {
	Province      string `json:"province"`       //描述：一级行政区(省级行政区):省、自治区、直辖市、特别行政区
	City          string `json:"city"`           //描述：二级行政区(地级行政区):地级市、地区、自治州、盟
	Area          string `json:"area"`           //描述：三级行政区(县级行政区):县、县级市、市辖区、自治县、旗、自治旗、矿区、林区、特区
	Town          string `json:"town"`           //描述：四级行政区(乡级行政区):乡、民族乡、镇、街道
	DetailAddress string `json:"detail_address"` //描述：详细地址
}
