package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBrand(t *testing.T) {
	data := `{
	  "took": 233,
	  "timed_out": false,
	  "_shards": {
	    "total": 1,
	    "successful": 1,
	    "skipped": 0,
	    "failed": 0
	  },
	  "hits": {
	    "total": {
	      "value": 1002,
	      "relation": "eq"
	    },
	    "max_score": null,
	    "hits": []
	  },
	  "aggregations": {
	    "category_agg": {
	      "doc_count_error_upper_bound": 0,
	      "sum_other_doc_count": 0,
	      "buckets": [
	        {
	          "key": "拉杆箱",
	          "doc_count": 714
	        },
	        {
	          "key": "牛奶",
	          "doc_count": 108
	        },
	        {
	          "key": "真皮包",
	          "doc_count": 60
	        },
	        {
	          "key": "拉拉裤",
	          "doc_count": 43
	        },
	        {
	          "key": "手机",
	          "doc_count": 35
	        },
	        {
	          "key": "老花镜",
	          "doc_count": 22
	        },
	        {
	          "key": "硬盘",
	          "doc_count": 14
	        },
	        {
	          "key": "曲面电视",
	          "doc_count": 3
	        },
	        {
	          "key": "non id nisi sit pariatur",
	          "doc_count": 1
	        },
	        {
	          "key": "reprehenderit",
	          "doc_count": 1
	        },
	        {
	          "key": "reprehenderit et",
	          "doc_count": 1
	        }
	      ]
	    }
	  }
	}`

	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	aggregations := result["aggregations"].(map[string]interface{})
	categoryAgg := aggregations["category_agg"].(map[string]interface{})
	buckets := categoryAgg["buckets"].([]interface{})

	for _, bucket := range buckets {
		bucketMap := bucket.(map[string]interface{})
		key := bucketMap["key"].(string)
		fmt.Println("Key:", key)
	}
}
