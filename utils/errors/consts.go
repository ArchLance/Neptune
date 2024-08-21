package errors

// 错误码
const (
	CodeUnknown        = 9999  // 未知错误
	CodePermissionDeny = 10001 // 权限不足
	CodeRequestFailed  = 10003 // 请求失败
	CodeParamInvalid   = 10005 // 非法参数
	CodeDbError        = 10007 // 数据库错误
	CodeLogicError     = 10009 // 基本业务逻辑错误
	CodeDataNotFound   = 10011 // 数据查找失败
	CodeDataExist      = 10013
)
