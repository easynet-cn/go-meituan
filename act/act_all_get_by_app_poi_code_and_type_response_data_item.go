package act

type ActAllGetByAppPoiCodeAndTypeResponseItem struct {
	ActName   string  `json:"act_name"`   //描述：活动名称
	ActType   int     `json:"act_type"`   //描述：优惠活动类型编码，参考现有“优惠活动类型”type字段
	ActStatus int     `json:"act_status"` //描述：活动状态：0-待生效，1-进行中，5-冻结
	ActCount  int     `json:"act_count"`  //描述：活动数量： 1）门店中此类型活动总数量。 2）非集合类活动：买赠、折扣、爆品活动，活动数量返回：1，代表有活动存在。 3）集合类活动：指定商品满减、加价购、X元M件、第N件优惠返回活动计数。 4）劵类活动：包含品牌券活动数量（非券数量）。
	ItemCount int     `json:"item_count"` //描述：活动内商品总量： 1）参加此类型活动的商品总数量。 2）不计算门店类活动，只计算集合类和商品类的活动商品数量。 3）商品数量为聚合后的数量，即同一个商品设置不同时段活动计为1个。 4）当品牌券设置为upc维度时，无品牌券商品数量；当品牌券设置为sku维度时，返回值包含商品数量（聚合值）。
	StartTime int     `json:"start_time"` //描述：活动开始时间，为此状态下全部活动的最早时间。
	EndTime   int     `json:"end_time"`   //描述：活动结束时间，为此状态下全部活动的最晚时间。
	ActIds    []int64 `json:"act_ids"`    //描述：活动id，非集合类活动：买赠、折扣、爆品活动及劵类活动不返回活动id。
}
