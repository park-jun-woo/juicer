//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 함수 시그니처에서 web::Path, web::Json, web::Query, web::Form extractor를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type extractorInfo struct {
	kind     string // "path", "json", "query", "form"
	typeName string // inner type name (e.g., "i64", "CreateUserRequest")
	rawType  string // full type text
}

func extractExtractors(funcNode *sitter.Node, src []byte) []extractorInfo {
	params := findChildByType(funcNode, "parameters")
	if params == nil {
		return nil
	}

	var result []extractorInfo
	for _, param := range childrenOfType(params, "parameter") {
		ext := parseExtractor(param, src)
		if ext != nil {
			result = append(result, *ext)
		}
	}
	return result
}

func parseExtractor(param *sitter.Node, src []byte) *extractorInfo {
	typeNode := findParamType(param)
	if typeNode == nil {
		return nil
	}

	typeText := nodeText(typeNode, src)

	if typeNode.Type() == "generic_type" {
		scopedType := findChildByType(typeNode, "scoped_type_identifier")
		if scopedType == nil {
			return nil
		}
		scopedText := nodeText(scopedType, src)
		kind := classifyExtractor(scopedText)
		if kind == "" {
			return nil
		}

		typeArgs := findChildByType(typeNode, "type_arguments")
		if typeArgs == nil {
			return nil
		}
		inner := extractTypeArgContent(typeArgs, src)

		return &extractorInfo{
			kind:     kind,
			typeName: inner,
			rawType:  typeText,
		}
	}

	return nil
}

func findParamType(param *sitter.Node) *sitter.Node {
	// parameter children: identifier, ":", type
	for i := 0; i < int(param.ChildCount()); i++ {
		child := param.Child(i)
		switch child.Type() {
		case "generic_type", "type_identifier", "scoped_type_identifier",
			"reference_type", "primitive_type":
			return child
		}
	}
	return nil
}

func classifyExtractor(scopedType string) string {
	switch scopedType {
	case "web::Path":
		return "path"
	case "web::Json":
		return "json"
	case "web::Query":
		return "query"
	case "web::Form":
		return "form"
	}
	return ""
}

func extractTypeArgContent(typeArgs *sitter.Node, src []byte) string {
	// type_arguments: <, inner_type, >
	// collect all named children between < and >
	var parts []string
	for i := 0; i < int(typeArgs.ChildCount()); i++ {
		child := typeArgs.Child(i)
		if child.IsNamed() {
			parts = append(parts, nodeText(child, src))
		}
	}
	if len(parts) == 1 {
		return parts[0]
	}
	// tuple type or multiple
	result := ""
	for i := 0; i < int(typeArgs.ChildCount()); i++ {
		child := typeArgs.Child(i)
		t := child.Type()
		if t == "<" || t == ">" {
			continue
		}
		result += nodeText(child, src)
	}
	return result
}
