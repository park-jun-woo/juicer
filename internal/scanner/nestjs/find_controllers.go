//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what src/ 디렉토리(없으면 root 자체)에서 TypeScript 파일 경로를 수집한다
package nestjs

import (
	"os"
	"path/filepath"
	"strings"
)

// findTSFiles walks root/src/ (falling back to root itself when root/src is
// absent, e.g. when the scan root already ends in src) and collects .ts file
// paths, excluding node_modules, dist, and .git directories.
func findTSFiles(root string) ([]string, error) {
	walkDir := filepath.Join(root, "src")
	info, err := os.Stat(walkDir)
	if err != nil || !info.IsDir() {
		// Fallback: scan root directly when root/src does not exist. This
		// covers the case where the caller already passes a .../src path.
		walkDir = root
	}
	var files []string
	err = filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
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
