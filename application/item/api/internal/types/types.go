// Code generated by goctl. DO NOT EDIT.
package types

type ItemReqAndResp struct {
	Brand        string `json:"brand"`
	Category     string `json:"category"`
	CommentCount int    `json:"commentCount"`
	Id           int    `json:"id"`
	Image        string `json:"image"`
	IsAD         bool   `json:"isAd"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Sold         int    `json:"sold"`
	Spec         string `json:"spec"`
	Status       int    `json:"status"`
	Stock        int    `json:"stock"`
}

type FindItemsByIdReq struct {
	Ids []string `form:"ids"`
}

type FindItemsByIdResp struct {
	Data []ItemDTO `json:"data, optional"`
}

type ItemDTO struct {
	Brand        string `json:"brand"`
	Category     string `json:"category"`
	CommentCount int    `json:"commentCount"`
	Id           int    `json:"id"`
	Image        string `json:"image"`
	IsAD         bool   `json:"isAd",gorm:"default:true;column:isAD"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Sold         int    `json:"sold"`
	Spec         string `json:"spec"`
	Status       int    `json:"status"`
	Stock        int    `json:"stock"`
}

type QueryItemPageReq struct {
	IsAsc    string `form:"isAsc, optional"`
	PageNo   int    `form:"pageNo, optional"`
	PageSize int    `form:"pageSize, optional"`
	SortBy   string `form:"sortBy, optional"`
}

type QueryItemPageResp struct {
	List  []ItemDTO `json:"list"`
	Pages int       `json:"pages"`
	Total int       `json:"total"`
}

type UpdateItemStatusReq struct {
	Id     int `path:"id"`
	Status int `path:"status"`
}

type DeductItemsReq struct {
	Order []OrderDetailDTO `json:"order"`
}

type OrderDetailDTO struct {
	ItemId int `json:"itemId"`
	Num    int `json:"num"`
}

type FindItemByIdReq struct {
	Id int `path:"id"`
}

type DelItemByIdReq struct {
	Id int `path:"id"`
}
