//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what Java import 문에서 DTO 소스 파일 경로를 추적한다
package quarkus

import (
	"os"
	"path/filepath"
	"strings"
)

func resolveImportPath(projectRoot, fqcn string) string {
	relative := strings.ReplaceAll(fqcn, ".", string(filepath.Separator)) + ".java"
	candidates := []string{
		filepath.Join(projectRoot, "src", "main", "java", relative),
		filepath.Join(projectRoot, relative),
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	return ""
}
