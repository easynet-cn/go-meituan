package shangou

type ActRetailDiscountBatchlimitRequestActDataItem struct {
	ItemId     int64 `json:"item_id"`     //描述：活动id
	OrderLimit int   `json:"order_limit"` //每单限购，表示同一订单中可享折扣价的数量。(1)此字段信息仅支持传入大于0的整数或1，传1则表示不限购。(2)如传大于0的整数，则最大值为10。如为爆品活动，则order_limit的最大值为999999999。（约等于无限制）
}
