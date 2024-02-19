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

# 收货地址管理接口

<a id="opIdfindMyAddressesUsingGET"></a>

## GET 查询当前用户地址列表

GET /addresses

> Response Examples

> 200 Response

```json
[
  {
    "city": "string",
    "contact": "string",
    "id": 0,
    "isDefault": 0,
    "mobile": "string",
    "notes": "string",
    "province": "string",
    "street": "string",
    "town": "string"
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
|*anonymous*|[[AddressDTO](#schemaaddressdto)]|false|none||[收货地址实体]|
|» AddressDTO|[AddressDTO](#schemaaddressdto)|false|none|AddressDTO|收货地址实体|
|»» city|string|false|none||市|
|»» contact|string|false|none||联系人|
|»» id|integer(int64)|false|none||id|
|»» isDefault|integer(int32)|false|none||是否是默认 1默认 0否|
|»» mobile|string|false|none||手机|
|»» notes|string|false|none||备注|
|»» province|string|false|none||省|
|»» street|string|false|none||详细地址|
|»» town|string|false|none||县/区|

<a id="opIdfindAddressByIdUsingGET"></a>

## GET 根据id查询地址

GET /addresses/{addressId}

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|addressId|path|integer| yes |地址id|

> Response Examples

> 200 Response

```json
{
  "city": "string",
  "contact": "string",
  "id": 0,
  "isDefault": 0,
  "mobile": "string",
  "notes": "string",
  "province": "string",
  "street": "string",
  "town": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[AddressDTO](#schemaaddressdto)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### Responses Data Schema

# Data Schema

<h2 id="tocS_AddressDTO">AddressDTO</h2>

<a id="schemaaddressdto"></a>
<a id="schema_AddressDTO"></a>
<a id="tocSaddressdto"></a>
<a id="tocsaddressdto"></a>

```json
{
  "city": "string",
  "contact": "string",
  "id": 0,
  "isDefault": 0,
  "mobile": "string",
  "notes": "string",
  "province": "string",
  "street": "string",
  "town": "string"
}

```

AddressDTO

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|city|string|false|none||市|
|contact|string|false|none||联系人|
|id|integer(int64)|false|none||id|
|isDefault|integer(int32)|false|none||是否是默认 1默认 0否|
|mobile|string|false|none||手机|
|notes|string|false|none||备注|
|province|string|false|none||省|
|street|string|false|none||详细地址|
|town|string|false|none||县/区|

