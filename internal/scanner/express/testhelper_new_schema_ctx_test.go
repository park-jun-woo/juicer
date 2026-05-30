//ff:func feature=scan type=test control=sequence topic=express
//ff:what newSchemaCtx 테스트 헬퍼
package express

import sitter "github.com/smacker/go-tree-sitter"

func newSchemaCtx() *scanContext {
	return &scanContext{
		schemas:   map[string]*sitter.Node{},
		schemaSrc: map[string][]byte{},
	}
}
