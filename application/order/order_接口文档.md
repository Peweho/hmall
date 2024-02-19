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

# 订单管理接口

<a id="opIdcreateOrderUsingPOST"></a>

## POST 创建订单

POST /orders

> Body Parameters

```json
{
  "addressId": 0,
  "details": [
    {
      "itemId": 0,
      "num": 0
    }
  ],
  "paymentType": 0
}
```

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|body|body|[OrderFormDTO](#schemaorderformdto)| no | OrderFormDTO|none|

> Response Examples

> 200 Response

```json
0
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|integer|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### Responses Data Schema

<a id="opIdqueryOrderByIdUsingGET"></a>

## GET 根据id查询订单

GET /orders/{id}

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|id|path|integer| yes ||id|

> Response Examples

> 200 Response

```json
{
  "closeTime": "2019-08-24T14:15:22Z",
  "commentTime": "2019-08-24T14:15:22Z",
  "consignTime": "2019-08-24T14:15:22Z",
  "createTime": "2019-08-24T14:15:22Z",
  "endTime": "2019-08-24T14:15:22Z",
  "id": 0,
  "payTime": "2019-08-24T14:15:22Z",
  "paymentType": 0,
  "status": 0,
  "totalFee": 0,
  "userId": 0
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[OrderVO](#schemaordervo)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### Responses Data Schema

<a id="opIdmarkOrderPaySuccessUsingPUT"></a>

## PUT 标记订单已支付

PUT /orders/{orderId}

### Params

|Name|Location|Type|Required|Title|Description|
|---|---|---|---|---|---|
|orderId|path|string| yes ||订单id|

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

# Data Schema

<h2 id="tocS_OrderVO">OrderVO</h2>

<a id="schemaordervo"></a>
<a id="schema_OrderVO"></a>
<a id="tocSordervo"></a>
<a id="tocsordervo"></a>

```json
{
  "closeTime": "2019-08-24T14:15:22Z",
  "commentTime": "2019-08-24T14:15:22Z",
  "consignTime": "2019-08-24T14:15:22Z",
  "createTime": "2019-08-24T14:15:22Z",
  "endTime": "2019-08-24T14:15:22Z",
  "id": 0,
  "payTime": "2019-08-24T14:15:22Z",
  "paymentType": 0,
  "status": 0,
  "totalFee": 0,
  "userId": 0
}

```

OrderVO

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|closeTime|string(date-time)|false|none||交易关闭时间|
|commentTime|string(date-time)|false|none||评价时间|
|consignTime|string(date-time)|false|none||发货时间|
|createTime|string(date-time)|false|none||创建时间|
|endTime|string(date-time)|false|none||交易完成时间|
|id|integer(int64)|false|none||订单id|
|payTime|string(date-time)|false|none||支付时间|
|paymentType|integer(int32)|false|none||支付类型，1、支付宝，2、微信，3、扣减余额|
|status|integer(int32)|false|none||订单的状态，1、未付款 2、已付款,未发货 3、已发货,未确认 4、确认收货，交易成功 5、交易取消，订单关闭 6、交易结束，已评价|
|totalFee|integer(int32)|false|none||总金额，单位为分|
|userId|integer(int64)|false|none||用户id|

<h2 id="tocS_OrderFormDTO">OrderFormDTO</h2>

<a id="schemaorderformdto"></a>
<a id="schema_OrderFormDTO"></a>
<a id="tocSorderformdto"></a>
<a id="tocsorderformdto"></a>

```json
{
  "addressId": 0,
  "details": [
    {
      "itemId": 0,
      "num": 0
    }
  ],
  "paymentType": 0
}

```

OrderFormDTO

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|addressId|integer(int64)|false|none||收货地址id|
|details|[[OrderDetailDTO](#schemaorderdetaildto)]|false|none||下单商品列表|
|paymentType|integer(int32)|false|none||支付类型|

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

