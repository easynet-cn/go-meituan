package shangou

//【第N件优惠】查询第N件优惠活动信息
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/act/nth/discount/info
//文档地址：https://open-shangou.meituan.com/home/docDetail/398
type ActNthDiscountInfoRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：1）字段描述：APP方门店id，即商家中台系统里门店的编码。
}
