package utils

const (
	CacheOrderKey  = "cache#order"
	CacheOrderTime = 3600
)

// 订单状态
const (
	NotPayment        = 1 // 未付款
	Paied             = 2 // 已付款
	Shipped           = 3 // 已发货
	Receipted         = 4 // 确认收货
	TradeCancellation = 5 //交易取消
	TradeEnd          = 6 //交易结束
)

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
