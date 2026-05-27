//ff:func feature=scan type=extract control=selection topic=nestjs
//ff:what PartialType/OmitType/PickType 팩토리 함수를 해석하여 필드를 반환한다
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// resolveDTOFactory handles factory-style extends: PartialType(Base), OmitType(Base, [...]), PickType(Base, [...]).
func resolveDTOFactory(call *sitter.Node, src []byte, filePath string, imports map[string]string, projectRoot string, cache map[string][]scanner.Field) []dtoField {
	funcIdent := findChildByType(call, "identifier")
	if funcIdent == nil {
		return nil
	}
	factoryName := nodeText(funcIdent, src)
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	baseClassName := extractFactoryBaseClass(args, src)
	if baseClassName == "" {
		return nil
	}
	parentFields := resolveParentDTOFields(baseClassName, filePath, imports, projectRoot, cache)
	switch factoryName {
	case "PartialType":
		return applyPartialType(parentFields)
	case "OmitType":
		omitNames := extractFactoryStringArray(args, src)
		return applyOmitType(parentFields, omitNames)
	case "PickType":
		pickNames := extractFactoryStringArray(args, src)
		return applyPickType(parentFields, pickNames)
	default:
		return scannerFieldsToDTOFields(parentFields)
	}
}
