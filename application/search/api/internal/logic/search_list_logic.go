package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/search/api/internal/svc"
	"hmall/application/search/api/internal/types"
	"hmall/pkg/util"
	"log"
	"strconv"
	"strings"
)

type SearchListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchListLogic {
	return &SearchListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 假设条件齐全
var query = map[string]any{
	"query": map[string]any{
		"bool": map[string]any{
			"must": []map[string]any{
				//{"match": map[string]any{"name": req.Key}},
				//{"match": map[string]any{"brand": req.Brand}},
			},
			"should": []map[string]any{},
			"filter": []map[string]any{
				//{"term": map[string]any{"category": map[string]any{"value": req.Category}}},
				//{"range": map[string]any{"price": map[string]any{"gte": req.MinPrice, "lte": req.MaxPrice}}},
			},
		},
	},
	"highlight": map[string]any{
		"fields": []map[string]any{
			//{"name": map[string]any{"pre_tags": "<em>", "post_tags": "</em>"}},
		},
	},
	"from": 0,  // 分⻚开始的位置，默认为0
	"size": 10, // 每⻚⽂档数量，默认10
}

func (l *SearchListLogic) SearchList(req *types.SearchListReq) (resp *types.SearchListResp, err error) {
	//确定分页大小
	pageSize := req.PageSize
	if req.PageSize == 0 {
		pageSize = types.EsDefaultPageSize
	}
	query["from"], query["size"] = req.PageNo, pageSize

	//品牌；如果有分类无关键字，聚合品牌 should，
	if req.Brand == "" {
		if req.Key == "" && req.Category != "" {
			buckets, err := l.BrandAggregation(req.Category)
			if err != nil {
				logx.Errorf("l.BrandAggregation, error: %v", err)
			}

			for _, bucket := range *buckets {
				bucketMap := bucket.(map[string]interface{})
				key := bucketMap["key"].(string)
				query["query"].(map[string]any)["bool"].(map[string]any)["should"] = append(
					query["query"].(map[string]any)["bool"].(map[string]any)["should"].([]map[string]any),
					map[string]any{"term": map[string]any{"Brand": map[string]any{"value": key}}})
			}
		}
	} else {
		query["query"].(map[string]any)["bool"].(map[string]any)["must"] = append(
			query["query"].(map[string]any)["bool"].(map[string]any)["must"].([]map[string]any),
			map[string]any{"term": map[string]any{"Brand": req.Brand}})
	}
	//关键字
	if req.Key != "" {
		query["query"].(map[string]any)["bool"].(map[string]any)["must"] = append(
			query["query"].(map[string]any)["bool"].(map[string]any)["must"].([]map[string]any),
			map[string]any{"match": map[string]any{"Name": req.Key}})

		query["highlight"].(map[string]any)["fields"] = append(
			query["highlight"].(map[string]any)["fields"].([]map[string]any),
			map[string]any{"Name": map[string]any{"pre_tags": "<em>", "post_tags": "</em>"}})
	}
	//分类
	if req.Category != "" {
		query["query"].(map[string]any)["bool"].(map[string]any)["filter"] = append(
			query["query"].(map[string]any)["bool"].(map[string]any)["filter"].([]map[string]any),
			map[string]any{"term": map[string]any{"Category": map[string]any{"value": req.Category}}})
	}
	//价格
	if req.MinPrice != 0 || req.MaxPrice != 0 {
		query["query"].(map[string]any)["bool"].(map[string]any)["filter"] = append(
			query["query"].(map[string]any)["bool"].(map[string]any)["filter"].([]map[string]any),
			map[string]any{"range": map[string]any{"Price": map[string]any{"gte": req.MinPrice, "lte": req.MaxPrice}}})
	}
	//序列化query
	marshal, err := json.Marshal(query)
	if err != nil {
		logx.Errorf("json.Marshal: %v, error: %v", query, err)
		return nil, err
	}
	//向es发起请求
	response, err := esapi.SearchRequest{
		Index: []string{types.EsItemsIndex},
		Body:  bytes.NewReader(marshal),
	}.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		logx.Errorf("esapi.SearchRequest: %v, error: %v", string(marshal), err)
		return nil, err
	}

	//处理数据
	start := strings.Index(response.String(), "]")
	respStr := response.String()[start+1:]

	var res Response
	err = json.Unmarshal([]byte(respStr), &res)
	if err != nil {
		logx.Errorf("json.Unmarshal: %v, error: %v", respStr, err)
		return nil, err
	}

	//获取用户id，用于创建用户搜索信息
	uid, err := util.GetUsr(l.ctx, types.JwtKey)
	if err != nil {
		logx.Errorf("util.GetUsr, error: %v", err)
		return nil, err
	}

	//构造返回对象
	//1、商品
	items := make([]types.SearchItemDTO, 0, len(res.Hits.HitDetails))
	for _, v := range res.Hits.HitDetails {
		if req.Key != "" {
			v.Source.Name = v.Highlight.Name[0] //处理高亮
		}
		items = append(items, v.Source)
		_ = l.AddUserSearchInfo(&UserSearchInfo{
			Uid:      strconv.Itoa(uid),
			Brand:    v.Source.Brand,
			ItemId:   strconv.FormatInt(v.Source.Id, 10),
			Category: v.Source.Category,
		})
	}

	return &types.SearchListResp{
		Items: items,
		Total: res.Hits.Total.Value,
		Pages: (req.PageNo / pageSize) + 1,
	}, nil
}

