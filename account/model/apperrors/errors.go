package apperrors

import (
	"errors"
	"fmt"
	"net/http"
)

type Type string

const (
	Authorization   Type = "AUTHORIZATION"   // 授权类错误：Unauthorized (401)
	BadRequest      Type = "BADREQUEST"      // 请求类错误：Bad request (400)
	Conflict        Type = "CONFLICT"        // 冲突类错误：Already exists (409)
	Internal        Type = "INTERNAL"        // 内部类错误：Server (500) 和 fallback errors
	NotFound        Type = "NOTFOUND"        // 未找到类错误：404
	PayloadTooLarge Type = "PAYLOADTOOLARGE" // 负载过大类错误：413
)

// Error defines the structure for an API error
// Error 定义了 API 错误的结构
type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

// Error satisfies the standard error interface
// we can return errors from this package as a regular old go _error_
// Error 满足标准的 error 接口
// 我们可以将这个包中的错误作为一个普通的 go _error_ 返回
func (e *Error) Error() string {
	return e.Message
}

// Status is the mapping errors to status codes
// of course, this is somewhat redundant since our errors already map http status codes
// Status 是将错误映射到状态码, 当然，这有点多余，因为我们的错误已经映射了 http 状态码
func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	default:
		return http.StatusInternalServerError
	}
}

// Status checks the runtime error
// and returns an http status code
// if the error is model.Error
// Status 检查运行时错误并返回 http 状态码, 如果错误是 model.Error, 则返回错误的状态码, 否则返回 500
func Status(err error) int {
	// if no error, return 200 OK
	if err == nil {
		return http.StatusOK
	}
	// if our custom error, return the status code
	var e *Error
	ok := errors.As(err, &e)
	// if not our custom error, return 500
	if !ok {
		return http.StatusInternalServerError
	}
	return e.Status()
}

// NewAuthorization creates a 401 error
func NewAuthorization(message string) *Error {
	return &Error{
		Type:    Authorization,
		Message: message,
	}
}

// NewBadRequest creates a 400 errors( validation, malformed request body, etc)
func NewBadRequest(message string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request: %v", message),
	}
}

// NewConflict creates a 409 error
func NewConflict(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v already exists: %v", name, value),
	}
}

// NewInternal creates a 500 error
func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error"),
	}
}

// NewNotFound creates a 404 error
func NewNotFound(name string, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("resource: %v with the value %v not found", name, value),
	}
}

// NewPayloadTooLarge creates a 413 error
func NewPayloadTooLarge(maxBodySize int64, contentLength int64) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("payload too large: max size is %v, but got %v", maxBodySize, contentLength),
	}
}
