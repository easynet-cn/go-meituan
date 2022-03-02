package act

//【商品满减】批量创建指定商品满减活动或创建店内满减活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/full/discount/batchsave
//文档地址：https://open-shangou.meituan.com/home/docDetail/318
type ActFullDiscountBatchsaveRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActInfo    string `url:"act_info"`     //是否必须：是；描述：活动信息，json格式字符串。
	ActDetails string `url:"act_info"`     //是否必须：是；描述：活动详情信息，json格式数组。活动满减不能超过5层，参数包括：(1)origin_price(必填项)，表示满的金额 ；(2)act_price(必填项)，表示减的金额。满、减金额都必须是大于0的整数，单位是元。举例：origin_price=20，act_price=2，表示满20元减2元。
	AppFoods   string `url:"app_foods"`    //是否必须：是；描述：活动商品信息，为json格式数组，此字段仅用于创建指定商品满减活动时上传。(1)参数包括：a)app_food_code(必填项)：APP方商品id，即商家中台系统里商品的编码（spu_code值）。b)day_limit(必填项)：当日活动商品库存，仅支持上传大于0的整数或1，如传1时则表示不限制商品活动库存。(2)同一次调用中可上传的商品信息不能超过200组，支持部分创建成功；如因某参数格式错误或商品code不存在，则本次调用将全部失败。(3)多规格商品不支持参加指定商品满减活动。
}
