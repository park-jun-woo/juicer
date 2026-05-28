//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what 프로젝트 전체에서 .ts 파일 경로를 수집한다 (node_modules, dist, build, .git, test, __tests__, spec 제외)
package hono

import (
	"os"
	"path/filepath"
	"strings"
)

func findTSFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			switch info.Name() {
			case "node_modules", "dist", "build", ".git", "test", "__tests__", "spec":
				return filepath.SkipDir
			}
			return nil
		}
		name := info.Name()
		if strings.HasSuffix(name, ".ts") && !strings.HasSuffix(name, ".d.ts") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
