//ff:func feature=scan type=test control=sequence topic=express
//ff:what newScanCtx 테스트 헬퍼
package express

import sitter "github.com/smacker/go-tree-sitter"

func newScanCtx(absRoot string) *scanContext {
	return &scanContext{
		parsed:      map[string]*fileInfo{},
		allRouters:  map[string]map[string]bool{},
		schemas:     map[string]*sitter.Node{},
		schemaSrc:   map[string][]byte{},
		absRoot:     absRoot,
		pathAliases: map[string]string{},
	}
}
