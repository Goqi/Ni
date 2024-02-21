package rdp

import (
	lib_rdp "Ni/pkg/js/libs/rdp"

	"Ni/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/rdp")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Types (value type)
			"CheckRDPAuthResponse": func() lib_rdp.CheckRDPAuthResponse { return lib_rdp.CheckRDPAuthResponse{} },
			"IsRDPResponse":        func() lib_rdp.IsRDPResponse { return lib_rdp.IsRDPResponse{} },
			"RDPClient":            func() lib_rdp.RDPClient { return lib_rdp.RDPClient{} },

			// Types (pointer type)
			"NewCheckRDPAuthResponse": func() *lib_rdp.CheckRDPAuthResponse { return &lib_rdp.CheckRDPAuthResponse{} },
			"NewIsRDPResponse":        func() *lib_rdp.IsRDPResponse { return &lib_rdp.IsRDPResponse{} },
			"NewRDPClient":            func() *lib_rdp.RDPClient { return &lib_rdp.RDPClient{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
