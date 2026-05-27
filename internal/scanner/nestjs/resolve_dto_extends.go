//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what extends 절에서 부모 DTO를 해석하여 필드를 반환한다
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// resolveDTOExtends inspects the class extends clause and returns parent fields.
// Handles plain extends (extends ParentDto), factory extends (extends PartialType(Base)),
// and OmitType/PickType patterns.
func resolveDTOExtends(cls *sitter.Node, src []byte, filePath string, imports map[string]string, projectRoot string, cache map[string][]scanner.Field) []dtoField {
	heritage := findChildByType(cls, "class_heritage")
	if heritage == nil {
		return nil
	}
	ext := findChildByType(heritage, "extends_clause")
	if ext == nil {
		return nil
	}
	call := findChildByType(ext, "call_expression")
	if call != nil {
		return resolveDTOFactory(call, src, filePath, imports, projectRoot, cache)
	}
	// tree-sitter may produce "identifier" or "type_identifier" for extends target
	typeID := findChildByType(ext, "type_identifier")
	if typeID == nil {
		typeID = findChildByType(ext, "identifier")
	}
	if typeID == nil {
		return nil
	}
	parentName := nodeText(typeID, src)
	return resolveParentDTO(parentName, filePath, imports, projectRoot, cache)
}
