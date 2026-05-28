//ff:func feature=scan type=parse control=sequence topic=spring
//ff:what 단일 Java 파일을 파싱하여 fileInfo를 반환한다
package spring

import (
	"os"
	"path/filepath"
)

func parseFile(absRoot, absPath string) (*fileInfo, error) {
	src, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	root, err := parseJava(src)
	if err != nil {
		return nil, err
	}
	relPath, _ := filepath.Rel(absRoot, absPath)
	return &fileInfo{
		absPath:     absPath,
		relPath:     relPath,
		projectRoot: absRoot,
		src:         src,
		root:        root,
		imports:     extractImports(root, src),
	}, nil
}
