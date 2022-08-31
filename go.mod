module gin-skeleton

go 1.16

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/spf13/cast v1.5.0
	github.com/spf13/viper v1.12.0
	go.uber.org/zap v1.17.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.8
)

replace cloud.google.com/go v1.12.1-0.20220712161005-5247643f0235 => github.com/googleapis/google-cloud-go v1.12.1-0.20220712161005-5247643f0235
