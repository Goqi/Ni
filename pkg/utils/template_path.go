package utils

import (
	"os"
	"path/filepath"
	"strings"

	"Ni/pkg/catalog/config"
)

const (
	// TemplatesRepoURL is the URL for files in nuclei-templates repository
	TemplatesRepoURL = "https://github.com/Goqi/pocs/blob/master/"
)

var configData *config.Config

func init() {
	configData, _ = config.ReadConfiguration()
}

// TemplatePathURL returns the Path and URL for the provided template
func TemplatePathURL(fullPath string) (string, string) {
	var templateDirectory string
	if configData != nil && configData.TemplatesDirectory != "" && strings.HasPrefix(fullPath, configData.TemplatesDirectory) {
		templateDirectory = configData.TemplatesDirectory
	} else {
		return "", ""
	}

	finalPath := strings.TrimPrefix(strings.TrimPrefix(fullPath, templateDirectory), "/")
	templateURL := TemplatesRepoURL + finalPath
	return finalPath, templateURL
}

// GetDefaultTemplatePath on default settings
// nuclei-templates路径
func GetDefaultTemplatePath() (string, error) {
	//home, err := homedir.Dir()
	home, err := os.Getwd()

	if err != nil {
		return "", err
	}
	return filepath.Join(home, "pocs"), nil
}
