//ff:func feature=scan type=parse control=sequence topic=laravel
//ff:what PHP 소스를 tree-sitter로 파싱하여 fileInfo를 반환한다
package laravel

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"
	php "github.com/smacker/go-tree-sitter/php"
)

func parseFile(absRoot, absPath string) (*fileInfo, error) {
	src, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", absPath, err)
	}
	parser := sitter.NewParser()
	parser.SetLanguage(php.GetLanguage())
	tree, err := parser.ParseCtx(context.Background(), nil, src)
	if err != nil {
		return nil, fmt.Errorf("tree-sitter parse %s: %w", absPath, err)
	}
	relPath, _ := filepath.Rel(absRoot, absPath)
	return &fileInfo{
		absPath: absPath,
		relPath: relPath,
		src:     src,
		root:    tree.RootNode(),
	}, nil
}
