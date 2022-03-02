package medicine

//无货取消订单触发库存清零相关配置
type StockConfig struct {
	Reset                 bool   `json:"reset"`                     //描述：是否开启“无货取消订单触发库存清零”配置的总开关： true开；false关。
	Type                  []int  `json:"type"`                      //描述：触发库存清零的取消订单方式： 1门店因无货取消订单；2买家因无货取消订单。
	LimitOpenSyncStock    bool   `json:"limit_open_sync_stock"`     //描述：触发库存清零后是否禁止开放平台接口更新库存： true是，禁止更新库存；false否，允许更新库存。
	Schedule              string `json:"schedule"`                  //描述：商品触发库存清零后，且limit_open_sync_stock为true的情况下，允许开放平台接口更新库存的时间点，日期为次日的“时:分”。limit_open_sync_stock为false时，无需关注此字段。
	SyncNextDay           bool   `json:"sync_next_day"`             //描述：是否允许次日自动补充库存： true是，自动补充库存；false否，不会自动补充库存。 注：如设置为“是”，则当商品开启“无货取消订单触发库存清零”配置，且该商品当天触发过清零规则并且当前库存为0，则该商品第二天零点会执行自动补充库存规则。
	SyncCount             int    `json:"sync_count"`                //描述：次日自动补充库存数量
	LimitOpenSyncStockNow bool   `json:"limit_open_sync_stock_now"` //描述：是否触发库存清零规则且当前禁止开放平台接口更新库存： true是，禁止更新库存；false否，允许更新库存。 库存清零规则具体配置请查看stock_config下参数内容。
}
