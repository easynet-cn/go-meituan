package act

//【商品券】查询门店下商品券活动
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/act/goods/coupon/list
//文档地址：https://open-shangou.meituan.com/home/docDetail/378
type ActGoodsCouponListRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActStatus  int    `url:"act_status"`   //是否必须：是；描述：活动当前状态，取值范围：1全部；0未生效；1已生效；2已结束。
	StartTime  int    `url:"start_time"`   //是否必须：是；描述：活动开始时间，传10位秒级的时间戳。注意：创建商品券活动时，活动开始和结束时间均分别自动转换为所传日期的00:00:00和23:59:59，查询店内活动时按转换后的时间为准。
	EndTime    int    `url:"end_time"`     //是否必须：是；描述：活动结束时间，传10位秒级的时间戳。注意：创建商品券活动时，活动开始和结束时间均分别自动转换为所传日期的00:00:00和23:59:59，查询店内活动时按转换后的时间为准。
	PageNum    int    `url:"page_num"`     //是否必须：是；描述：页码，表示从第几页开始查，1表示第一页。
	PageSize   int    `url:"page_size"`    //是否必须：是；描述：每页展示的数量
}
