//ff:func feature=scan type=test control=sequence topic=express
//ff:what round5Ctx 테스트 헬퍼
package express

import sitter "github.com/smacker/go-tree-sitter"

func round5Ctx() *scanContext {
	return &scanContext{
		parsed:         map[string]*fileInfo{},
		allRouters:     map[string]map[string]bool{},
		routerPrefixes: map[routerKey][]string{},
		pathAliases:    map[string]string{},
		schemas:        map[string]*sitter.Node{},
		schemaSrc:      map[string][]byte{},
	}
}
