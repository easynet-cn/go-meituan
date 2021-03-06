package order

type OrderGetOrderDetailResposneDataItemExtrasItem struct {
	ActDetailId     int                                                       `json:"act_detail_id"`      //描述：活动id，目前此字段信息不适用于闪购品类订单商品优惠活动查询，请参考下方sku_benefit_detail->wmAppOrderActDetails->act_id 字段。
	ReduceFee       float64                                                   `json:"reduce_fee"`         //描述：本活动优惠的金额，单位是元。表示在用户本需支付的订单金额基础上减的金额，包括美团承担活动费用和商户承担活动费用的总和。 特殊情况说明：满赠、买赠活动的赠品成本不属于给用户减的金额，即当type=5、23时，reduce_fee=0。
	MtCharge        float64                                                   `json:"mt_charge"`          //描述：活动优惠金额中美团承担成本，单位是元。该信息默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“20 活动费用分摊”。
	PoiCharge       float64                                                   `json:"poi_charge"`         //描述：活动优惠金额中商家承担成本，单位是元。该信息默认不推送，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“20 活动费用分摊”。 特殊情况说明：满赠、买赠活动的赠品成本不属于给用户减的金额，即当type=5、23时，reduce_fee=0，poi_charge有值，为赠品成本。
	Remark          string                                                    `json:"remark"`             //描述：优惠说明
	Type            int                                                       `json:"type"`               //描述：优惠活动类型，参考值： 1-首单减；2-满减；3-抵价券；4-套餐赠；5-满赠；6-超时赔付；7-特价菜；8-首单返优惠劵；9-使用优惠劵；10-运营发优惠劵；11-提前下单减；12-满返优惠劵；13-当面付返优惠券；14-随机返优惠券；15-兑换红包；16-满减配送费；17-折扣商品；18-美团专送再减；19-使用点评优惠券；20-第2份半价活动；21-会员免配送费；22-门店新用户立减；23-买赠活动；24-AB新客活动；25-减配送费；100-满减商家优惠券（下单返券）；101-使用商家优惠券；102-供应链自定义；103-进店领券（商家券）；26-满减AB；27-指定商品满减；50-微信钱包渠道首减；51-点评app渠道首减；52-美团app渠道首减；28-新客满减；29-新客满减AB；30-多阶梯满减配送费；31-满N件折；32-扫码购偶数件折扣；33-会员折扣/特价；34-买满件阶梯特价/折扣；35-组合特价/折扣；37-品牌折扣商品；40-加价购；41-新客折扣菜；42-红包兑换, 不在B端配置，通过数据写入服务来创建；117-商品券；88-虚拟币；45-外卖拼团；43-X元M件 （可多种商品总计件数）；118-商品折扣券；66-闪购会员折；48-拼团减配送费；53-新客专享减包装费；54-新客专享减配送费； 55-第N件优惠; 56-闪购爆品；57-新客专享减打包袋费（针对闪购新客爆品活动）；59-新客专享减配送费（针对闪购新客爆品活动）；201-会员首购优惠；202-会员限时特惠；304-减配送费劵；300-商家会员减配送费；305-联盟津贴；84-店外领券；123-好友助力；127-活动商品券（活动商品折扣券、兑换券）；124-商品兑换券、折扣券；111-商品运费券；105-闪购减配送券；106-闪购活动减配送券；60-闪购拼团；115-好友助力券。
	Type9CouponType int                                                       `json:"type_9_coupon_type"` //描述：优惠券细分类型，订单使用优惠券(type=9)才有该字段 2-优惠券用于商品，3-优惠券用于运费券
	ActExtendMsg    OrderGetOrderDetailResposneDataItemExtrasItemActExtendMsg `json:"act_extend_msg"`     //描述：活动扩展信息，当前仅支持满赠活动和买赠活动。该字段默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“37 订单优惠信息中添加活动扩展信息”。
}
