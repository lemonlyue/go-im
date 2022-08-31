package main

import (
	"flag"
	"fmt"
	"gin-skeleton/app/cmd"
	"gin-skeleton/bootstrap"
	bootstrapConfig "gin-skeleton/config"
	"gin-skeleton/pkg/config"
	"gin-skeleton/pkg/console"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"os"
)

func init()  {
	bootstrapConfig.Initialize()
}

func main()  {

	// main
	var rootCmd = &cobra.Command{
		Use: config.Get("app.name"),
		Short: config.Get("app.name"),
		Long: `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		
		// 
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Init Config
			var env string
			flag.StringVar(&env, "env", "", "load .env file")
			flag.Parse()
			config.InitConfig(env)

			// init logger
			bootstrap.SetupLogger()

			// 设置 gin 的运行模式，支持 debug, release, test
			// release 会屏蔽调试信息，官方建议生产环境中使用
			// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
			// 故此设置为 release，有特殊情况手动改为 debug 即可
			gin.SetMode(gin.ReleaseMode)

			// init db
			bootstrap.SetupDB()

			// init redis
			bootstrap.SetupRedis()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdTest,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}