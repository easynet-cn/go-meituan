package isv

type GetAccessTokenResponse struct {
	Status       int    `json:"status"`        //描述：表示请求成功或失败，参考错误码。
	AccessToken  string `json:"access_token"`  //描述：获取的access token。
	ExpiresIn    int64  `json:"expires_in"`    //描述：过期时间，单位秒。
	RefreshToken string `json:"refresh_token"` //描述：用于刷新access token的refresh token。
	ReExpiresIn  int64  `json:"re_expires_in"` //描述：过期时间，单位秒。
	State        string `json:"state"`         //描述：请求时传递了该参数，则会原样返回。
	Message      string `json:"message"`       //描述：错误描述。
}
