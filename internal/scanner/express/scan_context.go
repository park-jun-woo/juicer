//ff:type feature=scan type=model topic=express
//ff:what 스캔 컨텍스트 구조체 (Pass 1 결과)
package express

import sitter "github.com/smacker/go-tree-sitter"

type scanContext struct {
	parsed      map[string]*fileInfo
	allRouters  map[string]map[string]bool
	prefixMap   map[string]string
	absRoot     string
	pathAliases map[string]string
	schemas     map[string]*sitter.Node // schemaVarName -> z.object() node
	schemaSrc   map[string][]byte       // schemaVarName -> 해당 노드의 소스 바이트
}
