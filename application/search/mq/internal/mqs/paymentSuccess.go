package mqs

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/search/mq/internal/svc"
	"log"
	"strconv"
)

type PaymentSuccess struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccess {
	return &PaymentSuccess{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaymentSuccess) Consume(_, item string) error {
	var res SearchItemDTO
	if err := json.Unmarshal([]byte(item), &res); err != nil {
		logx.Errorf("json.Unmarshal: %v,error : %v", item, err)
		return err
	}

	resp, err := esapi.IndexRequest{
		Index:      "items",
		DocumentID: strconv.Itoa(int(res.Id)),
		Body:       bytes.NewReader([]byte(item)),
		Refresh:    "true",
	}.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		logx.Errorf("esapi.IndexRequest.Do, error : %v", err)
		return err
	}
	log.Println(resp)
	return nil
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
