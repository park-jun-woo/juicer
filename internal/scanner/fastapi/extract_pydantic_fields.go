//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what Pydantic 클래스 본문에서 필드 정의를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractPydanticFields extracts field definitions from a Pydantic model class body.
func extractPydanticFields(cls *sitter.Node, src []byte) []pydanticField {
	var fields []pydanticField
	body := findChildByType(cls, "block")
	if body == nil {
		return nil
	}
	for i := 0; i < int(body.ChildCount()); i++ {
		child := body.Child(i)
		f := tryExtractField(child, src)
		if f != nil {
			fields = append(fields, *f)
		}
	}
	return fields
}
