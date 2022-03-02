package order

type WmAppRetailForOrderPartRefundListItem struct {
	RetailRefundPartialEstimateCharge RetailRefundPartialEstimateCharge `json:"retail_refund_partial_estimate_charge"` //描述：部分退款的商品中预计费信息字段集合。 注意：仅提供有商家承担成本的优惠信息。
	AppFoodCode                       string                            `json:"app_food_code"`                         //描述：APP方商品id，即商家中台系统里商品的编码（spu_code值），字段信息限定长度不超过128个字符。
	AppMedicineCode                   string                            `json:"app_medicine_code"`                     //描述：APP方商品id，即商家中台系统里商品的编码（spu_code值），字段信息限定长度不超过128个字符。此字段内容与app_food_code字段一致，平台建议商家对接此字段。
	FoodName                          string                            `json:"food_name"`                             //描述：商品名称
	SkuId                             string                            `json:"sku_id"`                                //描述：商品sku唯一标识码，字段信息限定长度不超过40个字符。
	Upc                               string                            `json:"upc"`                                   //描述：商品的UPC码信息，即商品包装上的UPC/EAN码编号，长度一般8位或者13位。是商家同步商品信息时维护的UPC码，同一门店内，商品UPC码不允许重复。
	Spec                              string                            `json:"spec"`                                  //描述：商品sku的规格名称
	Count                             int                               `json:"count"`                                 //描述：本次退款的商品数量：(1)如为按件部分退款，此字段信息为本次退款商品sku的数量。(2)如为按克重退差价，此字段信息为0。
	BoxNum                            float64                           `json:"box_num"`                               //描述：商品sku单件需使用打包盒的数量。（商品维度，在创建/更新商品时维护的信息）
	BoxPrice                          float64                           `json:"box_price"`                             //描述：商品sku的单个打包盒的价格，单位是元。
	FoodPrice                         float64                           `json:"food_price"`                            //描述：当前商品sku参加商品类活动优惠后的金额（单价），单位是元。
	OriginFoodPrice                   float64                           `json:"origin_food_price"`                     //描述：商品sku优惠前原价(单价)，单位是元。此字段信息为当前订单中单件商品sku的原价。
	RefundPrice                       float64                           `json:"refund_price"`                          //描述：退款价格（单价），单位是元。此字段信息为当前订单中单件此商品sku的退款价格，是单价。(1)如购买多件商品sku，仅1件享优惠价，计算时商品优惠金额会进行等比分摊。(2)如商品是按克重退差价，refund_price字段信息是计算优惠分摊后，单件商品sku重量差异部分的价格。
	RefundedWeight                    float64                           `json:"refunded_weight"`                       //描述：商品sku的已退重量，单位是克/g。此字段仅适用于退差价类型，为商品sku标价重量与实拣重量的差异重量。
}
