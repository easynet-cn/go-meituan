package act

//将门店商品从所参加的活动中删除
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/deleteActsByProducts
//文档地址：
type ActDeleteActsByProductsRequest struct {
	AppPoiCode   string `url:"app_poi_code"`   //是否必须：是；描述：1）字段描述：APP方门店id，即商家中台系统里门店的编码。
	AppFoodCodes string `url:"app_food_codes"` //是否必须：是；描述：APP方商品id，即商家中台系统里商品的编码,支持批量上传商品编码，单次调用最多支持50个商品编码。 说明：(1) 不同门店之间商品id可以重复，同一门店内商品id不允许重复；(2)字段信息限定长度不超过128个字符。 不存在的code，会默认忽略，传入重复code会默认去重
	Type         int    `url:"type"`           //是否必须：是；描述：当前支持操作剔除商品的活动类型。 商品活动集合： 折扣(17) 爆品(56) 指定商品满减(27) 买赠(23) 加价购(40) X元M件(43) 第N件优惠(55) 券类活动集合： 商品券(117) 闪购折扣商品券(124) 闪购活动商品券(123) 闪购活动商品折扣券(127) 闪购商品配送券(111) 闪购商品活动配送券(114)
	Status       int    `url:"status"`         //是否必须：是；描述：当前支持操作剔除商品的活动状态。 1-活动已生效（目前仅支持从已生效活动中剔除某商品）。不传默认为1。
}
