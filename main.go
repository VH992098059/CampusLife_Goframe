package main

import (
	_ "demo3/internal/packed"

	_ "demo3/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"demo3/internal/cmd"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
