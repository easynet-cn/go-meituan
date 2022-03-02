package medicine

//门店药品分类
type MedicineCategory struct {
	AppPoiCode         string `json:"app_poi_code"`         //描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	CategoryCode       string `json:"category_code"`        //描述：一级分类code
	CategoryName       string `json:"category_name"`        //描述：一级分类名称
	Sequence           int    `json:"sequence"`             //描述：一级分类在当前门店内的排序序号，数字越小，前端排位越靠前。
	SecondCategoryCode string `json:"second_category_code"` //描述：二级分类code
	SecondCategoryName string `json:"second_category_name"` //描述：二级分类名称
	SecondSequence     int    `json:"second_sequence"`      //描述：二级分类在所属一级分类下的排序序号，数字越小，前端排位越靠前。
	CTime              int    `json:"ctime"`                //描述：创建一级分类的时间，为10位秒级的时间戳。
	UTime              int    `json:"utime"`                //描述：最近一次修改一级分类的时间，为10位秒级的时间戳。
}
