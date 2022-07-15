package main

import (
	_ "goframe/internal/packed"

	_ "goframe/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
