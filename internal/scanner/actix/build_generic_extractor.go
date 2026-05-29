//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what generic_type 노드(web::Path<...> 등)에서 extractor 정보를 구성한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func buildGenericExtractor(typeNode *sitter.Node, src []byte) *extractorInfo {
	scopedType := findChildByType(typeNode, "scoped_type_identifier")
	if scopedType == nil {
		return nil
	}
	kind := classifyExtractor(nodeText(scopedType, src))
	if kind == "" {
		return nil
	}
	typeArgs := findChildByType(typeNode, "type_arguments")
	if typeArgs == nil {
		return nil
	}
	return &extractorInfo{
		kind:     kind,
		typeName: extractTypeArgContent(typeArgs, src),
		rawType:  nodeText(typeNode, src),
	}
}
