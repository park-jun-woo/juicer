//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what app.config.ts에서 apiPrefix 기본값 리터럴을 추출한다
package nestjs

import (
	"os"
	"path/filepath"
	"regexp"
)

// configDefaultRe matches: process.env.API_PREFIX || 'value' (single or double quotes).
var configDefaultRe = regexp.MustCompile(`process\.env\.API_PREFIX\s*\|\|\s*['"]([^'"]+)['"]`)

// readConfigDefault reads the default apiPrefix value from config files.
// It looks for patterns like: apiPrefix: process.env.API_PREFIX || 'api'
func readConfigDefault(root string) string {
	candidates := []string{
		filepath.Join(root, "src", "config", "app.config.ts"),
		filepath.Join(root, "src", "config", "app-config.ts"),
	}
	for _, path := range candidates {
		src, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		m := configDefaultRe.FindSubmatch(src)
		if m != nil {
			return string(m[1])
		}
	}
	return ""
}
