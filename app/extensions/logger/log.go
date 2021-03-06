package logger

import (
	"context"
	"runtime/debug"
	"strings"

	"github.com/zhenligod/go-web/app/helper"

	"github.com/zhenligod/thingo/gutils"
	"github.com/zhenligod/thingo/logger"
)

/**
{
    "level":"info",
    "time_local":"2019-11-24T19:36:02.838+0800",
    "msg":"exec start",
    "request_method":"GET",
    "request_uri":"/v1/hello",
    "log_id":"c34a00fa-0906-2fff-b5e8-61cfcc2a19ff",
    "ua":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.80 Safari/537.36",
    "plat":"web",
    "tag":"v1_hello",
    "options":null,
    "ip":"127.0.0.1"
}
*/

func writeLog(ctx context.Context, levelName string, message string, options map[string]interface{}) {
	reqUri := getStringByCtx(ctx, "request_uri")
	tag := strings.Replace(reqUri, "/", "_", -1)
	tag = strings.Replace(tag, ".", "_", -1)
	tag = strings.TrimLeft(tag, "_")

	logId := getStringByCtx(ctx, "log_id")
	if logId == "" {
		logId = gutils.RndUuid()
	}

	ua := getStringByCtx(ctx, "user_agent")
	logInfo := map[string]interface{}{
		"tag":            tag,
		"request_uri":    reqUri,
		"log_id":         logId,
		"options":        options,
		"ip":             getStringByCtx(ctx, "client_ip"),
		"ua":             ua,
		"plat":           helper.GetDeviceByUa(ua), // 当前设备匹配
		"request_method": getStringByCtx(ctx, "request_method"),
	}

	switch levelName {
	case "info":
		logger.Info(message, logInfo)
	case "debug":
		logger.Debug(message, logInfo)
	case "warn":
		logger.Warn(message, logInfo)
	case "error":
		logger.Error(message, logInfo)
	case "emergency":
		logger.DPanic(message, logInfo)
	default:
		logger.Info(message, logInfo)
	}
}

func getStringByCtx(ctx context.Context, key string) string {
	return helper.GetStringByCtx(ctx, key)
}

// Info info log.
func Info(ctx context.Context, message string, context map[string]interface{}) {
	writeLog(ctx, "info", message, context)
}

// Debug debug.
func Debug(ctx context.Context, message string, context map[string]interface{}) {
	writeLog(ctx, "debug", message, context)
}

// Warn warn log.
func Warn(ctx context.Context, message string, context map[string]interface{}) {
	writeLog(ctx, "warn", message, context)
}

// Error error log.
func Error(ctx context.Context, message string, context map[string]interface{}) {
	writeLog(ctx, "error", message, context)
}

// Emergency致命错误或panic捕获
func Emergency(ctx context.Context, message string, context map[string]interface{}) {
	writeLog(ctx, "emergency", message, context)
}

// Recover 异常捕获处理
func Recover(ctx context.Context) {
	if err := recover(); err != nil {
		Emergency(ctx, "exec panic", map[string]interface{}{
			"error":       err,
			"error_trace": string(debug.Stack()),
		})

		return
	}
}
