package shangou

type ActFullDiscountBatchsaveRequestActInfo struct {
	ActName   string `json:"act_name"`   //是否必须：是；描述：活动名称, 仅限输入汉字、数字和字母，且限制长度不超过12个字符（6个汉字）
	StartTime int    `json:"start_time"` //是否必须：是；描述：活动开始时间(必填项)，传入10位秒级的时间戳；系统会自动将传入的时间转换为所传日期的00:00:00；活动开始时间不得早于当前时间。
	EndTime   int    `json:"end_time"`   //是否必须：是；描述：活动结束时间，为10位秒级的时间戳；系统会自动将传入的时间转换为所传日期的23:59:59；活动结束时间不得早于当前时间和开始时间。
	ActType   int    `json:"act_type"`   //是否必须：否；描述：活动类型，取值范围：0全店满减；1指定商品满减。如此字段不传则默认act_type=1，表示指定商品满减类型；act_type=0时，不可以传app_foods字段。
	WeeksTime string `json:"weeks_time"` //是否必须：否；描述：生效活动周期, 1、仅当act_type传0（全店满减）时，支持传本参数； 2、取值范围及意义：1,2,3,4,5,6,7；分别表示周一至周日，多个星期之间用英文逗号分隔； 3、非必传，创建时如不传，则默认周期为"1,2,3,4,5,6,7"。 4、上传的周期相互不能重复。
	Period    string `json:"period"`     //是否必须：否；描述：生效时间段， 1、仅当act_type传0（全店满减）时，支持传本参数； 2、精确到分钟，24小时制，格式为HH:MMHH:MM，最多支持上传3个时段，多个时段之间用英文逗号分隔； 3、每个时段的开始时间必须小于结束时间。传的多个时段不允许有重叠。 4、上传多个时段时，必须按时间先后顺序进行上传。每个时段最短为30分钟。 5、非必传，创建时如不传则默认时段为"00:00:23:59"。
	QutoDelay int    `json:"auto_delay"` //是否必须：否；描述：每次活动到期后自动延期30天，传1表示开启。 1、仅当act_type传0（全店满减）时，支持传本参数。 2、仅支持传1，表示开启，每当活动到期时自动延期30天。 3、非必传，创建活动时如不传则为关闭状态，不自动延期。
}
