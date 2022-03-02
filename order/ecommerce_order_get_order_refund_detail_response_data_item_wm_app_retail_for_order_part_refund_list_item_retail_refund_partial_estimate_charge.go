package order

type EcommerceOrderGetOrderRefundDetailResponseDataItemWmAppRetailForOrderPartRefundListItemRetailRefundPartialEstimateCharge struct {
	Total_origin_price            string                      `json:"total_origin_price"`            //描述：退款商品sku原价总价（含商品包装盒费），单位元
	total_activity_price          string                      `json:"total_activity_price"`          //描述：退款商品sku实付价总价（含商品包装盒费），单位元
	total_reduce_price            string                      `json:"total_reduce_price"`            //描述：退款商品sku总优惠金额，包括商家承担金额和美团承担金额，单位元
	total_poi_charge              string                      `json:"total_poi_charge"`              //描述：退款商品sku优惠中商家承担总金额，单位元
	order_retail_activity_details []OrderRetailActivityDetail `json:"order_retail_activity_details"` //描述：商品参与活动详情
}
