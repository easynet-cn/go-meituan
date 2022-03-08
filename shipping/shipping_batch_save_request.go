package shipping

//批量创建/更新配送范围
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/shipping/batchsave
//文档地址：https://open-shangou.meituan.com/home/docDetail/48
type ShippingBatchSaveRequest struct {
	AppPoiCode   string `url:"app_poi_code"`  //是否必须：是；描述：APP方门店id，传商家中台系统里门店的编码，最长不超过128个字符。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ShippingData string `url:"shipping_data"` //是否必须：是；描述：配送范围数据集合，json格式数组。 处理逻辑：(1)如果配送范围(app_shipping_code)原来不存在就新增。(2)如app_shipping_code存在就更新配送范围。(3)如原来存在的app_shipping_code，而本次没上传，则会删除此配送范围。（注意，不支持删除非接口创建的配送范围。）(4)配送范围都为多边形，不支持圆形。
}
