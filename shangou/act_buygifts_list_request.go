package shangou

//【买赠】批量查询门店买赠活动
//请求方式：GET
//文档地址：https://waimaiopen.meituan.com/api/v1/act/buygifts/list
//文档地址：https://open-shangou.meituan.com/home/docDetail/251
type ActBuygiftsListRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	Offset     int    `url:"offset"`       //是否必须：是；描述：分页查询偏移量，表示从第几条数据开始查，0表示第一条。
	Limit      int    `url:"limit"`        //是否必须：是；描述：分页查询每页展示条数，此字段信息最大不超过200。
}
