package shangou

//【折扣】批量修改药品折扣商品和爆品商品当日活动库存
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/retail/discount/batchstock
//文档地址：https://open-shangou.meituan.com/home/docDetail/308
type ActRetailDiscountBatchstockRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActType    int    `url:"act_type"`     //是否必须：否；描述：活动类型可取值： 1001折扣活动 56爆品活动 如不传本参数，则默认为修改折扣活动。
	ActData    string `url:"act_data"`     //是否必须：否；描述：折扣商品的当日活动库存信息，json格式数组。同一次调用中，此字段可上传活动商品数据不能超过200组，支持部分成功。
	ItemId     int64  `url:"item_id"`      //是否必须：否；描述：活动id
	DayLimit   int    `url:"day_limit"`    //是否必须：否；折扣商品当日活动库存，须传大于0的整数或1，如传1则表示不限制商品活动库存。
}
