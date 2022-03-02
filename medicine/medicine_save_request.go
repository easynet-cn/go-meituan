package medicine

//创建药品
//请求地址：https://waimaiopen.meituan.com/api/v1/medicine/save
type MedicineSaveRequest struct {
	AppPoiCode      string  `url:"app_poi_code"`      //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	AppMedicineCode string  `url:"app_medicine_code"` //是否必须：是；描述：APP方药品id，可使用商家中台系统里药品的编码（spu_code值）。 (1) 不同门店之间药品id可以重复，同一门店内药品id不允许重复； (2)字段信息限定长度不超过128个字符。 (3)此字段信息也会默认同步至skuid字段上，即同步至商家端后台药品详情中的“店内码/货号”字段上。
	Upc             string  `url:"upc"`               //是否必须：是；描述：药品UPC码，同一门店内药品UPC码不允许重复。
	MedicineNo      string  `url:"medicine_no"`       //是否必须：否；描述：药品的批准文号（国药准字号），系统会根据药品UPC码匹配美团标准库中药品的批准文号信息来展示。
	Spec            string  `url:"spec"`              //是否必须：否；描述：药品规格，系统会根据药品UPC码匹配美团标准库中药品的规格信息来展示。
	Price           float64 `url:"price"`             //是否必须：是；描述：药品售卖价格，单位是元；最多支持2位小数，不能为负数，最大支持30000.00
	Stock           int64   `url:"stock"`             //是否必须：是；描述：药品在当前门店的可售卖库存，不能为负数或小数，最大允许99999
	CategoryCode    string  `url:"category_code"`     //是否必须：与category_name二选一填入；描述：分类code：(1)此字段信息需传入门店内存在的末级分类的分类code。(2)一次调用可以传多个分类，最多不超过5个分类，用英文逗号隔开。(3)商家未传递分类code时，根据分类名称查询或操作分类
	CategoryName    string  `url:"category_name"`     //是否必须：与category_code二选一填入；描述：分类名称：(1)此字段信息需传入门店内已存在的末级分类的分类名称。(2)一次调用可以传多个分类，最多不超过5个分类，用英文逗号隔开。(3)商家未传递分类code时，根据分类名称查询或操作分类
	IsSoldOut       int     `url:"is_sold_out"`       //是否必须：否；描述：药品上下架状态。字段取值范围：0上架，1下架。 本参数如传的不是0,1或不传，则默认为下架状态。
	Sequence        int     `url:"sequence"`          //是否必须：否；描述：药品在所属末级分类下的排序序号，同一分类下药品排序序号数字越小，前端排位越靠前。
	LimitSaleInfo   string  `url:"limit_sale_info"`   //是否必须：否；描述：药品限购详情
}
