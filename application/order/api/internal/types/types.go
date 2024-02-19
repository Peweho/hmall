// Code generated by goctl. DO NOT EDIT.
package types

type CreateOrdeReq struct {
	AddressId   int         `json:"addressId"`
	PaymentType int         `json:"paymentType"`
	Details     []DetailDTO `json:"details"`
}

type DetailDTO struct {
	ItemId int `json:"itemId"`
	Num    int `json:"num"`
}

type FindOrderByIdReq struct {
	Id int `path:"id"`
}

type FindOrderByIdResp struct {
	FindOrderByIdVO
}

type FindOrderByIdVO struct {
	Id          int    `json:"id"`
	PayTime     string `json:"payTime"`
	PaymentType int    `json:"paymentType"`
	Status      int    `json:"status"`
	TotalFee    int    `json:"totalFee"`
	UserId      int    `json:"userId"`
	CloseTime   string `json:"closeTime"`
	CommentTime string `json:"commentTime"`
	ConsignTime string `json:"consignTime"`
	CreateTime  string `json:"createTime"`
	EndTime     string `json:"endTime"`
}

type MarkOrderReq struct {
	OrderId string `path:"orderId"`
}