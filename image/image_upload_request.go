package image

//图片上传API
type ImageUploadRequest struct {
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	ImgData    []byte `url:"img_data"`     //是否必须：是；描述：图片字节流，图片转成 byte 数组的格式。注意：此字段不参与签名计算。
	ImgName    string `url:"img_name"`     //是否必须：是；描述：图片名称，图片格式支持 [jpg][jpeg][png][bmp]
}
