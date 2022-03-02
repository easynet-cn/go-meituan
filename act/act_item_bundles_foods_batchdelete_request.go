package act

//【X元M件】删除X元M件活动中的商品
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/item/bundles/foods/batchdelete
//文档地址：https://open-shangou.meituan.com/home/docDetail/336
type ActItemBundlesFoodsBatchdeleteRequest struct {
	AppPoiCode   string `url:"app_poi_code"`   //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActId        int64  `url:"act_id"`         //是否必须：是；描述：活动id
	AppFoodCodes string `url:"app_food_codes"` //是否必须：是；描述：活动商品，传APP方商品id，即商家中台系统里商品的编码（spu_code值）。 注意：app_food_code即为药品的app_medicine_code。同一次调用中，app_food_codes字段可上传的需删除的活动商品数不能超过100个，以英文逗号隔开；支持活动商品部分被删除成功。如活动中全部商品被删除，则此活动同时也会被自动删除。
}
