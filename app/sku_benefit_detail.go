package app

type SkuBenefitDetail struct {
	AppFoodCode          string   `json:"app_food_code"`        //是否必须：否；描述：APP方商品id，即商家中台系统里商品的编码(spu_code值)，此app_food_code等同于app_medicine_code。(1)不同门店之间商品id可以重复，同一门店内商品id不允许重复。(2)字段信息限定长度不超过128个字符。(3)如此字段信息推送的是商品名称或信息为空，则表示商家没有维护商品编码，请商家自行维护。
	AppMedicineCode      string   `json:"app_medicine_code"`    //是否必须：否；描述：APP方商品id，即商家中台系统里商品的编码(spu_code值)，(1)不同门店之间商品id可以重复，同一门店内商品id不允许重复。(2)字段信息限定长度不超过128个字符。(3)如此字段信息推送的是商品名称或信息为空，则表示商家没有维护商品编码，请商家自行维护。此字段内容与app_food_code字段一致，平台建议商家对接此字段。
	Name                 string   `json:"name"`                 //是否必须：否；描述：商品名称
	Sku_id               string   `json:"sku_id"`               //是否必须：否；描述：SKU码(商家的规格编码)，是商品sku唯一标识码，药品都属于单规格。此字段信息与app_food_code相同，等同于药品app_medicine_code。
	Upc                  string   `json:"upc"`                  //是否必须：否；描述：商品的UPC码信息，即商品包装上的UPC/EAN码编号，长度一般8位或者13位。是商家同步商品信息时维护的UPC码，同一门店内，商品UPC码不允许重复。
	Count                *int     `json:"count"`                //是否必须：否；描述：订单中此商品sku的购买数量
	TotalOriginPrice     *float64 `json:"totalOriginPrice"`     //是否必须：否；描述：商品sku所参加的全部活动优惠前的原价，单位是元。
	OriginPrice          *float64 `json:"originPrice"`          //是否必须：否；描述：单件商品sku参加活动优惠前的原单价，单位是元。
	TotalReducePrice     *float64 `json:"totalReducePrice"`     //是否必须：否；描述：商品sku参加活动所优惠的总金额，单位是元。
	TotalActivityPrice   *float64 `json:"totalActivityPrice"`   //是否必须：否；描述：商品sku参活动优惠后的总价，单位是元。
	ActivityPrice        *float64 `json:"activityPrice"`        //是否必须：否；描述：单件商品sku参加活动优惠后的单价，单位是元。
	TotalMtCharge        *float64 `json:"totalMtCharge"`        //是否必须：否；描述：商品sku参加活动优惠的美团承担成本总金额，单位是元。
	TotalPoiCharge       *float64 `json:"totalPoiCharge"`       //是否必须：否；描述：商品sku参加活动优惠的商家承担成本总金额，单位是元。
	TotalBoxPrice        *float64 `json:"totalBoxPrice"`        //是否必须：否；描述：商品sku使用包装盒的总价，单位是元。药品类无商品维度包装盒，药品商家无需参考此字段信息。
	BoxPrice             *float64 `json:"boxPrice"`             //是否必须：否；描述：商品sku使用的包装盒单价，即单个包装盒的价格，单位是元。药品类无商品维度包装盒，药品商家无需参考此字段信息。
	BoxNumber            *float64 `json:"boxNumber"`            //是否必须：否；描述：单件此商品sku需使用包装盒的数量。药品类无商品维度包装盒，药品商家无需参考此字段信息。
	WmAppOrderActDetails string   `json:"wmAppOrderActDetails"` //是否必须：否；描述：参与活动详情
}
