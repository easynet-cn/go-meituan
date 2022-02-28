package shangou

//【X元M件】批量删除X元M件活动
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/act/item/bundles/batchdelete
//文档地址：https://open-shangou.meituan.com/home/docDetail/332
type ActItemBundlesBatchdeleteRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ActIds     string `url:"act_ids"`      //是否必须：是；描述：活动id，同一次调用中，支持上传待删除的活动id不能超过100个，以英文逗号分隔。不支持部分活动删除成功。
}
