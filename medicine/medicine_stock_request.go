package medicine

//批量更新药品库存
//请求地址：https://waimaiopen.meituan.com/api/v1/medicine/stock
type MedicineStockRequest struct {
	AppPoiCode   string `url:"app_poi_code"`  //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	MedicineData string `url:"medicine_data"` //是否必须：是；描述：药品库存数据集合的json格式数组， (1)上传参数包括： a）app_medicine_code(必填项)，APP方药品id，即商家中台系统里药品的编码（spu_code值）。 b）app_poi_code(必填项)，APP方门店id，即商家中台系统里门店的编码。须与外层的app_poi_code字段信息一致。 c）stock(必填项)，为药品的库存，传非负整数。 (2)可传药品库存数据限定不超过200组。 (3)支持药品库存部分同步成功，部分同步失败。
}
