package shangou

type OrderGetOrderDetailResposneDataItemMedicareDetail struct {
	AppFoodCode           string `json:"app_food_code"`         //描述：APP方商品id，即商家中台系统里商品的编码(spu_code值)，此app_food_code等同于app_medicine_code。(1)不同门店之间商品id可以重复，同一门店内商品id不允许重复。(2)字段信息限定长度不超过128个字符。(3)如此字段信息推送的是商品名称或信息为空，则表示商家没有维护商品编码，请商家自行维护。
	AppMedicineCode       string `json:"app_medicine_code"`     //描述：APP方商品id，即商家中台系统里商品的编码(spu_code值)，内容等同于app_food_code编码，平台建议对接此字段。
	FoodName              string `json:"food_name"`             //描述：商品名称
	MedicareTotalPayPrice string `json:"medicareTotalPayPrice"` //描述：医保报销金额
	Quantity              int    `json:"quantity"`              //描述：此sku医保报销的数量
	SkuId                 string `json:"sku_id"`                //描述：SKU码(商家的规格编码)，是商品sku唯一标识码，药品都属于单规格。此字段信息与app_food_code相同，等同于药品app_medicine_code。
}
