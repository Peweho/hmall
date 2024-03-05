package types

const (
	CacheCartKey  = "cache#order"
	CacheCartTime = 3600
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

// 订单状态
const (
	NotPayment        = 1 // 未付款
	Paied             = 2 // 已付款
	Shipped           = 3 // 已发货
	Receipted         = 4 // 确认收货
	TradeCancellation = 5 //交易取消
	TradeEnd          = 6 //交易结束
)

// 秒杀商品相关
const (
	// 秒杀商品状态
	FalshNotStart = iota
	FalshStart
	FalshEndDecut
	FalshEndNotDecut

	// 秒杀商品超时时间
	FalshTimeOut     = 1
	CacheFlashStatus = "item#flash"
)
