package act

//查询商品参加的活动
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/act/all/get/byAppFoodCodes
//文档地址：https://open-shangou.meituan.com/home/docDetail/427
type ActAllGetByAppFoodCodesRequest struct {
	AppPoiCode   string `url:"app_poi_code"`   //是否必须：是；描述：APP方门店id
	Status       *int   `url:"status"`         //是否必须：否；描述：要查询的活动状态：0-活动未开始；1-活动已生效；5-活动冻结中；如果不传则查询全部状态
	AppFoodCodes string `url:"app_food_codes"` //是否必须：是；描述：需要查询的商品app_food_code，多个code之间用英文逗号分隔，最多支持50个商品
	QueryType    *int   `url:"query_type"`     //是否必须：否；描述：查询类型：0或不填-查询非买赠非券类活动；1-查询买赠活动；2-查询通过SKU创建的券；3-查询通过UPC创建的券
	PageNum      int    `url:"page_num"`       //是否必须：是；描述：页码：需为大于等于1的整数，从1开始
	PageSize     int    `url:"page_size"`      //是否必须：是；描述：每页大小：需为大于等于1的整数，最大值为50
}
