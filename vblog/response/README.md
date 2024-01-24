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