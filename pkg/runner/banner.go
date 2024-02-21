package runner

import (
	"fmt"

	"Ni/pkg/catalog/config"
	"github.com/projectdiscovery/gologger"
)

var banner = fmt.Sprintf(`Cell-nuclei二开 v%s`, config.Version)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("https://github.com/Goqi/Cell\n\n")
}
