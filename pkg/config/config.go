package config

import (
	"gin-skeleton/pkg/helpers"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

// viper instance
var viper *viperlib.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func init() {
	// init viper
	viper = viperlib.New()
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("appenv")
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	// load env
	loadEnv(env)
	// load config
	loadConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func loadEnv(envSuffix string) {
	basePath := ""
	if curPath, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			basePath = strings.Replace(strings.Replace(curPath, match(curPath, "/test*"), "", 1), match(curPath, `\test*`), "", 1)
		} else {
			basePath = curPath
		}
	}

	envPath := ".env"
	envPath = basePath + "/" + envPath
	if len(envSuffix) > 0 {
		filepath := ".env." + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			envPath = filepath
		}
	}
	viper.SetConfigFile(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// watch config change
	viper.WatchConfig()
}

func match(curPath string, matchStr string) string {
	re := regexp.MustCompile(matchStr)
	loc := re.FindStringIndex(curPath)
	matchPath := ""
	if loc != nil {
		matchPath = curPath[loc[0]:]
	}
	return matchPath
}

// get env value
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// Add config
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
