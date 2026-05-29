//ff:func feature=prisma type=parse control=sequence topic=prisma
//ff:what 입력 경로가 디렉터리면 *.prisma glob, 파일이면 단일 파일 목록 반환
package prisma

import (
	"fmt"
	"os"
	"path/filepath"
)

// collectPrismaFiles returns the .prisma files to parse: a directory is
// globbed for *.prisma, otherwise the path is used as a single file.
func collectPrismaFiles(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("stat %s: %w", path, err)
	}
	if !info.IsDir() {
		return []string{path}, nil
	}
	files, err := filepath.Glob(filepath.Join(path, "*.prisma"))
	if err != nil {
		return nil, fmt.Errorf("glob: %w", err)
	}
	return files, nil
}
