package pop3

import (
	lib_pop3 "Ni/pkg/js/libs/pop3"

	"Ni/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/pop3")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Types (value type)
			"IsPOP3Response": func() lib_pop3.IsPOP3Response { return lib_pop3.IsPOP3Response{} },
			"Pop3Client":     func() lib_pop3.Pop3Client { return lib_pop3.Pop3Client{} },

			// Types (pointer type)
			"NewIsPOP3Response": func() *lib_pop3.IsPOP3Response { return &lib_pop3.IsPOP3Response{} },
			"NewPop3Client":     func() *lib_pop3.Pop3Client { return &lib_pop3.Pop3Client{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
