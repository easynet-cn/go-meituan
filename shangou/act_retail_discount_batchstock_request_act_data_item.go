package shangou

type ActRetailDiscountBatchstockRequestActDataItem struct {
	ItemId   int64 `json:"item_id"`   //是否必须：是；描述：活动id
	DayLimit int   `json:"day_limit"` //是否必须：是；描述：折扣商品当日活动库存，须传大于0的整数或1，如传1则表示不限制商品活动库存。
}
