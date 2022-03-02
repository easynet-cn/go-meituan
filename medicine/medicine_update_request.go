package medicine

//更新药品
//请求地址：https://waimaiopen.meituan.com/api/v1/medicine/update
type MedicineUpdateRequest struct {
	AppPoiCode      string  `url:"app_poi_code"`      //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	AppMedicineCode string  `url:"app_medicine_code"` //是否必须：是；描述：APP方药品id，即商家中台系统里药品的编码（spu_code值）。
	Upc             string  `url:"upc"`               //是否必须：否；描述：药品UPC码，展示美团标准库中此药品的UPC码，不支持修改。
	MedicineNo      string  `url:"medicine_no"`       //是否必须：否；描述：药品的批准文号（国药准字号），展示美团标准库中此药品的批准文号信息，不支持修改。
	Spec            string  `url:"spec"`              //是否必须：否；描述：药品规格，展示美团标准库中此药品的规格信息，不支持修改。
	Price           float64 `url:"price"`             //是否必须：否；描述：药品售卖价格，单位是元；最多支持2位小数，不能为负数，最大支持30000.00
	Stock           int     `url:"stock"`             //是否必须：否；描述：药品在当前门店的可售卖库存，不能为负数或小数，最大支持99999
	CategoryCode    string  `url:"category_code"`     //是否必须：与category_name二选一填入 ；描述：分类code：(1)此字段信息需传入门店内存在的末级分类的分类code。(2)一次调用可以传多个分类，最多不超过5个分类，用英文逗号隔开。(3)此字段所传入的分类信息将覆盖商品所有原分类。(4)商家未传递分类code时，根据分类名称查询或操作分类
	CategoryName    string  `url:"category_name"`     //是否必须：与category_code二选一填入 ；描述：分类名称：(1)此字段信息需传入门店内已存在的末级分类的分类名称。(2)一次调用可以传多个分类，最多不超过5个分类，用英文逗号隔开。(3)此字段所传入的分类信息将覆盖商品所有原分类(根据药品分类code信息定位)。(4)商家未传递分类code时，根据分类名称查询或操作分类
	IsSoldOut       int     `url:"is_sold_out"`       //是否必须：否；描述：药品上下架状态，字段取值范围：0上架，1下架。
	Sequence        int     `url:"sequence"`          //是否必须：否；描述：药品在所属末级分类下的排序序号，同一分类下药品排序序号数字越小，前端排位越靠前。
	LimitSaleInfo   string  `url:"limit_sale_info"`   //是否必须：否；描述：药品限购详情
}
