package shangou

type MedicineStockRequestMedicineData struct {
	AppPoiCode      string `json:"app_poi_code"`      //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。须与外层的app_poi_code字段信息一致。
	AppMedicineCode string `json:"app_medicine_code"` //是否必须：是；描述：APP方药品id，即商家中台系统里药品的编码（spu_code值）
	Stock           int    `json:"stock"`             //是否必须：是；描述：药品的库存，传非负整数，最大支持99999
}
