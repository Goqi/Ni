package nucleicloud

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"Ernuclei/pkg/catalog/config"
)

// ReadCatalogChecksum reads catalog checksum from nuclei-templates repository
func ReadCatalogChecksum() map[string]string {
	config, _ := config.ReadConfiguration()
	if config == nil {
		return nil
	}

	checksumFile := filepath.Join(config.TemplatesDirectory, "templates-checksum.txt")
	file, err := os.Open(checksumFile)
	if err != nil {
		return nil
	}
	defer file.Close()

	checksums := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.SplitN(scanner.Text(), ":", 2)
		if len(text) < 2 {
			continue
		}
		path := strings.TrimPrefix(text[0], "nuclei-templates/")
		if strings.HasPrefix(path, ".") {
			continue
		}
		checksums[path] = text[1]
	}
	return checksums
}
