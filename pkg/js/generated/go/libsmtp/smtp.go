package smtp

import (
	lib_smtp "Ni/pkg/js/libs/smtp"

	"Ni/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/smtp")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Types (value type)
			"IsSMTPResponse": func() lib_smtp.IsSMTPResponse { return lib_smtp.IsSMTPResponse{} },
			"SMTPClient":     func() lib_smtp.SMTPClient { return lib_smtp.SMTPClient{} },
			"SMTPMessage":    func() lib_smtp.SMTPMessage { return lib_smtp.SMTPMessage{} },

			// Types (pointer type)
			"NewIsSMTPResponse": func() *lib_smtp.IsSMTPResponse { return &lib_smtp.IsSMTPResponse{} },
			"NewSMTPClient":     func() *lib_smtp.SMTPClient { return &lib_smtp.SMTPClient{} },
			"NewSMTPMessage":    func() *lib_smtp.SMTPMessage { return &lib_smtp.SMTPMessage{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
