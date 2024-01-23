# 业务异常

## 为什么需要业务异常

判断 令牌过期

业务异常的使用方式：
```go
err := Biz_Call()
if err == TokenExpired {

}
```

字符串比对，可能造成误杀
```go
access token expired %f minutes

hasPrefix("access token expired")
```

设计一套业务专用的业务异常，通常设计为异常码（Error Code）:
```go
// err.ErrorCode == xxxx 
if exception.IsTokenExpired(err) {

}
```

## 怎么设计业务异常

想方设法扩展error

本身需要兼容Error的场景：
```go
func XXX() error 
```

go 里面的Error是个接口
```go
type error interface {
    Error() string
}
```

fmt包里提供的Error实现
```go

```

如何定义自定义异常
```go
func NewAPIException(code int, msg string) *APIExcetion {
	return &APIExcetion{
		code: code,
		msg: msg,
	}
}

// error的自定义实现
type APIExcetion struct {
	code int
	msg string
}

func(e *APIExcetion) Error() string {
	return e.msg
}

func(e *APIExcetion) Code() int {
	return e.code
}
```

## 定义业务异常

1. 定义 TokenExpired 5000

```go
// 这个模块定义的业务异常
// token expired %f minutes
// 约定俗成： ErrXXXXXX 来定义自定义异常当以，方便快速在包内搜索
var (
	ErrAccessTokenExpired = exception.NewAPIException(5000, "access token expired")
	ErrRefreshTokenExpired = exception.NewAPIException(5001, "refresh token expired")
)
```

2. 使用自定义异常
```go
if  aDelta > 0 {
		return ErrAccessTokenExpired.WithMessagef("access token expired %f minutes", aDelta)
		// return fmt.Errorf("access token expired %f minutes", aDelta)
	}
```

3. 如何判断异常是否相等

1. 基于断言后根据Code来进行业务异常判断
```go
	// 通过断言来获取一个exception
	if e, ok := err.(*exception.APIExcetion); ok {
		t.Log(e.String())
		// 判断该异常是不是 TokenExpired 异常
		if e.Code == token.ErrAccessTokenExpired.Code {
			t.Log(e.String())
		}
	}
```

```go
// exception.IsException(err, token.ErrAccessTokenExpired)
// 给一个异常判断的方法
func IsException(err error, e *APIExcetion) bool {
	if target, ok := err.(*APIExcetion); ok {
		return target.Code == e.Code
	}
	return false
}
```