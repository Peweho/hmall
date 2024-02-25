package utils

// 支付单状态
const (
	NotCommit  = iota // 待提交
	NotPay            // 待支付
	PayTimeOut        // 支付超时
	PaySuccess        // 支付成功
)

// 支付单逻辑删除
const (
	NotDelete = iota
	Deleted
)
