package vnc

import (
	lib_vnc "Ni/pkg/js/libs/vnc"

	"Ni/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/vnc")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Types (value type)
			"IsVNCResponse": func() lib_vnc.IsVNCResponse { return lib_vnc.IsVNCResponse{} },
			"VNCClient":     func() lib_vnc.VNCClient { return lib_vnc.VNCClient{} },

			// Types (pointer type)
			"NewIsVNCResponse": func() *lib_vnc.IsVNCResponse { return &lib_vnc.IsVNCResponse{} },
			"NewVNCClient":     func() *lib_vnc.VNCClient { return &lib_vnc.VNCClient{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
