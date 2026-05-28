//ff:func feature=scan type=parse control=sequence topic=flask
//ff:what 단일 Python 파일을 파싱하여 fileInfo를 반환한다
package flask

import (
	"os"
	"path/filepath"
)

// parseFile parses a single Python file and returns its fileInfo.
func parseFile(absRoot, absPath string) (*fileInfo, error) {
	src, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	root, err := parsePython(src)
	if err != nil {
		return nil, err
	}
	relPath, _ := filepath.Rel(absRoot, absPath)
	return &fileInfo{
		absPath: absPath,
		relPath: relPath,
		src:     src,
		root:    root,
	}, nil
}
