package act

type ActRetailDiscountBatchsaveResponseExtraInfo struct {
	TaskId         string `json:"task_id"`          //描述：当商家使用异步排队方式请求接口时返回，商家可使用任务id查询任务进度
	TaskDetailLink string `json:"task_detail_link"` //描述：创建异步排队任务成功，返回任务详情链接，通过浏览器打开链接查看任务处理结果
}
