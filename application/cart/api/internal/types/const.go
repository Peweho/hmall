package types

const CacheItemKey = "cache#item"

// jwt 令牌关键字
const (
	JwtKey = "Id"
)

// 返回码
const (
	OK             = 200
	Created        = 201
	Unauthorized   = 401
	Forbidden      = 403
	NotFound       = 404
	PwdInCorrect   = 405
	MoneyNotenough = 406
)
const (
	MSgAddCache = "0"
	MSgDelCache = "1"
)
