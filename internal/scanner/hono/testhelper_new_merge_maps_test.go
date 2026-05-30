//ff:func feature=scan type=test control=sequence topic=hono
//ff:what newMergeMaps 테스트 헬퍼
package hono

import sitter "github.com/smacker/go-tree-sitter"

func newMergeMaps() (map[string]*fileInfo, map[string]map[string]bool, map[string]string, map[string]*sitter.Node, *[]routeGroup, map[string]map[string]string) {
	return map[string]*fileInfo{},
		map[string]map[string]bool{},
		map[string]string{},
		map[string]*sitter.Node{},
		&[]routeGroup{},
		map[string]map[string]string{}
}
