package shangou

type ActAllGetByAppFoodCodesResponseDataActListItem struct {
	AppFoodCode    string `json:"app_food_code"`    //描述：商品app_food_code
	TypeName       string `json:"type_name"`        //描述：活动类型名称
	Type           int    `json:"type"`             //描述：活动类型：2-满减；17-折扣商品；20-第二份半价；23-买赠；27-指定商品满减；40-加价换购；43-X元M件；56-爆品；
	Status         int    `json:"status"`           //描述：活动状态：0-活动未开始；1-活动已生效；5-活动冻结中
	StartTime      int    `json:"start_time"`       //描述：活动开始时间，10位秒级时间戳
	EndTime        int    `json:"end_time"`         //描述：活动结束时间，10位秒级时间戳
	ActId          int64  `json:"act_id"`           //描述：活动id
	CanModifyPrice int    `json:"can_modify_price"` //描述：商品价格是否可修改：0-可以修改；1-不能修改
}
