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

# 用户相关接口

<a id="opIdloginUsingPOST"></a>

## POST 用户登录接口

POST /users/login

> Body 请求参数

```json
{
  "password": "string",
  "rememberMe": true,
  "username": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|[LoginFormDTO](#schemaloginformdto)| 否 | LoginFormDTO|none|

> 返回示例

> 200 Response

```json
{
  "balance": 0,
  "token": "string",
  "userId": 0,
  "username": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[UserLoginVO](#schemauserloginvo)|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|Forbidden|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|Inline|

### 返回数据结构

<a id="opIddeductMoneyUsingPUT"></a>

## PUT 扣减余额

PUT /users/money/deduct

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|amount|query|string| 否 ||支付金额|
|pw|query|string| 否 ||支付密码|

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

<h2 id="tocS_UserLoginVO">UserLoginVO</h2>

<a id="schemauserloginvo"></a>
<a id="schema_UserLoginVO"></a>
<a id="tocSuserloginvo"></a>
<a id="tocsuserloginvo"></a>

```json
{
  "balance": 0,
  "token": "string",
  "userId": 0,
  "username": "string"
}

```

UserLoginVO

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|balance|integer(int32)|false|none||none|
|token|string|false|none||none|
|userId|integer(int64)|false|none||none|
|username|string|false|none||none|

<h2 id="tocS_LoginFormDTO">LoginFormDTO</h2>

<a id="schemaloginformdto"></a>
<a id="schema_LoginFormDTO"></a>
<a id="tocSloginformdto"></a>
<a id="tocsloginformdto"></a>

```json
{
  "password": "string",
  "rememberMe": true,
  "username": "string"
}

```

LoginFormDTO

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|password|string|true|none||用户名|
|rememberMe|boolean|false|none||是否记住我|
|username|string|true|none||用户名|

