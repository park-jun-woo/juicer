//ff:func feature=scan type=test control=sequence topic=express
//ff:what firstImportSpec 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstImportSpec(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	specs := findAllByType(fi.Root, "import_specifier")
	if len(specs) == 0 {
		t.Fatal("no import_specifier")
	}
	return specs[0]
}
