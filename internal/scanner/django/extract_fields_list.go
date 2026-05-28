//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what Meta body에서 fields = [...] 대입문의 필드명 리스트를 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// extractFieldsList extracts the fields = [...] assignment from a Meta body.
func extractFieldsList(metaBody *sitter.Node, src []byte) []string {
	for _, stmtNode := range childrenOfType(metaBody, "expression_statement") {
		assign := findChildByType(stmtNode, "assignment")
		if assign == nil {
			continue
		}
		left := findChildByType(assign, "identifier")
		if left == nil || nodeText(left, src) != "fields" {
			continue
		}
		listNode := findChildByType(assign, "list")
		if listNode == nil {
			continue
		}
		return extractStringLiterals(listNode, src)
	}
	return nil
}
