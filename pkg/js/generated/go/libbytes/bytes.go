package bytes

import (
	lib_bytes "Ni/pkg/js/libs/bytes"

	"Ni/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/bytes")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"NewBuffer": lib_bytes.NewBuffer,

			// Var and consts

			// Types (value type)
			"Buffer": func() lib_bytes.Buffer { return lib_bytes.Buffer{} },

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
