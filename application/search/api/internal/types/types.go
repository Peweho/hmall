// Code generated by goctl. DO NOT EDIT.
package types

type SearchListReq struct {
	Brand    string `form:"brand,optional"`
	Category string `form:"category,optional"`
	IsAsc    string `form:"isAsc,optional"`
	Key      string `form:"key,optional"`
	MaxPrice int    `form:"maxPrice,optional"`
	MinPrice int    `form:"minPrice,optional"`
	PageNo   int    `form:"pageNo,optional"`
	PageSize int    `form:"pageSize,optional"`
	SortBy   string `form:"sortBy,optional"`
}

type SearchListResp struct {
	Items []SearchItemDTO `json:"items"`
	Pages int             `json:"pages"`
	Total int             `json:"total"`
}

type SearchItemDTO struct {
	Brand        string `json:"brand"`
	Category     string `json:"category"`
	CommentCount int64  `json:"commentCount"`
	Id           int64  `json:"id"`
	Image        string `json:"image"`
	IsAD         bool   `json:"isAD"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Sold         int64  `json:"sold"`
	Spec         string `json:"spec"`
	Status       int64  `json:"status"`
	Stock        int64  `json:"stock"`
}

type GetHotItemsReq struct {
	Num int `form:"num"`
}

type GetHotItemsResp struct {
	Items []SearchItemDTO `json:"items"`
}
