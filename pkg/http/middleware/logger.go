package middleware

import (
	"bytes"
	"cloud_callback/internal/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"time"
)

const DateFormat = "2006-01-02 15:04:05"

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	b2 := b
	w.body.Write(b2)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LoggerMiddleware() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		writer := CustomResponseWriter{ResponseWriter: c.Writer, body: bytes.NewBufferString("")}
		c.Writer = writer
		start := time.Now()
		url := c.Request.URL.Path
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Error("LoggerMiddleware-->ReadAll", zap.Error(err))
			c.Abort()
		}
		// 再次放入數據
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		log.Info(fmt.Sprintf("%s-->Request", url),
			zap.String("startTime", start.Format(DateFormat)),
			zap.String("method", c.Request.Method),
			zap.String("body", string(body)))
		c.Next()
		end := time.Now()
		costTime := end.Sub(start).Seconds()
		log.Info(fmt.Sprintf("%s-->Response", url),
			zap.String("endTime", end.Format(DateFormat)),
			zap.Float64("costTime", costTime),
			zap.String("respBody", writer.body.String()))
	}
}
