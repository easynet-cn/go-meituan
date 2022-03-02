package act

type ActMarkupRrepurchaseBatchsaveRequestActData struct {
	ActId     *int64                                                    `json:"act_id"`     //是否必须：否，更新时必填；描述：活动id，仅更新时填写。本接口支持添加商品至该活动中；支持更新该活动中商品的活动价、活动库存、每单限购的信息。
	ActName   *string                                                   `json:"act_name"`   //是否必须：否，创建时必填；描述：活动名称，不超过12个字符。更新活动时不支持更新此活动名称。
	StartTime *int                                                      `json:"start_time"` //是否必须：否，创建时必填；描述：活动开始时间，传10位秒级的时间戳。(1)系统会自动将传入的时间转换为所传日期的00:00:00。(2)活动开始时间不得早于当前时间。(3)更新活动时不支持更新此时间。
	EndTime   *int                                                      `json:"end_time"`   //是否必须：否，创建时必填；描述：活动结束时间，传10位秒级的时间戳。(1)系统会自动将传入的时间转换为所传日期的23:59:59。(2)活动结束时间不得早于当前时间和开始时间。(3)更新活动时不支持更新此时间。
	ActPrice  *float64                                                  `json:"act_price"`  //是否必须：否，创建时必填；描述：换购门槛，表示可参加本次换购活动的订单金额(不包含运费)，单位元。须为1～999999.99之间的数字。更新活动时，不支持修改此信息。
	ActRemark *string                                                   `json:"actRemark"`  //是否必须：否；描述：参加换购的活动商品数据，json格式；同一次调用中，可上传的活动商品数据不能超过200组。更新活动时，支持添加活动商品；支持修改已有活动商品的相关信息。 注意：app_food_code即为药品的app_medicine_code。
	AppFoods  []ActMarkupRrepurchaseBatchsaveRequestActDataAppFoodsItem `json:"app_foods"`  //是否必须：否；描述：参加换购的活动商品数据，json格式；同一次调用中，可上传的活动商品数据不能超过200组。更新活动时，支持添加活动商品；支持修改已有活动商品的相关信息。 注意：app_food_code即为药品的app_medicine_code。
}
