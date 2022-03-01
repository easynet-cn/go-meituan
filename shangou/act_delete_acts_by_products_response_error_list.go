package shangou

type ActDeleteActsByProductsResponseErrorListItem struct {
	Code        int    `json:"code"`          //描述：错误码
	Msg         string `json:"msg"`           //描述：错误信息
	AppFoodCode string `json:"app_food_code"` //描述：错误的app_food_code
}
