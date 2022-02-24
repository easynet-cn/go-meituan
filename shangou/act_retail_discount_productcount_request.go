package shangou

//【折扣】修改折扣/爆品活动门店级每单限购商品总数
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/retail/discount/productcount
//文档地址：https://open-shangou.meituan.com/home/docDetail/429
type ActRetailDiscountProductcountRequest struct {
	AppPoiCode   string `url:"app_poi_code"`   //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActType      int    `url:"act_type"`       //是否必须：是；描述：活动类型可取值： 1001折扣活动 56爆品活动 如不传本参数，则默认为修改折扣活动。
	ActProdCount int    `url:"act_prod_count"` //是否必须：是；描述：每单限购活动商品的总数量。支持传大于0的整数或1；传1表示不限
}
