package shangou

//【第N件优惠】删除第N件优惠活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/nth/discount/delete
//文档地址：https://open-shangou.meituan.com/home/docDetail/397
type ActNthDiscountDeleteRequest struct {
	AppPoiCode   string `url:"app_poi_code"`   //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。
	ActId        string `url:"act_id"`         //是否必须：是；描述：1）字段描述：第N件优惠活动id。 2）仅支持单个活动id。
	AppFoodCodes string `url:"app_food_codes"` //是否必须：是；描述：1）字段描述：需从活动中删除的商品编码。多个用英文逗号隔开 2）如仅输入活动id，此字段为空，则删除活动及活动下全部商品。 3）如输入food code则删除活动内此商品。 4）最多支持200个app_food_code输入。
}
