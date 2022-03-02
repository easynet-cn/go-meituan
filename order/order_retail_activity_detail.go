package order

type OrderRetailActivityDetail struct {
	OriginActivityId int64  `json:"origin_activity_id"` //描述：退款商品sku参与活动的活动id
	Type             int    `json:"type"`               //描述：参与活动的活动类型，参考值：1-首单立减；2-满减优惠；4-套餐惠赠优惠；9-美团红包；11-提前下单减优惠；17-折扣商品；18-美团专送再减；20-第二份半价优惠；22-门店新客立减；27-指定商品满减；40-加价购；43-X元M件；46-超值换购；66-会员折扣商品；101-商家代金券优惠；117-商品优惠券；118-商品折扣券；900-首单红包优惠。
	PoiCharge        string `json:"poi_charge"`         //描述：在此活动类型下退款商品sku优惠中分摊的商家承担金额（针对本次部分退款）
}
