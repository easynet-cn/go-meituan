package shangou

//药品限购信息
type LimitSaleInfo struct {
	LimitSale bool   `json:"limitSale"` //是否必须：是；描述：是否限制购买数量： (1)字段取值范围：1.true开启；false关闭。(2)如开启限制，则对SPU下所有SKU生效，每个SKU均按设置规则执行限购。(3)如入参选择“不限制”，则后续参数入参也无效。(4)如入参选择“限制”，则开启单个买家限购功能。
	Type      int    `json:"type"`      //是否必须：如限购开启则必填；描述：限购规则枚举值：1,2。不传默认为1； 1代表限制开始结束周期内每X天内下单顾客限购数量，X为frequency字段值，2代表限制整个开始结束周期内下单顾客限购数量。
	Frequency int    `json:"frequency"` //是否必须：否；描述：限购循环天数： 最大31，最小1。
	Begin     string `json:"begin"`     //是否必须：如限购开启则必填；描述：限购开始日期： 开始时间不得早于当前日期的7天前。
	End       string `json:"end"`       //是否必须：如限购开启则必填；描述：限购结束日期： 结束日期不得早于开始日期，且与开始日期相距不得超出90天。
	Count     int    `json:"count"`     //是否必须：如限购开启则必填；描述：限购数量： (1)字段范围：整数，最小为1。(2)如入参为小数，则向下取整（例如：入参数量为2.8，最终取2）。
}