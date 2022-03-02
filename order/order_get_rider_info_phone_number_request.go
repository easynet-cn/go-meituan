package order

//拉取骑手真实手机号
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/order/getRiderInfoPhoneNumber
//文档地址：https://open-shangou.meituan.com/home/docDetail/389
type OrderGetRiderInfoPhoneNumberRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：非必填；描述：门店id
	Offset     int    `url:"offset"`       //是否必须：是；描述：分页查询的偏移量
	Limit      int    `url:"limit"`        //是否必须：是；描述：每条页数，需小于等于1000
}
