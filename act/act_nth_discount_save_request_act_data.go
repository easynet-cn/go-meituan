package act

type ActNthDiscountSaveRequestActData struct {
	ActId          *int64                                            `json:"act_id"`           //是否必须：否；描述：活动id，此参数有值时认为是更新操作
	ActName        string                                            `json:"act_name"`         //是否必须：创建时必填，不支持更新；描述：1）字段描述：促销活动主题。仅为便于商家标识活动，不展示在C端。 2）长度：12字符。
	StartTime      *int                                              `json:"start_time"`       //是否必须：创建时必填，不支持更新；描述：1）字段描述：活动开始时间。 2）格式：10位秒级的时间戳。 3）开始时间不得早于当前日期 4）时间自动换算成所传日期的00:00:00。
	EndTime        *int                                              `json:"end_time"`         //是否必须：创建时必填，不支持更新；描述：1）字段描述：活动结束时间。 2）格式：10位秒级的时间戳。 3）结束时间不得早于当前时间和开始时间。 4）时间自动换算成所传日期的23:59:59。
	ActSettingType *int                                              `json:"act_setting_type"` //是否必须：创建时必填，不支持更新；描述：1）字段描述：活动子类型选择。 2）取值范围：0-第N件X折；1-第N件X元。 3）活动信息的字段规则根据此字段而变化。
	ActDetails     []ActNthDiscountSaveRequestActDataActDetailsItem  `json:"act_details"`      //是否必须：创建时必填，不支持更新；描述：1）字段描述：活动具体信息。 2）至少有一阶，最多三阶。
	ActProducts    []ActNthDiscountSaveRequestActDataActProductsItem `json:"act_products"`     //是否必须：创建时必填，不支持更新；描述：1）字段描述：参加活动的商品信息集合。 2）每个门店内参加第N件优惠活动的商品个数限制为200个。
}
