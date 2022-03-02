package medicine

type MedicineIsSoldOutRequestMedicineData struct {
	AppMedicineCode string `json:"app_medicine_code"` //是否必须：是；描述：APP方药品id，即商家中台系统里药品的编码
	IsSoldOut       int    `json:"is_sold_out"`       //是否必须：是；描述：药品上下架状态；取值范围：0上架，1下架
}
