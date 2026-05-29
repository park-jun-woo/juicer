//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what Pass 1: 파일 파싱, 라우터 수집, import 해석, use-prefix 수집, Zod 스키마 수집
package express

import sitter "github.com/smacker/go-tree-sitter"

func scanPass1(tsFiles []string, absRoot string) *scanContext {
	parsed := make(map[string]*fileInfo)
	allRouters := make(map[string]map[string]bool)
	schemas := make(map[string]*sitter.Node)
	schemaSrc := make(map[string][]byte)
	aliases := loadTsconfigPaths(absRoot)
	var allMounts []mountEntry
	for _, path := range tsFiles {
		mounts := scanOneFilePass1(path, parsed, allRouters, absRoot, aliases, schemas, schemaSrc)
		allMounts = append(allMounts, mounts...)
	}
	routerPrefixes := resolveRouterPrefixes(allMounts, allRouters)
	ctx := &scanContext{
		parsed:         parsed,
		allRouters:     allRouters,
		routerPrefixes: routerPrefixes,
		absRoot:        absRoot,
		pathAliases:    aliases,
		schemas:        schemas,
		schemaSrc:      schemaSrc,
	}
	// 미해결 스키마 import 추적 (최대 1회)
	resolveSchemaImports(ctx)
	return ctx
}
