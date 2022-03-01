package shangou

type ActDeleteActsByProductsResponse struct {
	Data      string                                         `json:"data"`       //描述：ok表示成功(部分成功或全部成功)，ng全部失败
	ErrorList []ActDeleteActsByProductsResponseErrorListItem `json:"error_list"` //描述：部分失败时的一些提示信息
}
