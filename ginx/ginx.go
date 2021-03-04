package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result result
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
	//
	c          *gin.Context
	statusCode int
}

// NewResult .
func NewResult(c *gin.Context) *Result {
	r := &Result{
		c:          c,
		statusCode: http.StatusOK,
	}
	return r
}

// Response .
func (r *Result) Response() {
	r.c.JSON(r.statusCode, r)
}

// ContentType .
// gin/binding.MIMEJSON
func (r *Result) ContentType(value string) *Result {
	key := "Content-Type"
	r.c.Writer.Header().Del(key)
	r.c.Header(key, value)
	return r
}

// WriteString .
func (r *Result) WriteString(data string) (int, error) {
	return r.c.Writer.WriteString(data)
}

// Set set
func (r *Result) Set(code int, msg string, data interface{}) *Result {
	r.Code = code
	r.Msg = msg
	r.Data = data
	return r
}

// Set200 200
func (r *Result) Set200(msg string, data interface{}) *Result {
	r.statusCode = http.StatusOK
	r.Set(200, msg, data)
	return r
}

// Status http status
func (r *Result) Status(code int) *Result {
	r.statusCode = code
	return r
}

// Data .
func (r *Result) SetData(data interface{}) *Result {
	r.Data = data
	return r
}

// Set400 400
func (r *Result) Set400(msg string) *Result {
	r.statusCode = http.StatusBadRequest
	r.Set(400, msg, nil)
	return r
}

// Set500 500
func (r *Result) Set500(msg string) *Result {
	r.statusCode = http.StatusInternalServerError
	r.Set(500, msg, nil)
	return r
}

// Response.
func Response(c *gin.Context, status int, code int, msg string, data interface{}) {
	c.JSON(status, Result{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	return
}

// Ok .
func Ok(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: 200,
		Msg:  msg,
	})
	return
}

// Faild .
func Faild(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusBadRequest, Result{
		Code: code,
		Msg:  msg,
	})
	return
}

// FaildWithStatus .
func FaildWithStatus(c *gin.Context, status int, code int, msg string) {
	c.JSON(status, Result{
		Code: code,
		Msg:  msg,
	})
	return
}
