//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 클래스 노드를 Serializer로 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parseSerializerClass parses a class node as a Serializer if it is one.
func parseSerializerClass(classNode *sitter.Node, fi fileInfo) *serializerInfo {
	nameNode := findChildByType(classNode, "identifier")
	if nameNode == nil {
		return nil
	}
	parents := extractParentClasses(classNode, fi.src)
	if !isSerializerSubclass(parents) {
		return nil
	}
	si := &serializerInfo{name: nodeText(nameNode, fi.src)}
	body := findChildByType(classNode, "block")
	if body != nil {
		si.fields = extractSerializerFields(body, fi.src)
		if len(si.fields) == 0 {
			si.fields = metaFieldsToScannerFields(extractMetaFields(body, fi.src))
		}
	}
	return si
}
