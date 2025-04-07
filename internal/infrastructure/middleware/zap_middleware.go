package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GinzapExtraFields(c *gin.Context) []zapcore.Field {
	fields := []zapcore.Field{}

	if requestID := c.Request.Header.Get("X-Request-Id"); requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
	}

	if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		bodyBytes, err := io.ReadAll(tee)
		if err == nil && len(bodyBytes) > 0 {
			var parsed interface{}
			if json.Unmarshal(bodyBytes, &parsed) == nil {
				fields = append(fields, zap.Any("body", parsed))
			} else {
				fields = append(fields, zap.String("body", string(bodyBytes)))
			}
			c.Request.Body = io.NopCloser(&buf)
		} else if err != nil {
			fields = append(fields, zap.String("body", "<failed to read>"))
		}
	}

	return fields
}
