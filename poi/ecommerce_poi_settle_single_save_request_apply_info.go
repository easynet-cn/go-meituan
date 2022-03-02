package poi

//门店信息
type EcommercePoiSettleSingleSaveRequestApplyInfo struct {
	AppPoiCode               string `json:"app_poi_code"`                 //是否必须：是；描述：门店编码：创建时(type=0)传入的门店编码不能重复，例如已经创建一个编码为A的门店，下次新建时则不能使用A作为门店编码
	SinglePoiBasicInfo       string `json:"single_poi_basic_info"`        //是否必须：是；描述：门店基本信息
	SinglePoiCertInfos       string `json:"single_poi_cert_infos"`        //是否必须：是；描述：门店资质信息，json数组，根据type传入不同类型的资质信息。
	SinglePoiLegalPersonInfo string `json:"single_poi_legal_person_info"` //是否必须：是；描述：法人信息，json对象
	SinglePoiCooperationInfo string `json:"single_poi_cooperation_info"`  //是否必须：是；描述：配送方式
}
