package medicine

//批量创建药品
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/medicine/batchsave
//文档地址：https://open-shangou.meituan.com/home/docDetail/86
type MedicineBatchSaveRequest struct {
	AppPoiCode   string `url:"app_poi_code"`  //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	MedicineData string `url:"medicine_data"` //是否必须：是；描述：多个药品数据集合的json格式数组：(1)可传参数及逻辑与单个接口【medicine/save 创建药品】中各参数规则一致，请查看单个接口文档的详细描述。(2)可传药品数据限定不能超过200组。(3)支持药品部分同步成功，部分同步失败。
}
