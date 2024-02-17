---
title: hmall——go-zero
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

# hmall——go-zero

Base URLs:

# Authentication

# 购物车相关接口

<a id="opIdqueryMyCartsUsingGET"></a>

## GET 查询购物车列表

GET /carts

> Response Examples

> 200 Response

```json
[
  {
    "createTime": "2019-08-24T14:15:22Z",
    "id": 0,
    "image": "string",
    "itemId": 0,
    "name": "string",
    "newPrice": 0,
    "num": 0,
    "price": 0,
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
|*anonymous*|[[CartVO](#schemacartvo)]|false|none||[购物车VO实体]|
|» CartVO|[CartVO](#schemacartvo)|false|none|CartVO|购物车VO实体|
|»» createTime|string(date-time)|false|none||创建时间|
|»» id|integer(int64)|false|none||购物车条目id|
|»» image|string|false|none||商品图片|
|»» itemId|integer(int64)|false|none||sku商品id|
|»» name|string|false|none||商品标题|
|»» newPrice|integer(int32)|false|none||商品最新价格|
|»» num|integer(int32)|false|none||购买数量|
|»» price|integer(int32)|false|none||价格,单位：分|
|»» spec|string|false|none||商品动态属性键值集|
|»» status|integer(int32)|false|none||商品最新状态|
|»» stock|integer(int32)|false|none||商品最新库存|

<a id="opIdaddItem2CartUsingPOST"></a>

## POST 添加商品到购物车

POST /carts

> Body Parameters

```json
{
  "image": "string",
  "itemId": 0,
  "name": "string",
  "price": 0,
  "spec": "string"
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|body|body|[CartFormDTO](#schemacartformdto)| no | CartFormDTO|none|

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

<a id="opIdupdateCartUsingPUT"></a>

## PUT 更新购物车数据

PUT /carts

> Body Parameters

```json
{
  "createTime": "2019-08-24T14:15:22Z",
  "id": 0,
  "image": "string",
  "itemId": 0,
  "name": "string",
  "num": 0,
  "price": 0,
  "spec": "string",
  "updateTime": "2019-08-24T14:15:22Z",
  "userId": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|body|body|[Cart](#schemacart)| no | Cart|none|

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

<a id="opIddeleteCartItemByIdsUsingDELETE"></a>

## DELETE 批量删除购物车中商品

DELETE /carts

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|ids|query|string| no ||购物车条目id集合|

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

<a id="opIddeleteCartItemUsingDELETE"></a>

## DELETE 删除购物车中商品

DELETE /carts/{id}

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

<h2 id="tocS_Cart">Cart</h2>

<a id="schemacart"></a>
<a id="schema_Cart"></a>
<a id="tocScart"></a>
<a id="tocscart"></a>

```json
{
  "createTime": "2019-08-24T14:15:22Z",
  "id": 0,
  "image": "string",
  "itemId": 0,
  "name": "string",
  "num": 0,
  "price": 0,
  "spec": "string",
  "updateTime": "2019-08-24T14:15:22Z",
  "userId": 0
}

```

Cart

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|createTime|string(date-time)|false|none||none|
|id|integer(int64)|false|none||none|
|image|string|false|none||none|
|itemId|integer(int64)|false|none||none|
|name|string|false|none||none|
|num|integer(int32)|false|none||none|
|price|integer(int32)|false|none||none|
|spec|string|false|none||none|
|updateTime|string(date-time)|false|none||none|
|userId|integer(int64)|false|none||none|

<h2 id="tocS_CartFormDTO">CartFormDTO</h2>

<a id="schemacartformdto"></a>
<a id="schema_CartFormDTO"></a>
<a id="tocScartformdto"></a>
<a id="tocscartformdto"></a>

```json
{
  "image": "string",
  "itemId": 0,
  "name": "string",
  "price": 0,
  "spec": "string"
}

```

CartFormDTO

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|image|string|false|none||商品图片|
|itemId|integer(int64)|false|none||商品id|
|name|string|false|none||商品标题|
|price|integer(int32)|false|none||价格,单位：分|
|spec|string|false|none||商品动态属性键值集|

<h2 id="tocS_CartVO">CartVO</h2>

<a id="schemacartvo"></a>
<a id="schema_CartVO"></a>
<a id="tocScartvo"></a>
<a id="tocscartvo"></a>

```json
{
  "createTime": "2019-08-24T14:15:22Z",
  "id": 0,
  "image": "string",
  "itemId": 0,
  "name": "string",
  "newPrice": 0,
  "num": 0,
  "price": 0,
  "spec": "string",
  "status": 0,
  "stock": 0
}

```

CartVO

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|createTime|string(date-time)|false|none||创建时间|
|id|integer(int64)|false|none||购物车条目id|
|image|string|false|none||商品图片|
|itemId|integer(int64)|false|none||sku商品id|
|name|string|false|none||商品标题|
|newPrice|integer(int32)|false|none||商品最新价格|
|num|integer(int32)|false|none||购买数量|
|price|integer(int32)|false|none||价格,单位：分|
|spec|string|false|none||商品动态属性键值集|
|status|integer(int32)|false|none||商品最新状态|
|stock|integer(int32)|false|none||商品最新库存|

