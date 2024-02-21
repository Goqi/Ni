package structs

import (
	lib_structs "Ni/pkg/js/libs/structs"

	"Ni/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/structs")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"Pack":            lib_structs.Pack,
			"StructsCalcSize": lib_structs.StructsCalcSize,
			"Unpack":          lib_structs.Unpack,

			// Var and consts

			// Types (value type)

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
