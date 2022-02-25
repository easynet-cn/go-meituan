package shangou

//【商家券】批量新增商家券（店内发券）活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/instore/coupon/batchsave
//文档地址：https://open-shangou.meituan.com/home/docDetail/315
type ActInstoreCouponBatchsaveRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	LimitTime  string `url:"limit_time"`   //是否必须：是；描述：活动时间，json格式字符串。时间会精确到开始时间的当前分钟数的0秒，结束时间的当前分钟数的59秒。
	ActData    string `url:"act_data"`     //是否必须：是；描述：商家券活动信息，json格式数组。(1)同一次调用中，最多可传3组优惠券信息，门店内同一时间优惠券不能超过3张。(2)同一次调用中，不支持优惠券部分创建成功。
}
