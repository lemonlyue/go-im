package main

import (
	"flag"
	"fmt"
	"gin-skeleton/app/cmd"
	"gin-skeleton/app/cmd/make"
	"gin-skeleton/bootstrap"
	bootstrapConfig "gin-skeleton/config"
	"gin-skeleton/pkg/config"
	"gin-skeleton/pkg/console"
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

			// init db
			bootstrap.SetupDB()

			// init redis
			bootstrap.SetupRedis()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdTest,
		cmd.CmdMigrate,
		make.CmdMake,
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