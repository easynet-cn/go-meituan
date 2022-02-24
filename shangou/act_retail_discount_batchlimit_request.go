package shangou

//【折扣】批量修改药品折扣商品和爆品商品活动每单限购数量
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/retail/discount/batchlimit
//文档地址：https://open-shangou.meituan.com/home/docDetail/309
type ActRetailDiscountBatchlimitRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActType    int    `url:"act_type"`     //是否必须：否；描述：活动类型可取值： 1001折扣活动 56爆品活动 如不传本参数，则默认为修改折扣活动。
	ActData    string `url:"act_data"`     //是否必须：否；描述：折扣商品的每单限购的信息，json格式数组。同一次调用中，此字段可上传活动商品数据不能超过200组，支持部分商品修改成功。
}
