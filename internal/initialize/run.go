package initialize

import (
	"fmt"
	"github.com/hainguyen27798/open-search-engine/global"
)

func Run() {
	LoadEnv()
	InitValidator()
	InitMongo()
	InitQDrant()

	r := InitRouter()
	err := r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
	if err != nil {
		panic(err)
	}
}
