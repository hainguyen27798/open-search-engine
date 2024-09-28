package initialize

import (
	"fmt"
	"github.com/hainguyen27798/open-typesense-search/global"
)

func Run() {
	LoadEnv()
	InitTypesense()

	r := InitRouter()
	err := r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
	if err != nil {
		panic(err)
	}
}
