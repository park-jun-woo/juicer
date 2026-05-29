//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 프로젝트 전체에서 TS/JS 소스 파일(.ts/.tsx/.js/.jsx/.mjs/.cjs) 경로를 수집한다 (node_modules/dist/build/.git 및 공통 테스트 디렉터리·테스트 파일·.d.ts 제외)
package express

import (
	"os"
	"path/filepath"

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
		if isCollectableSource(name) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
