//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what 테스트 헬퍼: PHP 소스 문자열을 fileInfo로 파싱
package laravel

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	php "github.com/smacker/go-tree-sitter/php"
)

func mustParsePHP(t *testing.T, src string) fileInfo {
	t.Helper()
	parser := sitter.NewParser()
	parser.SetLanguage(php.GetLanguage())
	tree, err := parser.ParseCtx(context.Background(), nil, []byte(src))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	return fileInfo{absPath: "test.php", relPath: "test.php", src: []byte(src), root: tree.RootNode()}
}
