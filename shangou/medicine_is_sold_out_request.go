package shangou

//药品批量上下架
//请求地址：https://waimaiopen.meituan.com/api/v1/medicine/isSoldOut
type MedicineIsSoldOutRequest struct {
	AppPoiCode   string `url:"app_poi_code"`  //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	MedicineData string `url:"medicine_data"` //是否必须：是；描述：药品数据集合的json格式数组， (1)上传参数包括： a）app_medicine_code(必填项)，APP方药品id，即商家中台系统里药品的编码（spu_code值）。 b）is_sold_out(必填项)，药品上下架状态；取值范围：0上架，1下架。 (2)可传药品数据限定不超过200组；接口支持部分成功。
}
