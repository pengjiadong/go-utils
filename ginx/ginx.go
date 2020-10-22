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
