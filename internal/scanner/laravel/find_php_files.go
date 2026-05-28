//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 프로젝트 루트에서 PHP 파일 경로를 수집한다 (vendor, storage, bootstrap, .git 등 제외)
package laravel

import (
	"os"
	"path/filepath"
	"strings"
)

func findPHPFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if skipDirs[info.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(info.Name(), ".php") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
