package act

type ActNthDiscountProductsResponse struct {
	Data  string `json:"data"`  //描述：查询成功，则为活动商品信息集合 查询失败，则为 ng
	Error string `json:"error"` //描述：失败信息
}
