package shangou

type DistanceMarkupFactor struct {
	Distance                float64 `json:"distance"`                   //是否必须：是；描述：超限距离， 1.单位km 2.正数且最多支持1位小数 3.最大支持传1000 4.多个规则之间不能重复
	MarkupNum               float64 `json:"markupNum"`                  //是否必须：是；描述：加价金额， 1.单位元 2.正数且最多支持1位小数 3.最大支持传10000
	WeightMarkupExecuteType int     `json:"weight_markup_execute_type"` //是否必须：否；描述：按重量加价规则的执行类型：1启用，2不启用，3保持原值
}
