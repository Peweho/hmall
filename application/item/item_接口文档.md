---
title: 黑马商城——go-zero
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.20"

---

# 黑马商城——go-zero

Base URLs:

# Authentication

# 商品管理相关接口

<a id="opIdqueryItemByIdsUsingGET"></a>

## GET 根据id批量查询商品

GET /items

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|ids|query|array[string]| yes |ids|

> Response Examples

> 200 Response

```json
[
  {
    "brand": "string",
    "category": "string",
    "commentCount": 0,
    "id": 0,
    "image": "string",
    "isAD": true,
    "name": "string",
    "price": 0,
    "sold": 0,
    "spec": "string",
    "status": 0,
    "stock": 0
  }
]
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|*anonymous*|[[ItemDTO](#schemaitemdto)]|false|none||[商品实体]|
|» ItemDTO|[ItemDTO](#schemaitemdto)|false|none|ItemDTO|商品实体|
|»» brand|string|false|none||品牌名称|
|»» category|string|false|none||类目名称|
|»» commentCount|integer(int32)|false|none||评论数|
|»» id|integer(int64)|false|none||商品id|
|»» image|string|false|none||商品图片|
|»» isAD|boolean|false|none||是否是推广广告，true/false|
|»» name|string|false|none||SKU名称|
|»» price|integer(int32)|false|none||价格（分）|
|»» sold|integer(int32)|false|none||销量|
|»» spec|string|false|none||规格|
|»» status|integer(int32)|false|none||商品状态 1-正常，2-下架，3-删除|
|»» stock|integer(int32)|false|none||库存数量|

<a id="opIdsaveItemUsingPOST"></a>

## POST 新增商品

POST /items

> Body Parameters

```json
{
  "brand": "string",
  "category": "string",
  "commentCount": 0,
  "id": 0,
  "image": "string",
  "isAD": true,
  "name": "string",
  "price": 0,
  "sold": 0,
  "spec": "string",
  "status": 0,
  "stock": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|body|body|[ItemDTO](#schemaitemdto)| no | ItemDTO|none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### Responses Data Schema

<a id="opIdupdateItemUsingPUT"></a>

## PUT 更新商品

PUT /items

> Body Parameters

```json
{
  "brand": "string",
  "category": "string",
  "commentCount": 0,
  "id": 0,
  "image": "string",
  "isAD": true,
  "name": "string",
  "price": 0,
  "sold": 0,
  "spec": "string",
  "status": 0,
  "stock": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|body|body|[ItemDTO](#schemaitemdto)| no | ItemDTO|none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### Responses Data Schema

<a id="opIdqueryItemByPageUsingGET"></a>

## GET 分页查询商品

GET /items/page

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|isAsc|query|string| no ||是否升序|
|pageNo|query|integer| no ||页码|
|pageSize|query|integer| no ||页码|
|sortBy|query|string| no ||排序方式|

> Response Examples

> 200 Response

```json
{
  "list": [
    {
      "brand": "string",
      "category": "string",
      "commentCount": 0,
      "id": 0,
      "image": "string",
      "isAD": true,
      "name": "string",
      "price": 0,
      "sold": 0,
      "spec": "string",
      "status": 0,
      "stock": 0
    }
  ],
  "pages": 0,
  "total": 0
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### Responses Data Schema

<a id="opIdupdateItemStatusUsingPUT"></a>

## PUT 更新商品状态

PUT /items/status/{id}/{status}

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|id|path|integer| yes ||id|
|status|path|integer| yes ||status|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### Responses Data Schema

<a id="opIddeductStockUsingPUT"></a>

## PUT 批量扣减库存

PUT /items/stock/deduct

> Body Parameters

```json
[
  {
    "itemId": 0,
    "num": 0
  }
]
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|body|body|[OrderDetailDTO](#schemaorderdetaildto)| no ||none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### Responses Data Schema

<a id="opIdqueryItemByIdUsingGET"></a>

## GET 根据id查询商品

GET /items/{id}

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|id|path|integer| yes ||id|

> Response Examples

> 200 Response

```json
{
  "brand": "string",
  "category": "string",
  "commentCount": 0,
  "id": 0,
  "image": "string",
  "isAD": true,
  "name": "string",
  "price": 0,
  "sold": 0,
  "spec": "string",
  "status": 0,
  "stock": 0
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[ItemDTO](#schemaitemdto)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### Responses Data Schema

<a id="opIddeleteItemByIdUsingDELETE"></a>

## DELETE 根据id删除商品

DELETE /items/{id}

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|id|path|integer| yes ||id|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|No Content|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|

### Responses Data Schema

# Data Schema

<h2 id="tocS_OrderDetailDTO">OrderDetailDTO</h2>

<a id="schemaorderdetaildto"></a>
<a id="schema_OrderDetailDTO"></a>
<a id="tocSorderdetaildto"></a>
<a id="tocsorderdetaildto"></a>

```json
{
  "itemId": 0,
  "num": 0
}

```

OrderDetailDTO

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|itemId|integer(int64)|false|none||商品id|
|num|integer(int32)|false|none||商品购买数量|

<h2 id="tocS_PageDTO«ItemDTO»">PageDTO«ItemDTO»</h2>

<a id="schemapagedto«itemdto»"></a>
<a id="schema_PageDTO«ItemDTO»"></a>
<a id="tocSpagedto«itemdto»"></a>
<a id="tocspagedto«itemdto»"></a>

```json
{
  "list": [
    {
      "brand": "string",
      "category": "string",
      "commentCount": 0,
      "id": 0,
      "image": "string",
      "isAD": true,
      "name": "string",
      "price": 0,
      "sold": 0,
      "spec": "string",
      "status": 0,
      "stock": 0
    }
  ],
  "pages": 0,
  "total": 0
}

```

PageDTO«ItemDTO»

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|list|[[ItemDTO](#schemaitemdto)]|false|none||[商品实体]|
|pages|integer(int64)|false|none||none|
|total|integer(int64)|false|none||none|

<h2 id="tocS_ItemDTO">ItemDTO</h2>

<a id="schemaitemdto"></a>
<a id="schema_ItemDTO"></a>
<a id="tocSitemdto"></a>
<a id="tocsitemdto"></a>

```json
{
  "brand": "string",
  "category": "string",
  "commentCount": 0,
  "id": 0,
  "image": "string",
  "isAD": true,
  "name": "string",
  "price": 0,
  "sold": 0,
  "spec": "string",
  "status": 0,
  "stock": 0
}

```

ItemDTO

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|brand|string|false|none||品牌名称|
|category|string|false|none||类目名称|
|commentCount|integer(int32)|false|none||评论数|
|id|integer(int64)|false|none||商品id|
|image|string|false|none||商品图片|
|isAD|boolean|false|none||是否是推广广告，true/false|
|name|string|false|none||SKU名称|
|price|integer(int32)|false|none||价格（分）|
|sold|integer(int32)|false|none||销量|
|spec|string|false|none||规格|
|status|integer(int32)|false|none||商品状态 1-正常，2-下架，3-删除|
|stock|integer(int32)|false|none||库存数量|

