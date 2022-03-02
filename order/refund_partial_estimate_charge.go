package order

type RefundPartialEstimateCharge struct {
	TotalFoodAmount       string `json:"total_food_amount"`       //描述：本次退款商品的总金额，不含包装盒费，单商品总价=商品原价+赠品原价，位元
	BoxAmount             string `json:"box_amount"`              //描述：本次退款商品包装盒费总价，单位元
	ActivityPoiAmount     string `json:"activity_poi_amount"`     //描述：本次退款商家活动费用总支出成本，含赠品成本，单位元
	ActivityMeituanAmount string `json:"activity_meituan_amount"` //描述：本次退款美团活动补贴总金额，单位元
	ActivityAgentAmount   string `json:"activity_agent_amount"`   //描述：本次退款代理商活动承担金额，单位元
	PlatformChargeFee     string `json:"platform_charge_fee"`     //描述：本次退款平台服务费总金额，单位元
	SettleAmount          string `json:"settle_amount"`           //描述：结算金额，单位元
	Productpreferences    string `json:"productpreferences"`      //描述：商家活动支出分摊到商品上的优惠总金额，单位元
	NotProductpreferences string `json:"not_productpreferences"`  //描述：商家活动支出未分摊到商品上的总金额，单位元
}
