//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 필요한 스키마 중 아직 해결되지 않은 것만 set으로 모은다
package express

import sitter "github.com/smacker/go-tree-sitter"

func buildUnresolvedSet(needed []string, schemas map[string]*sitter.Node) map[string]bool {
	set := make(map[string]bool)
	for _, name := range needed {
		if _, ok := schemas[name]; !ok {
			set[name] = true
		}
	}
	return set
}
