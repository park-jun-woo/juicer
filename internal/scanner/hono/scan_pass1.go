//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what Pass 1: 모든 파일을 파싱하여 scanContext를 생성한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func scanPass1(tsFiles []string, absRoot string) *scanContext {
	parsed := make(map[string]*fileInfo)
	honoVars := make(map[string]map[string]bool)
	basePaths := make(map[string]string)
	schemas := make(map[string]*sitter.Node)
	imports := make(map[string]map[string]string)
	var allGroups []routeGroup

	for _, path := range tsFiles {
		mergePass1Result(path, absRoot, parsed, honoVars, basePaths, schemas, &allGroups, imports)
	}

	prefixMap := resolveRoutePrefixes(allGroups, basePaths, honoVars, imports)
	return &scanContext{
		parsed:    parsed,
		honoVars:  honoVars,
		basePaths: basePaths,
		groups:    allGroups,
		schemas:   schemas,
		prefixMap: prefixMap,
		imports:   imports,
		absRoot:   absRoot,
	}
}
