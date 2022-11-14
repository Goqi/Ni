package runner

import (
	"fmt"

	"Ernuclei/pkg/catalog/config"
	"github.com/projectdiscovery/gologger"
)

var banner = fmt.Sprintf(`
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/   %s
`, config.Version)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tprojectdiscovery.io\n\n")

	gologger.Print().Label("WRN").Msgf("Use with caution. You are responsible for your actions.\n")
	gologger.Print().Label("WRN").Msgf("Developers assume no liability and are not responsible for any misuse or damage.\n")
}
