package shangou

//【加价购】删除加价换购营销活动接口
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/markup/repurchase/delete
//文档地址：https://open-shangou.meituan.com/home/docDetail/364
type ActMarkupRepurchaseDeleteRequest struct {
	AppPoiCode   string `url:"app_poi_code"`   //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActId        string `url:"act_id"`         //是否必须：是；描述：活动id
	AppFoodCodes string `url:"app_food_codes"` //是否必须：是；描述：待删除的活动商品，传APP方商品id，即商家中台系统里商品的编码（spu_code值）。(1)同一次调用中，app_food_codes字段可上传需删除的活动商品数不能超过100个，以英文逗号隔开；支持活动商品部分被删除成功。(2)如活动中全部商品被删除，则此活动同时也会被自动删除。(3)如不传此字段，则会删除整个活动。 注意：app_food_code即为药品的app_medicine_code。
}
