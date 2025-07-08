package utils

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// Logger logging utility with trace id
type Logger struct {
	traceId string
}

// NewLogger creates a new logger instance
func NewLogger(c *gin.Context) *Logger {
	traceId, exists := c.Get("traceId")
	if !exists {
		return &Logger{traceId: "unknown"}
	}
	return &Logger{traceId: traceId.(string)}
}

// NewLoggerFromContext creates logger instance from context
func NewLoggerFromContext(ctx context.Context) *Logger {
	traceId := ctx.Value("traceId")
	if traceId == nil {
		return &Logger{traceId: "unknown"}
	}
	return &Logger{traceId: traceId.(string)}
}

// Info logs info level message
func (l *Logger) Info(msg string, args ...interface{}) {
	logMsg := fmt.Sprintf("[INFO] [TraceId: %s] %s", l.traceId, fmt.Sprintf(msg, args...))
	log.Println(logMsg)
}

// Error logs error level message
func (l *Logger) Error(msg string, args ...interface{}) {
	logMsg := fmt.Sprintf("[ERROR] [TraceId: %s] %s", l.traceId, fmt.Sprintf(msg, args...))
	log.Println(logMsg)
}

// Debug logs debug level message
func (l *Logger) Debug(msg string, args ...interface{}) {
	logMsg := fmt.Sprintf("[DEBUG] [TraceId: %s] %s", l.traceId, fmt.Sprintf(msg, args...))
	log.Println(logMsg)
}

// Warn logs warn level message
func (l *Logger) Warn(msg string, args ...interface{}) {
	logMsg := fmt.Sprintf("[WARN] [TraceId: %s] %s", l.traceId, fmt.Sprintf(msg, args...))
	log.Println(logMsg)
}
