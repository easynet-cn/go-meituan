package shangou

//更改门店公告信息
type PoiUpdatepromoteinfoRequest struct {
	AppPoiCode    string `url:"app_poi_codes"`  //是否必须：是；描述：APP方门店id，传商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	PromotionInfo string `url:"promotion_info"` //是否必须：是；描述：门店公告信息
}
