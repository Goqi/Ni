package postgres

import (
	lib_postgres "Ni/pkg/js/libs/postgres"

	"Ni/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/postgres")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Types (value type)
			"PGClient": func() lib_postgres.PGClient { return lib_postgres.PGClient{} },

			// Types (pointer type)
			"NewPGClient": func() *lib_postgres.PGClient { return &lib_postgres.PGClient{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
