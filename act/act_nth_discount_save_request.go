package act

//【第N件优惠】创建/更新第N件优惠活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/nth/discount/save
//文档地址：https://open-shangou.meituan.com/home/docDetail/395
type ActNthDiscountSaveRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：门店id
	ActData    string `url:"act_data"`     //是否必须：是；描述：活动数据，包括活动信息以及活动所包含的商品
}
