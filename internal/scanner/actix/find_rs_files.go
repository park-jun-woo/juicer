//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 프로젝트 루트에서 Rust 파일 경로를 수집한다
package actix

import (
	"os"
	"path/filepath"
	"strings"
)

func findRsFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if info.IsDir() {
			if skipDirs[name] {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(name, ".rs") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
