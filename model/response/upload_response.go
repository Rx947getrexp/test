package response

// PathUploadParam 上传图片返回地址
type PathUploadParam struct {
	Url string `json:"url"` //地址
}

type DataResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data PathUploadParam `json:"data"`
}
