package medicine

//查询门店药品列表请求
//请求地址：https://waimaiopen.meituan.com/api/v1/medicine/list
type MedicineListRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	Offset     int    `url:"offset"`       //是否必须：是；描述：分页查询偏移量，表示从第几页开始查询，需根据公式得到：页码(舍弃小数)=offset/limit+1。例如：offset=23，limit=5，根据公式计算结果表示从第5页开始查询，且每页5个商品，即本次请求结果将展示门店内第21～第25条商品数据。(2)【limit】字段有值时，此字段必填。
	Limit      int    `url:"limit"`        //是否必须：是；描述：分页每页展示药品数量，须为大于0的整数，最多支持200。
}
