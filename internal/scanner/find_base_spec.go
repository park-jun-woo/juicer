//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 프로젝트 루트에서 기존 openapi.yaml 파일을 탐색한다
package scanner

import "os"

func FindBaseSpec(root string) string {
	candidates := []string{
		root + "/openapi.yaml",
		root + "/openapi.yml",
		root + "/api/openapi.yaml",
		root + "/api/openapi.yml",
		root + "/docs/openapi.yaml",
		root + "/docs/openapi.yml",
	}
	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}
