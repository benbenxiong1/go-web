package cmd

import (
	"github.com/spf13/cobra"
)

// CmdPlay 临时调试代码
var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
	Args:  cobra.NoArgs,
}

func runPlay(cmd *cobra.Command, args []string) {
	//// 存进去 redis 中
	//redis.Redis.Set("hello", "hi form redis", 10*time.Second)
	//// 从 redis 中取出
	//console.Success(redis.Redis.Get("hello"))
}
