//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what 단일 파일의 Pass 1 결과를 전체 컨텍스트에 병합한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func mergePass1Result(path, absRoot string, parsed map[string]*fileInfo, honoVars map[string]map[string]bool, basePaths map[string]string, schemas map[string]*sitter.Node, allGroups *[]routeGroup, imports map[string]map[string]string) {
	r := scanOneFilePass1(path, absRoot)
	if r == nil {
		return
	}
	parsed[path] = r.fi
	if len(r.vars) > 0 {
		honoVars[path] = r.vars
	}
	for k, v := range r.bp {
		basePaths[prefixKey(path, k)] = v
	}
	for k, v := range r.schemas {
		schemas[k] = v
	}
	*allGroups = append(*allGroups, r.groups...)
	if len(r.imports) > 0 {
		imports[path] = r.imports
	}
}
