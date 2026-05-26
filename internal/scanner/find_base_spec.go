//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 프로젝트 루트에서 기존 openapi.yaml 파일을 탐색한다
package scanner

import (
	"os"
	"path/filepath"
)

func FindBaseSpec(root string) string {
	candidates := []string{
		filepath.Join(root, "openapi.yaml"),
		filepath.Join(root, "openapi.yml"),
		filepath.Join(root, "api", "openapi.yaml"),
		filepath.Join(root, "api", "openapi.yml"),
		filepath.Join(root, "docs", "openapi.yaml"),
		filepath.Join(root, "docs", "openapi.yml"),
	}
	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}
