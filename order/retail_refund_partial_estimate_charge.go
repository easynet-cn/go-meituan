package order

type RetailRefundPartialEstimateCharge struct {
	TotalOriginPrice           string                      `json:"total_origin_price"`            //描述：退款商品sku原价总价（含商品包装盒费），单位元
	TotalActivityPrice         string                      `json:"total_activity_price"`          //描述：退款商品sku实付价总价（含商品包装盒费），单位元
	TotalReducePrice           string                      `json:"total_reduce_price"`            //描述：退款商品sku总优惠金额，包括商家承担金额和美团承担金额，单位元
	TotalPoiCharge             string                      `json:"total_poi_charge"`              //描述：退款商品sku优惠中商家承担总金额，单位元
	OrderRetailActivityDetails []OrderRetailActivityDetail `json:"order_retail_activity_details"` //描述：商品参与活动详情
}
