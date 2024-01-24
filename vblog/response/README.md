# 数据响应格式统一

## 为什么要统一？

```json
{
    code: 5000
    reason: "token expired"
    message: "过期10分钟"
}
```

```json
{
    code: 0
    data: []/{}
}
```

通过外面多包一层code字段来判断 来判断这次API请求是否成功

```json
{
    "code": 0
    "message": "失败原因"
    data: {}/[]
}
```

## 怎么设计统一返回

HTTP 协议，HTTP Status Code来判断一个请求是否成功 4xx，5xx

借助于Http Status Code，来表达 API请求是否成功

成功返回：直接返回数据内容，对象本身
```json
{}/[]
```

失败返回：直接业务异常，APIException对象，如果err 不是APIException，就需要进行转换
```json
"code": 0
"reason": "expired"
"message": "过期了10分钟"
```