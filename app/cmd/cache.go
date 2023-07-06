package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "go-web/pkg/cache"
    "go-web/pkg/console"
)

var CmdCache = &cobra.Command{
    Use:   "cache",
    Short: "HERE PUTS THE COMMAND DESCRIPTION",
}

var CmdCacheClear = &cobra.Command{
    Use:   "clear",
    Short: "",
    Run:   runCacheClear,
}

var CmdCacheForget = &cobra.Command{
    Use:   "forget",
    Short: "",
    Run:   runCacheForget,
}

var cacheKey string

func init() {
    // 注册 cache 命令的子命令
    CmdCache.AddCommand(
        CmdCacheClear,
        CmdCacheForget,
    )

    CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")
    err := CmdCacheForget.MarkFlagRequired("key")
    if err != nil {
        console.Error(err.Error())
    }
}

func runCacheClear(cmd *cobra.Command, args []string) {
    cache.Flush()
    console.Success("Cache cleared.")
}

func runCacheForget(cmd *cobra.Command, args []string) {
    cache.Forget(cacheKey)
    console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}
