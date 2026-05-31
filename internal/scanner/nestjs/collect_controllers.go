//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what TypeScript 파일 목록에서 컨트롤러를 수집한다
package nestjs

import (
	"os"
	"path/filepath"
)

// collectControllers reads and parses TS files, returning controllers found.
func collectControllers(tsFiles []string, absRoot string) []controllerWithFile {
	var result []controllerWithFile
	for _, file := range tsFiles {
		src, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		astRoot, err := parseTypeScript(src)
		if err != nil {
			continue
		}
		relPath, _ := filepath.Rel(absRoot, file)
		controllers := extractControllers(astRoot, src, relPath, file, absRoot)
		for _, ci := range controllers {
			result = append(result, controllerWithFile{info: ci, absFile: file})
		}
	}
	return result
}