type Response struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Hits struct {
	Total      Total       `json:"total"`
	MaxScore   float64     `json:"max_score"`
	HitDetails []HitDetail `json:"hits"`
}

type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type HitDetail struct {
	Index     string              `json:"_index"`
	ID        string              `json:"_id"`
	Score     float64             `json:"_score"`
	Source    types.SearchItemDTO `json:"_source"`
	Highlight Highlight           `json:"highlight"`
}

type Highlight struct {
	Name []string `json:"name"`
}

type UserSearchInfo struct {
	Uid      string `json:"uid"`
	Brand    string `json:"brand"`
	Category string `json:"category"`
	ItemId   string `json:"item_id"`
}

func (l *SearchListLogic) BrandAggregation(category string) (*[]any, error) {
	BrandAggregation := `{
				"query": {
 					"bool": {
      					"filter": [
							{"term": {"category": ` + category + `}}
						]
    				}
				}, 
				"size": 0,
					"aggs": {
						"brand_agg": {
							"terms": {
								"field": "brand","size": 20
							}
						}
					}
				}`
	//向es发起请求
	resp, err := esapi.SearchRequest{
		Index: []string{types.EsItemsIndex},
		Body:  strings.NewReader(BrandAggregation),
	}.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		logx.Errorf("esapi.SearchRequest: %v, error: %v", query, err)
		return nil, err
	}
	//处理数据
	start := strings.Index(resp.String(), "]")
	data := resp.String()[start+1:]

	var result map[string]interface{}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	//取出buket
	aggregations := result["aggregations"].(map[string]interface{})
	categoryAgg := aggregations["category_agg"].(map[string]interface{})
	buckets := categoryAgg["buckets"].([]interface{})

	return &buckets, nil
}

func (l *SearchListLogic) AddUserSearchInfo(info *UserSearchInfo) error {
	marshal, _ := json.Marshal(info)
	resp, err := esapi.CreateRequest{
		Index:      types.EsUserSearchInfoIndex,
		Body:       bytes.NewReader(marshal),
		DocumentID: info.Uid + "_" + info.ItemId,
	}.Do(l.ctx, l.svcCtx.Es)
	if err != nil {
		logx.Errorf("esapi.CreateRequest: %v, error: %v", string(marshal), err)
		return err
	}
	log.Println("AddUserSearchInfo:", resp.String())
	return nil
}

// 数据库导入es
func (l *SearchListLogic) DatabasesToEs() error {
	item, err := l.svcCtx.ItemModel.FindItem(l.ctx, 1000, 0)
	if err != nil {
		logx.Errorf("ItemModel.FindItem, error:", err)
		return err
	}

	for _, v := range item {
		marshal, _ := json.Marshal(v)
		_, err := esapi.CreateRequest{
			Index:      types.EsItemsIndex,
			DocumentID: strconv.FormatInt(v.Id, 10),
			Body:       bytes.NewReader(marshal),
		}.Do(context.Background(), l.svcCtx.Es)
		if err != nil {
			logx.Errorf("esapi.CreateRequest: %v, error: %v", string(marshal), err)
			return err
		}
	}
	return nil
}
