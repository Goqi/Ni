package fs

import (
	lib_fs "Ni/pkg/js/libs/fs"

	"Ni/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/fs")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"ListDir":          lib_fs.ListDir,
			"ReadFile":         lib_fs.ReadFile,
			"ReadFileAsString": lib_fs.ReadFileAsString,
			"ReadFilesFromDir": lib_fs.ReadFilesFromDir,

			// Var and consts

			// Types (value type)

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
