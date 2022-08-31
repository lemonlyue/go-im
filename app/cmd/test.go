package cmd

import (
	"gin-skeleton/pkg/console"
	"gin-skeleton/pkg/redis"
	"github.com/spf13/cobra"
	"time"
)

var CmdTest = &cobra.Command{
	Use: "test",
	Run: runTest,
	Short: "cmd test",
}

func runTest(cmd *cobra.Command, args []string)  {
	redis.Redis.Set("go", "Hello World", 10 * time.Second)
	console.Success(redis.Redis.Get("go"))
}
