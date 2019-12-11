package errors

type BaseErrMsg struct {
	ErrMsg   string `json:"err_msg"`
	ErrMsgEn string `json:"err_msg_en"`
}

type ErrCode int

var (
	baseErrors = map[ErrCode]BaseErrMsg{}
)

const (
	ErrSuccess ErrCode = iota
	ErrUnknownError
	ErrUnstableNetwork
	ErrPermissionDeny
	ErrServiceUnderMaintaining
	ErrTooMuchRequest
	ErrServiceNotFound
)

func init() {
	// err code definition rule:
	// 	0				System Reserved: Success
	//  1 - 99			System Reserved: Basic Error

	//  1000 -   	Business Logic

	baseErrors[0] = BaseErrMsg{ErrMsg: "成功", ErrMsgEn: "success"}

	baseErrors[1] = BaseErrMsg{ErrMsg: "未知错误 %s", ErrMsgEn: "unknown error"}
	baseErrors[2] = BaseErrMsg{ErrMsg: "网络波动, 请重新尝试", ErrMsgEn: "unstable network connection, please retry"}
	baseErrors[3] = BaseErrMsg{ErrMsg: "权限不足", ErrMsgEn: "permission denied"}
	baseErrors[4] = BaseErrMsg{ErrMsg: "服务维护中, 请稍后", ErrMsgEn: "service maintaining, please wait"}
	baseErrors[5] = BaseErrMsg{ErrMsg: "访问量过大, 请稍后重试", ErrMsgEn: "router under flow control"}
	baseErrors[6] = BaseErrMsg{ErrMsg: "请求的服务不存在", ErrMsgEn: "service not found"}
	baseErrors[7] = BaseErrMsg{ErrMsg: "缺少参数: %s", ErrMsgEn: "lack of parameter"}
	baseErrors[8] = BaseErrMsg{ErrMsg: "缺少参数", ErrMsgEn: "lack of parameter"}
	baseErrors[9] = BaseErrMsg{ErrMsg: "%s", ErrMsgEn: "invalid parameter"}
	baseErrors[10] = BaseErrMsg{ErrMsg: "您的请求过于频繁, 请稍后再试", ErrMsgEn: "request frequency control"}
	baseErrors[11] = BaseErrMsg{ErrMsg: "访问受限: %s", ErrMsgEn: "request frequency control [black list]"}
	baseErrors[12] = BaseErrMsg{ErrMsg: "不支持的方法", ErrMsgEn: "method not support"}
	baseErrors[13] = BaseErrMsg{ErrMsg: "JSON参数错误 %s", ErrMsgEn: "json param format error"}
	baseErrors[14] = BaseErrMsg{ErrMsg: "操作正在执行中，请稍后重试", ErrMsgEn: "operation is working, waiting"}
	baseErrors[15] = BaseErrMsg{ErrMsg: "JSON解析出错 %s", ErrMsgEn: "json parse error"}
	baseErrors[16] = BaseErrMsg{ErrMsg: "函数栈溢出", ErrMsgEn: "func stack overflow"}

	baseErrors[20] = BaseErrMsg{ErrMsg: "类型错误: %s", ErrMsgEn: "invalid type"}
	baseErrors[30] = BaseErrMsg{ErrMsg: "'%s'长度不得大于%d", ErrMsgEn: "field too long"}

	baseErrors[40] = BaseErrMsg{ErrMsg: "未定义的环境变量: %s", ErrMsgEn: "undefined environment variables"}
	baseErrors[41] = BaseErrMsg{ErrMsg: "无法识别的环境变量: %s=%s", ErrMsgEn: "undefined environment variables"}

	baseErrors[50] = BaseErrMsg{ErrMsg: "数据层类型错误: %s", ErrMsgEn: "failed when transform type on orm"}
	baseErrors[51] = BaseErrMsg{ErrMsg: "数据库操作失败: %s", ErrMsgEn: "db op error"}
	baseErrors[52] = BaseErrMsg{ErrMsg: "无法获取数据库连接: %s", ErrMsgEn: "Failed to get db-connection"}

}
