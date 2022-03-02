package medicine

type MedicinePriceRequestMedicineData struct {
	AppMedicineCode string  `json:"app_medicine_code"` //是否必须：是；描述：APP方药品id，即商家中台系统里药品的编码（spu_code值）
	Price           float64 `json:"price"`             //是否必须：是；描述：药品的价格，单位是元；最多支持2位小数，不能为负数，最大支持30000.00
}
