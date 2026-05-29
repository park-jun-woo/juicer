//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 프로젝트 전체에서 .ts 파일 경로를 수집한다 (node_modules/dist/build/.git 및 공통 테스트 디렉터리·테스트 파일 제외)
package fastify

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func findTSFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if info.IsDir() {
			switch name {
			case "node_modules", "dist", "build", ".git":
				return filepath.SkipDir
			}
			if scanner.IsTestDir(name) {
				return filepath.SkipDir
			}
			return nil
		}
		if scanner.IsTestFile(name) {
			return nil
		}
		if strings.HasSuffix(name, ".ts") && !strings.HasSuffix(name, ".d.ts") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
