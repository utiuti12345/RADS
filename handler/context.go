package handler

type Context interface {
	String(code int, s string) error
	JSON(statusCode int,i interface{}) error
	Param(p string) string
	Bind(i interface{}) error
}