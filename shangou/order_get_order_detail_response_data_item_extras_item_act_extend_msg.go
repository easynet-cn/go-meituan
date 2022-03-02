package shangou

type OrderGetOrderDetailResposneDataItemExtrasItemActExtendMsg struct {
	AppFoodCode          string `json:"app_food_code"`           //描述：主品app_food_code，只在买赠活动中支持展示，如买赠活动的主品没有维护app_food_code，则为空
	AppMedicineCode      string `json:"app_medicine_code"`       //描述：主品app_medicine_code，只在买赠活动中支持展示，如买赠活动的主品没有维护则为空,此字段内容与app_food_code字段一致，平台建议商家对接此字段。
	SkuId                string `json:"sku_id"`                  //描述：主品sku_id，只在买赠活动中支持展示，如买赠活动的主品没有维护sku_id，则为空
	GiftsAppFoodCode     string `json:"gifts_app_food_code"`     //描述：赠品app_food_code，如赠品没有维护app_food_code或非店内商品，则为空
	GiftsAppMedicineCode string `json:"gifts_app_medicine_code"` //描述：赠品app_medicine_code，如赠品没有维护app_medicine_code或非店内商品，则为空，此字段内容与app_food_code字段一致，平台建议商家对接此字段
	GiftsSkuOd           string `json:"gifts_sku_id"`            //描述：赠品sku_id，如赠品没有维护sku_id或非店内商品，则为空
	GiftsName            string `json:"gifts_name"`              //描述：赠品名称
	GiftNum              int    `json:"gift_num"`                //描述：赠品数量
}
