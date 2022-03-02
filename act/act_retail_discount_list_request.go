package act

//【折扣】查询门店药品折扣商品和爆品商品活动
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/act/retail/discount/list
//文档地址：https://open-shangou.meituan.com/home/docDetail/306
type ActRetailDiscountListRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActType    int    `url:"act_type"`     //是否必须：否；描述：活动类型可取值： 1001折扣活动 56爆品活动 如不传本参数，则默认为查询折扣活动。
	Offset     int    `url:"offset"`       //是否必须：是；描述：分页查询偏移量，表示从第几条数据开始查，0表示第一条。
	Limit      int    `url:"limit"`        //是否必须：是；描述：分页查询每页展示条数，此字段信息最大不超过200。
}
