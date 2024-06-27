package main

import (
	_ "apis/internal/packed"

	_ "apis/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"apis/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
