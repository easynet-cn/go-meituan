package shangou

//门店药品列表项
type MedicineListItem struct {
	Name            string  `json:"name"`              //描述：药品名称，展示在美团标准库中此药品的名称。
	AppMedicineCode string  `json:"app_medicine_code"` //描述：APP方药品id，即商家中台系统里药品的编码（spu_code值）。
	Upc             string  `json:"upc"`               //描述：药品UPC码，展示美团标准库中此药品的UPC码。
	MedicineNo      string  `json:"medicine_no"`       //描述：药品的批准文号（国药准字号），展示美团标准库中此药品的批准文号信息。
	Spec            string  `json:"spec"`              //描述：药品规格，展示美团标准库中此药品的规格信息。
	Price           float64 `json:"price"`             //描述：药品售卖价格，单位是元。
	Stock           int     `json:"stock"`             //描述：药品在当前门店的可售卖库存，若此字段信息为空，则表示当前商品库存为“无限”。
	CategoryCode    string  `json:"category_code"`     //描述：分类code，展示当前药品所在末级分类的分类code列表，多个分类code以英文逗号隔开。
	CategoryName    string  `json:"category_name"`     //描述：分类名称，展示当前药品所在末级分类的分类名称列表，多个分类名称以英文逗号隔开。
	IsSoldOut       int     `json:"is_sold_out"`       //描述：药品当前的上下架状态：0上架；1下架。
	Sequence        int     `json:"sequence"`          //描述：药品在所属末级分类下的排序序号，同一分类下药品排序序号数字越小，前端排位越靠前。
	Ctime           int     `json:"ctime"`             //描述：创建药品的时间，为10位秒级的时间戳。
	Utime           int     `json:"utime"`             //描述：最近一次更新药品的时间，为10位秒级的时间戳。
	MedicineType    int     `json:"medicine_type"`     //描述：药品的处方类型，1OTC；2非药品；3处方药；4通用品。
	LimitSaleInfo   string  `json:"limit_sale_info"`   //描述：药品限购详情
	StockConfig     string  `json:"stock_config"`      //描述：无货取消订单触发库存清零相关配置
}
