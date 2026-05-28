//ff:type feature=scan type=model topic=hono
//ff:what 스캔 컨텍스트 구조체 (Pass 1 결과)
package hono

import sitter "github.com/smacker/go-tree-sitter"

type scanContext struct {
	parsed     map[string]*fileInfo
	honoVars   map[string]map[string]bool // file -> set of Hono instance var names
	basePaths  map[string]string          // varName -> basePath
	groups     []routeGroup               // all route groups across files
	schemas    map[string]*sitter.Node    // schemaVarName -> z.object() node
	prefixMap  map[string]string          // varName -> resolved prefix
	absRoot    string
}
