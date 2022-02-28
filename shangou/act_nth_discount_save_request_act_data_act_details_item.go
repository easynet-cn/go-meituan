package shangou

type ActNthDiscountSaveRequestActDataActDetailsItem struct {
	Index    int     `json:"index"`    //是否必须：创建时必填，不支持更新；描述：1）字段描述：活动商品件数。 2）取值范围：N：1≤N≤5。
	Discount float64 `json:"discount"` //是否必须：创建时必填，不支持更新；描述：1）字段描述：活动商品折扣系数。 2）取值范围：0≤X≤9.99，最多支持两位小数。 3）此范围当“活动子类型”为0时有效。
	Price    float64 `json:"price"`    //是否必须：创建时必填，不支持更新；描述：1）字段描述：活动商品售价 2）取值范围：0≤X，最多支持两位小数。 3）此范围当“活动子类型”为1时有效。
}
