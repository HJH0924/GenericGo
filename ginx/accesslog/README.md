# AccessLog

## 使用 `zap` 和 `viper` 实现动态读取配置文件中的参数来动态改变是否读取请求体和响应体

```go
func GetLoggerMiddleware(zl logger.Logger) gin.HandlerFunc {
	ginLogger := accesslog.NewBuilder(func(ctx context.Context, al *accesslog.AccessLog) {
		zl.Debug("HTTP AccessLog",
			logger.Field{Key: "Method", Val: al.Method},
			logger.Field{Key: "URL", Val: al.URL},
			logger.Field{Key: "StatusCode", Val: al.StatusCode},
			logger.Field{Key: "Duration", Val: al.Duration},
			logger.Field{Key: "ReqBody", Val: al.ReqBody},
			logger.Field{Key: "RespBody", Val: al.RespBody},
		)
	})
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		// 当配置文件变更时，更新日志记录的行为
		allowReqBody := viper.GetBool("config.log.allowReqBody")
		allowRespBody := viper.GetBool("config.log.allowRespBody")
		zl.Debug("Configuration updated",
			logger.Field{Key: "config.log.allowReqBody", Val: allowReqBody},
			logger.Field{Key: "config.log.allowRespBody", Val: allowRespBody},
		)
		ginLogger.AllowReqBody(allowReqBody).AllowRespBody(allowRespBody)
	})
	return ginLogger.AllowReqBody(viper.GetBool("config.log.allowReqBody")).AllowRespBody(viper.GetBool("config.log.allowRespBody")).Build()
}
```

