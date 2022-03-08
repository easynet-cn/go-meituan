package poi

//查询结算银行信息
//请求方式：GET
//请求地址：https://waimaiopen.meituan.com/api/v1/ecommerce/poi/settle/settlement/bankInfo
//文档地址：https://open-shangou.meituan.com/home/docDetail/678
type PoiSettleSettlementBankInfoRequest struct {
	BankId int `url:"bank_id"` //是否必须：否；描述：银行id，不传则查询银行信息，查询银行支行id本字段必传
	CityId int `url:"city_id"` //是否必须：否；描述：城市id,不传则查询银行信息，查询银行支行id本字段必传，需传入市区城市id，参考【查询城市信息】接口，如果为直辖市，则使用直辖市的省编码来查询，如：110000（北京市）、120000（天津市）、310000（上海市）、500000（重庆市）
}
