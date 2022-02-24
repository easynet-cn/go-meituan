package shangou

type ActBuygiftsListResponseDataItem struct {
	ItemName         string `json:"item_name"`           //描述：买赠活动名称
	AppFoodCode      string `json:"app_food_code"`       //描述：活动商品，为APP方商品id，即商家中台系统里商品的编码（spu_code值）。 注意：app_food_code即为药品的app_medicine_code。
	GiftsDayLimit    int    `json:"gifts_day_limit"`     //描述：当日活动赠品库存，为大于0的正整数或1，如为1则表示不限制活动赠品库存。
	StartTime        int    `json:"start_time"`          //描述：活动开始时间，为10位秒级的时间戳；系统已自动将传入的时间转换为所传日期的00:00:00。
	EndTime          int    `json:"end_time"`            //描述：活动结束时间，为10位秒级的时间戳；系统已自动将传入的时间转换为所传日期的23:59:59。
	GiftsType        int    `json:"gifts_type"`          //描述：赠品类型，参考值：0非本商品非在售商品；1本商品；2非本商品在售商品。
	GiftsName        string `json:"gifts_name"`          //描述：赠品名称
	GiftsAppFoodCode string `json:"gifts_app_food_code"` //描述：赠品的app_food_code信息，非活动本商品的店内在售商品。 注意：app_food_code即为药品的app_medicine_code。
	BuyNum           int    `json:"buy_num"`             //描述：达到送赠品条件的商品购买数量
	GiftNum          int    `json:"gift_num"`            //描述：达到购买数量时赠送的赠品数量
	BuyGiftsModel    string `json:"buy_gifts_model"`     //描述：补贴模式，目前仅有“商家进货补贴模式(走结算流程)”模式。
	Charge           string `json:"charge"`              //描述：成本相关，单位是元。参数包括： giftsCharge，赠品成本； giftsPoiCharge，商家承担成本； giftsMtCharge，美团承担成本。
	ItemId           int    `json:"item_id"`             //描述：活动id
	Status           int    `json:"status"`              //描述：活动当前的状态，参考值：2待生效；1生效；0过期。
}
