package isv

//快捷接口（直接获取access_token接口）
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/oauth/authorize
//文档地址：https://open-shangou.meituan.com/home/guide/medical/10686
type GetAccessTokenRequest struct {
	AppPoiCode   string `url:"app_poi_code"`  //是否必须：是；描述：门店代码。
	ResponseType string `url:"response_type"` //是否必须：是；描述：授权方式，此处response_type=token。
}
