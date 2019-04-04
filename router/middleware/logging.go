package middleware

import (
	"github.com/gin-gonic/gin"
	"bytes"
	"time"
	"regexp"
	"io/ioutil"
	"github.com/lexkong/log"
	"apiserver/handler"
	"encoding/json"
	"apiserver/pkg/errno"
	"github.com/willf/pad"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func(w bodyLogWriter) Write(b []byte)(int,error){
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc{
	return func(c *gin.Context){
		start := time.Now().UTC()
		path := c.Request.URL.Path

		req := regexp.MustCompile("/v1/user|/login")
		if !req.MatchString(path){
			return
		}

		// Skip for the health check requests.
		if path == "/middware/health" || path == "/middware/ram" || path == "/middware/cpu" || path == "/middware/disk" {
			return
		}
		var bodyBytes []byte
		if c.Request.Body != nil{
			bodyBytes,_ =ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		method := c.Request.Method
		ip := c.ClientIP()

		blw := &bodyLogWriter{
			body: bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		var response handler.Response

		if err := json.Unmarshal(blw.body.Bytes(),&response);err!= nil{
			log.Errorf(err, "response body can not unmarshal to model.Response struct, body: `%s`", blw.body.Bytes())
			code = errno.InternalServerError.Code
			message = err.Error()
		}else{
			code = response.Code
			message = response.Message
		}

		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)
	}
}