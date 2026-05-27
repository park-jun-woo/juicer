//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 클래스의 부모가 globalModels에 있으면 모델로 등록한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// tryInheritClass checks whether cls has a parent in globalModels.
// If so, it registers the class as a model in fi and globalModels, returning true.
func tryInheritClass(cls *sitter.Node, fi *fileInfo, globalModels map[string]*fileInfo) bool {
	nameNode := findChildByType(cls, "identifier")
	if nameNode == nil {
		return false
	}
	name := nodeText(nameNode, fi.src)
	if _, exists := fi.models[name]; exists {
		return false
	}
	parents := collectParentNames(cls, fi.src)
	for _, p := range parents {
		if _, known := globalModels[p]; known {
			visited := map[string]bool{}
			fi.models[name] = resolveFieldsWithInheritance(cls, fi.root, fi.src, visited)
			globalModels[name] = fi
			return true
		}
	}
	return false
}
