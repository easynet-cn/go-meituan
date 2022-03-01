package shangou

//查询门店内活动信息
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/act/all/get/byAppPoiCodeAndType
//文档地址：https://open-shangou.meituan.com/home/docDetail/445
type ActAllGetByAppPoiCodeAndTypeRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。
	Status     int    `url:"status"`       //是否必须：是；描述：活动状态：0-待生效，1-进行中，5-冻结。其中劵类活动没有5-冻结状态。
	Type       int    `url:"type"`         //是否必须：是；描述：活动类型：1-门店商品类+集合类活动，2-劵类（支持返回的券活动：商家券/活动商家券/商品券/活动商品券/配送券/活动配送券/商品配送券/活动商品配送券。）
}
