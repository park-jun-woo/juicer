//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 클래스 노드를 APIView로 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parseAPIViewClass parses a class node as an APIView if it is one. The class
// index lets it resolve custom intermediate base classes transitively.
func parseAPIViewClass(classNode *sitter.Node, fi fileInfo, idx classIndex) *apiviewInfo {
	nameNode := findChildByType(classNode, "identifier")
	if nameNode == nil {
		return nil
	}
	parents := extractParentClasses(classNode, fi.src)
	if !isAPIViewSubclass(parents, idx) || isViewSetSubclass(parents, idx) {
		return nil
	}
	view := &apiviewInfo{
		name: nodeText(nameNode, fi.src),
		file: fi.relPath,
		line: int(nameNode.StartPoint().Row) + 1,
	}
	body := findChildByType(classNode, "block")
	if body != nil {
		view.methods = extractHTTPMethods(body, fi.src)
		view.serializerClass = extractClassAttribute(body, "serializer_class", fi.src)
	}
	return view
}
