//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 컨트롤러 클래스명에서 파일 경로를 추적하고 메서드를 파싱한다 (PSR-4)
package laravel

import (
	"path/filepath"
)

// resolveController finds and parses a controller file from its class name.
// It searches common Laravel paths: app/Http/Controllers/**/*.php
func resolveController(absRoot, className string, parsedFiles map[string]*fileInfo) *fileInfo {
	for _, fi := range parsedFiles {
		if classMatches(fi, className) {
			return fi
		}
	}

	candidates := []string{
		filepath.Join(absRoot, "app", "Http", "Controllers", className+".php"),
		filepath.Join(absRoot, "app", "Http", "Controllers", "Api", className+".php"),
		filepath.Join(absRoot, "app", "Http", "Controllers", "API", className+".php"),
	}
	for _, candidate := range candidates {
		if fi := parseControllerCandidate(absRoot, candidate); fi != nil {
			return fi
		}
	}
	return nil
}
