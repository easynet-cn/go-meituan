package poi

//查询城市信息
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/ecommerce/poi/cityInfo
//文档地址：https://open-shangou.meituan.com/home/docDetail/679
type EcommercePoiCityInfoRequest struct {
	CityId int `url:"city_id"` //是否必须：否；描述：城市id,不传默认查省份信息，传入则查询下一级城市信息
}
