//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what 프로젝트 루트에서 Python 파일 경로를 수집한다
package flask

import (
	"os"
	"path/filepath"
	"strings"
)

// findPyFiles walks root and collects .py file paths,
// excluding venv, __pycache__, .git, dist, etc.
func findPyFiles(root string) ([]string, error) {
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
		if strings.HasSuffix(name, ".py") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
