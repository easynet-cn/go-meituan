package shangou

type ActNthDiscountInfoResponse struct {
	Data  string `json:"data"`  //描述：查询成功，则为活动信息集合 查询失败，则为 ng 活动信息中各字段与创建更新接口的act_data保持一致，活动查询接口不会返回活动下商品信息，即act_products字段为null
	Error string `json:"error"` //描述：失败信息
}
