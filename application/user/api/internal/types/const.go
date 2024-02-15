package types

// jwt 令牌关键字
const (
	JwtKey = "Id"
)

// 错误码
const (
	OK             = 200
	Created        = 201
	Unauthorized   = 401
	Forbidden      = 403
	NotFound       = 404
	PwdInCorrect   = 405
	MoneyNotenough = 406
)
