//ff:type feature=scan type=model topic=hono
//ff:what 단일 파일의 Pass 1 결과 구조체
package hono

import sitter "github.com/smacker/go-tree-sitter"

type pass1FileResult struct {
	fi      *fileInfo
	vars    map[string]bool
	bp      map[string]string
	schemas map[string]*sitter.Node
	groups  []routeGroup
}
