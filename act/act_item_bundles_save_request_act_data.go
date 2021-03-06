package act

type ActItemBundlesSaveRequestActData struct {
	ActId     *int64                                         `json:"act_id"`     //是否必须：仅更新时填写；描述：活动id，仅更新时填写。本接口支持更新时，修改商品的活动库存；支持添加商品至该活动中。
	ActName   *string                                        `json:"act_name"`   //是否必须：仅创建时填写；描述：活动名称，不超过12个字符，且仅支持汉字、数字和字。更新活动时不支持更新此活动名称。
	StartTime *int                                           `json:"start_time"` //是否必须：仅创建时填写；描述：活动开始时间，传10位秒级的时间戳。(1)系统会自动将传入的时间转换为所传日期的00:00:00。(2)活动开始时间不得早于当前时间。(3)更新活动时不支持更新此时间。
	EndTime   *int                                           `json:"end_time"`   //是否必须：仅创建时填写；描述：活动结束时间，传10位秒级的时间戳。(1)系统会自动将传入的时间转换为所传日期的23:59:59。(2)活动结束时间不得早于当前时间和开始时间。(3)更新活动时不支持更新此时间。
	ActPrice  *float64                                       `json:"act_price"`  //是否必须：仅创建时填写；描述：活动价格，单位元，表示“X元M件”中的X元；须为1~999999.99间数字。更新活动时不支持更新此信息。
	ActMum    *int                                           `json:"act_num"`    //是否必须：仅创建时填写；描述：活动数量，表示“X元M件”中的M件；必须为1~20间整数。更新活动时不支持更新此信息。
	AppFoods  []ActItemBundlesSaveRequestActDataAppFoodsItem `json:"app_foods"`  //是否必须：仅更新时填写；描述：参加活动的商品信息，json格式数组。同一次调用中，支持上传的活动商品数据不能超过200组。注意：参加活动的商品原价必须大于X/M。例如：X=10，M=2，则参加活动的商品原价都必须大于5元。 注意：app_food_code即为药品的app_medicine_code。
}
