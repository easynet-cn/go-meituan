package shangou

//【第N件优惠】查询第N件优惠活动中的商品信息
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/act/nth/discount/products
//文档地址：https://open-shangou.meituan.com/home/docDetail/400
type ActNthDiscountProductsRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：PP方门店id，即商家中台系统里门店的编码。
	ActId      int64  `url:"act_id"`       //是否必须：是；描述：1）字段描述：第N件优惠活动id。 2）仅支持单个活动id。
	Offset     int    `url:"offset"`       //是否必须：是；描述：表示分页的页数，
	Limit      int    `url:"limit"`        //是否必须：是；描述：单页个数，最大不超过200
}
