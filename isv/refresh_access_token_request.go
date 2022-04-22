package isv

//通过refresh_token刷新access_token的接口
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/oauth/token
//文档地址：https://open-shangou.meituan.com/home/guide/medical/10686
type RefreshAccessTokenRequest struct {
	GrantType    string `url:"grant_type"`    //是否必须：是；描述：授权类型，此处传refresh_token。
	RefreshToken string `url:"refresh_token"` //是否必须：是；描述：授权类型，请求获取access token成功后返回的refresh token。
}
