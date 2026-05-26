//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what src/ 디렉토리에서 TypeScript 파일 경로를 수집한다
package nestjs

import (
	"os"
	"path/filepath"
	"strings"
)

// findTSFiles walks root/src/ and collects .ts file paths,
// excluding node_modules, dist, and .git directories.
func findTSFiles(root string) ([]string, error) {
	srcDir := filepath.Join(root, "src")
	info, err := os.Stat(srcDir)
	if err != nil || !info.IsDir() {
		return nil, nil
	}
	var files []string
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		base := info.Name()
		if info.IsDir() {
			switch base {
			case "node_modules", "dist", ".git":
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(base, ".ts") && !strings.HasSuffix(base, ".d.ts") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
