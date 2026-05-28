//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what Meta 내부 클래스에서 fields 리스트를 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// extractMetaFields extracts the fields list from a Meta inner class.
func extractMetaFields(body *sitter.Node, src []byte) []string {
	for _, classNode := range findAllByType(body, "class_definition") {
		nameNode := findChildByType(classNode, "identifier")
		if nameNode == nil || nodeText(nameNode, src) != "Meta" {
			continue
		}
		metaBody := findChildByType(classNode, "block")
		if metaBody == nil {
			continue
		}
		fields := extractFieldsList(metaBody, src)
		if fields != nil {
			return fields
		}
	}
	return nil
}
