package reqctx

type Context interface {
	JSON(code int, i interface{}) error
	QueryParam(name string) string
}
