//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 프로젝트 루트에서 Python 파일 경로를 수집한다 (venv, __pycache__, migrations 등 제외)
package django

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// findPyFiles walks root and collects .py file paths,
// excluding venv, __pycache__, .git, dist, migrations, etc.
func findPyFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		name := info.Name()
		if info.IsDir() {
			if skipDirs[name] || scanner.IsTestDir(name) {
				return filepath.SkipDir
			}
			return nil
		}
		if scanner.IsTestFile(name) {
			return nil
		}
		if strings.HasSuffix(name, ".py") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
