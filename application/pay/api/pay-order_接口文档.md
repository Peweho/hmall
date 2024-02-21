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

# 支付相关接口

<a id="opIdapplyPayOrderUsingPOST"></a>

## POST 生成支付单

POST /pay-orders

> Body 请求参数

```json
{
  "amount": 0,
  "bizOrderNo": 0,
  "orderInfo": "string",
  "payChannelCode": "string",
  "payType": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|[PayApplyDTO](#schemapayapplydto)| 否 | PayApplyDTO|none|

> 返回示例

> 200 Response

```json
"string"
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### 返回数据结构

<a id="opIdtryPayOrderByBalanceUsingPOST"></a>

## POST 尝试基于用户余额支付

POST /pay-orders/{id}

> Body 请求参数

```json
{
  "id": 0,
  "pw": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|path|string| 是 ||支付单id|
|body|body|[PayOrderFormDTO](#schemapayorderformdto)| 否 | PayOrderFormDTO|none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### 返回数据结构

# 数据模型

<h2 id="tocS_PayOrderFormDTO">PayOrderFormDTO</h2>

<a id="schemapayorderformdto"></a>
<a id="schema_PayOrderFormDTO"></a>
<a id="tocSpayorderformdto"></a>
<a id="tocspayorderformdto"></a>

```json
{
  "id": 0,
  "pw": "string"
}

```

PayOrderFormDTO

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|false|none||支付订单id不能为空|
|pw|string|false|none||支付密码|

<h2 id="tocS_PayApplyDTO">PayApplyDTO</h2>

<a id="schemapayapplydto"></a>
<a id="schema_PayApplyDTO"></a>
<a id="tocSpayapplydto"></a>
<a id="tocspayapplydto"></a>

```json
{
  "amount": 0,
  "bizOrderNo": 0,
  "orderInfo": "string",
  "payChannelCode": "string",
  "payType": 0
}

```

PayApplyDTO

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|amount|integer(int32)|false|none||支付金额必须为正数|
|bizOrderNo|integer(int64)|false|none||业务订单id不能为空|
|orderInfo|string|false|none||订单中的商品信息不能为空|
|payChannelCode|string|false|none||支付渠道编码不能为空|
|payType|integer(int32)|false|none||支付方式不能为空|

