package poi

//获取三方门店映射信息
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/ecommerce/poi/bound/list
//文档地址：https://open-shangou.meituan.com/home/docDetail/442
type EcommercePoiBoundListRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：三方门店ID
	PageNum    int    `url:"page_num"`     //是否必须：是；描述：页数，从1开始。
	PageSize   int    `url:"page_size"`    //是否必须：是；描述：每页条数，最大支持200。
}
