//ff:func feature=scan type=test control=sequence topic=hono
//ff:what minimalCtx 테스트 헬퍼
package hono

import sitter "github.com/smacker/go-tree-sitter"

func minimalCtx() *scanContext {
	return &scanContext{
		parsed:    map[string]*fileInfo{},
		honoVars:  map[string]map[string]bool{},
		basePaths: map[string]string{},
		schemas:   map[string]*sitter.Node{},
		prefixMap: map[string]string{},
		imports:   map[string]map[string]string{},
	}
}
