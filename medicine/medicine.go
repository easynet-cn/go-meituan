package medicine

type Medicine struct {
	AppMedicineCode string  `json:"app_medicine_code"` //是否必须：是；描述：APP方药品id，可使用商家中台系统里药品的编码（spu_code值）。 (1) 不同门店之间药品id可以重复，同一门店内药品id不允许重复； (2)字段信息限定长度不超过128个字符。 (3)此字段信息也会默认同步至skuid字段上，即同步至商家端后台药品详情中的“店内码/货号”字段上。
	Upc             string  `json:"upc"`               //是否必须：是；描述：药品UPC码，同一门店内药品UPC码不允许重复。
	Price           float64 `json:"price"`             //是否必须：是；描述：药品价格，最大允许30000.00
	Stock           int     `json:"stock"`             //是否必须：是；描述：药品库存,最大允许99999
	CategoryCode    string  `json:"category_code"`     //是否必须：与category_name二选一填入；描述：分类code：(1)此字段信息需传入门店内存在的末级分类的分类code。(2)一次调用可以传多个分类，最多不超过5个分类，用英文逗号隔开。(3)商家未传递分类code时，根据分类名称查询或操作分类
	CategoryName    string  `json:"category_name"`     //是否必须：与category_code二选一填入；描述：分类名称：(1)此字段信息需传入门店内已存在的末级分类的分类名称。(2)一次调用可以传多个分类，最多不超过5个分类，用英文逗号隔开。(3)商家未传递分类code时，根据分类名称查询或操作分类
	Sequence        int     `json:"sequence"`          //是否必须：是；描述：一级分类在当前门店内的排序序号，数字越小，前端排位越靠前。
	LimitSaleInfo   string  `json:"limit_sale_info"`   //是否必须：是；描述：限购信息
}
