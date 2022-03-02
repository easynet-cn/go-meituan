package comment

type CommentQueryResponse struct {
	Data []CommentQueryRequestDataItem
	Code int    //描述：错误码
	Msg  string //描述：报错信息
}
