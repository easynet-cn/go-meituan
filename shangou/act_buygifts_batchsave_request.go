package shangou

//【买赠】批量创建买赠活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/buygifts/batchsave
//文档地址：https://open-shangou.meituan.com/home/docDetail/247
type ActBuygiftsBatchsaveRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActData    string `url:"act_data"`     //是否必须：是；描述：买赠活动信息集合，json格式数组。(1)参数包括：a) app_food_code(必填项)，为APP方商品id，即商家中台系统里商品的编码（spu_code值），字段类型string。 注意：app_food_code即为药品的app_medicine_code。b) start_time(必填项)，为活动开始时间，字段类型int。传10位秒级的时间戳，系统会自动将传入的时间转换为所传日期的00:00:00；活动开始时间不得早于当前时间。 c) end_time(必填项)：活动结束时间，字段类型int。传10位秒级的时间戳。系统会自动将传入的时间转换为所传日期的23:59:59；活动结束时间不得早于当前时间和开始时间。 d) gifts_type(必填项) ，为赠品类型，字段类型int。取值范围：0非本商品非在售商品；1本商品；2非本商品在售商品。e) gifts_name，为赠品名称，字段类型string。场景说明：当gifts_type=0时，gifts_name必填，最多填写12个字长度；当gifts_type=2时，不能上传gifts_name字段，但须上传gifts_app_food_code字段且不能为本商品，需传赠品(非本商品在售商品)的app_food_code；当gifts_type=1时，gifts_name和gifts_app_food_code都无需填写。 注意：gifts_app_food_code即为药品的app_medicine_code。f) buy_num(必填项)，达到送赠品条件的商品购买数量，字段类型int。须上传大于0的正整数。g) gifts_num(必填项)，达到购买数量时赠送的赠品数量，字段类型int。须上传大于0的为正整数。例如buy_num=2，gifts_num=3，则代表买2件商品，赠送3件赠品。h) gifts_charge(必填项)：单个赠品的成本，单位是元，字段类型float。此字段信息须大于0，最多支持两位小数。注意，赠品成本为商家成本，不涉及美团成本。i) gifts_day_limit(必填项)，为当日活动赠品库存，字段类型int。须上传大于0的正整数或1，如传1则表示不限制活动赠品库存。(2)活动规则：a)同一商品同一生效时段仅可参加一个买赠活动；多规格商品不支持参加活动。b)买赠与折扣商品、第二份半价、指定商品满减活动均互斥。(3)同一次调用可上传的活动商品数据最多不能超过200组。
}
