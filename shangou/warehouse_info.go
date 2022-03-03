package shangou

type WarehouseInfo struct {
	WasehouseCode string `json:"wasehouse_code"` //描述：商家系统中的仓库编码
	Count         int    `json:"count"`          //描述：该仓库下对应的商品数量
}
