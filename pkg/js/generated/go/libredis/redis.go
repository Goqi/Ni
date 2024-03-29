package redis

import (
	lib_redis "Ni/pkg/js/libs/redis"

	"github.com/dop251/goja"
	"Ni/pkg/js/gojs"
)

var (
	module = gojs.NewGojaModule("nuclei/redis")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"Connect":           lib_redis.Connect,
			"GetServerInfo":     lib_redis.GetServerInfo,
			"GetServerInfoAuth": lib_redis.GetServerInfoAuth,
			"IsAuthenticated":   lib_redis.IsAuthenticated,
			"RunLuaScript":      lib_redis.RunLuaScript,

			// Var and consts

			// Types (value type)

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
